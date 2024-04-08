package repository

import (
	"eatry/pkg/domain"
	"eatry/pkg/helper"
	interfaces "eatry/pkg/repository/interfaces"
	"eatry/pkg/utils/models"
	"fmt"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}
func (o *OrderRepository) DoesCartExist(userID int) (bool, error) {
	var exist bool
	err := o.DB.Raw("select exists(select 1 from carts where user_id=?)", userID).Scan(&exist).Error
	if err != nil {
		return false, err
	}
	return exist, nil
}
func (o *OrderRepository) AddressExist(orderBody models.OrderIncoming) (bool, error) {
	var count int
	err := o.DB.Raw("select count(*)from addresses where user_id=? and id=?", orderBody.UserID, orderBody.AddressID).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (o *OrderRepository) GetCouponDiscountPrice(UserID int, GrandTotal float64) (float64, error) {
	discountPrice, err := helper.GetCouponDiscountPrice(UserID, GrandTotal, o.DB)
	if err != nil {
		return 0.0, err
	}

	return discountPrice, nil
}
func (o *OrderRepository) UpdateCouponDetails(discount_price float64, UserID int) error {

	if discount_price != 0.0 {
		err := o.DB.Exec("update used_coupons set used = true where user_id = ?", UserID).Error
		if err != nil {
			return err
		}
	}
	return nil
}
func (o *OrderRepository) CreateOrder(orderDetails domain.Order) error {

	err := o.DB.Create(&orderDetails).Error
	if err != nil {
		return err
	}
	return nil

}
func (o *OrderRepository) AddOrderItems(orderItemDetails domain.UserOrderItem, userID int, ProductID uint, Quantity float64) error {

	// after creating the order delete all cart items and also update the quantity of the product
	err := o.DB.Omit("id").Create(&orderItemDetails).Error
	if err != nil {
		return err
	}

	err = o.DB.Exec("delete from carts where user_id = ? and product_id = ?", userID, ProductID).Error
	if err != nil {
		return err
	}

	err = o.DB.Exec("update products set quantity = quantity - ? where id = ?", Quantity, ProductID).Error
	if err != nil {
		return err
	}

	return nil

}
func (o *OrderRepository) GetBriefOrderDetails(orderID string) (domain.OrderSuccessResponse, error) {

	var orderSuccessResponse domain.OrderSuccessResponse
	o.DB.Raw("select order_id,shipment_status from orders where order_id = ?", orderID).Scan(&orderSuccessResponse)
	return orderSuccessResponse, nil

}
func (o *OrderRepository) GetOrderDetailsByOrderId(orderID string) (models.CombinedOrderDetails, error) {
	var orderDetails models.CombinedOrderDetails

	err := o.DB.Raw("select  orders.order_id,orders.final_price,orders.shipment_status,orders.payment_status,users.name,users.email,users.phone,addresses.house_name,addresses.street,addresses.city,addresses.district,addresses.state,addresses.pin from orders join users on  orders.user_id = users.id join addresses on orders.address_id=addresses.id where orders.order_id=? ", orderID).Scan(&orderDetails).Error
	if err != nil {
		return models.CombinedOrderDetails{}, nil
	}

	return orderDetails, nil
}
func (o *OrderRepository) GetOrders(orderID string) (domain.Order, error) {
	var body domain.Order

	if err := o.DB.Raw("select * from orders where order_id=?", orderID).Scan(&body).Error; err != nil {
		return domain.Order{}, err
	}
	return body, nil
}
func (o *OrderRepository) UserOrderRelationship(orderID string, userID int) (int, error) {

	var testUserID int
	err := o.DB.Raw("select user_id from orders where order_id = ?", orderID).Scan(&testUserID).Error
	if err != nil {
		return -1, err
	}
	return testUserID, nil

}
func (o *OrderRepository) GetProductDetailsFromOrders(orderID string) ([]models.OrderProducts, error) {

	var orderProductDetails []models.OrderProducts
	if err := o.DB.Raw("select product_id,quantity from user_order_items where order_id = ?", orderID).Scan(&orderProductDetails).Error; err != nil {
		return []models.OrderProducts{}, err
	}

	return orderProductDetails, nil
}
func (o *OrderRepository) GetShipmentStatus(orderID string) (string, error) {

	var shipmentStatus string
	err := o.DB.Raw("select shipment_status from orders where order_id = ?", orderID).Scan(&shipmentStatus).Error
	if err != nil {
		return "", err
	}

	return shipmentStatus, nil

}
func (o *OrderRepository) GetOrderDetails(userID int, page int, count int) ([]models.FullOrderDetails, error) {
	// details of order created byt his particular user
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count

	var orderDetails []models.OrderDetails
	o.DB.Raw("select order_id,final_price,shipment_status,payment_status from orders where user_id = ? limit ? offset ? ", userID, count, offset).Scan(&orderDetails)
	fmt.Println(orderDetails)

	var fullOrderDetails []models.FullOrderDetails
	// for each order select all the associated products and their details
	for _, od := range orderDetails {

		var orderProductDetails []models.OrderProductDetails
		o.DB.Raw("select user_order_items.product_id,products.name,user_order_items.quantity,user_order_items.total_price from user_order_items inner join products on user_order_items.product_id = products.id where user_order_items.order_id = ?", od.OrderId).Scan(&orderProductDetails)
		fullOrderDetails = append(fullOrderDetails, models.FullOrderDetails{OrderDetails: od, OrderProductDetails: orderProductDetails})

	}

	return fullOrderDetails, nil

}
func (o *OrderRepository) CancelOrders(orderID string) error {
	shipmentStatus := "cancelled"
	err := o.DB.Exec("update orders set shipment_status = ? where order_id = ?", shipmentStatus, orderID).Error
	if err != nil {
		return err
	}
	var paymentMethod int
	err = o.DB.Raw("select payment_method_id from orders where order_id = ?", orderID).Scan(&paymentMethod).Error
	if err != nil {
		return err
	}
	if paymentMethod == 1 || paymentMethod == 3 {
		err = o.DB.Exec("update orders set payment_status = 'refunded'  where order_id = ?", orderID).Error
		if err != nil {
			return err
		}
	}
	return nil
}
func (o *OrderRepository) UpdateQuantityOfProduct(orderProducts []models.OrderProducts) error {
	for _, od := range orderProducts {
		var quantity int
		if err := o.DB.Raw("select quantity from products where id = ?", od.ProductId).Scan(&quantity).Error; err != nil {
			return err
		}
		od.Quantity += quantity
		if err := o.DB.Exec("update products set quantity = ? where id = ?", od.Quantity, od.ProductId).Error; err != nil {
			return err
		}
	}
	return nil
}
func (o *OrderRepository) GetOrderDetailsBrief(page int) ([]models.CombinedOrderDetails, error) {

	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 2
	var orderDetails []models.CombinedOrderDetails

	err := o.DB.Raw("select orders.order_id,orders.final_price,orders.shipment_status,orders.payment_status,users.name,users.email,users.phone,addresses.house_name,addresses.state,addresses.pin,addresses.street,addresses.city from orders inner join users on orders.user_id = users.id inner join addresses on users.id = addresses.user_id limit ? offset ?", 2, offset).Scan(&orderDetails).Error

	if err != nil {
		return []models.CombinedOrderDetails{}, nil
	}

	return orderDetails, nil
}
func (o *OrderRepository) GetWalletAmount(UserID int) (float64, error) {

	var walletAvailable float64
	err := o.DB.Raw("select wallet_amount from wallets where user_id = ?", UserID).Scan(&walletAvailable).Error
	if err != nil {
		return 0.0, err
	}

	return walletAvailable, nil
}
func (o *OrderRepository) UpdateWalletAmount(walletAmount float64, UserID int) error {

	err := o.DB.Exec("update wallets set wallet_amount = ? where user_id = ? ", walletAmount, UserID).Error
	if err != nil {
		return err
	}
	return nil

}
func (o *OrderRepository) GetPaymentStatus(orderID string) (string, error) {
	var paymentstatus string
	err := o.DB.Raw("select payment_status from orders where order_id=?", orderID).Scan(&paymentstatus).Error
	if err != nil {
		return "", err
	}
	return paymentstatus, nil

}
func (o *OrderRepository) GetPriceoftheproduct(orderID string) (float64, error) {
	var a float64
	err := o.DB.Raw("select grand_total from orders where order_id=?", orderID).Scan(&a).Error
	if err != nil {
		return 0.0, err
	}
	return a, nil
}
func (o *OrderRepository) CheckOrderID(orderID string) (bool, error) {

	var count int
	err := o.DB.Raw("select count(*) from orders where order_id = ?", orderID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil

}
func (o *OrderRepository) ApproveOrder(orderID string) error {

	err := o.DB.Exec("update orders set shipment_status = 'order placed',approval=true where order_id = ?", orderID).Error
	if err != nil {
		return err
	}
	return nil
}
func (o *OrderRepository) GetOrderDetailsofAproduct(orderID string) (models.OrderDetails, error) {
	var OrderDetails models.OrderDetails

	if err := o.DB.Raw("select order_id,final_price,shipment_status,payment_status from orders where order_id = ?", orderID).Scan(&OrderDetails).Error; err != nil {
		return models.OrderDetails{}, err
	}
	return OrderDetails, nil
}

func (o *OrderRepository) GetAddressDetailsFromID(orderID string) (models.Address, error) {
	var address models.Address
	var addressId int

	if err := o.DB.Raw("select address_id from orders where order_id =?", orderID).Scan(&address).Error; err != nil {
		return models.Address{}, err
	}
	if err := o.DB.Raw("select * from addresses where id=?", addressId).Scan(&address).Error; err != nil {
		return models.Address{}, err
	}
	return models.Address{}, nil
}
func (o *OrderRepository) GetOrderDetailsByID(orderID string) (models.CombinedOrderDetails, error) {
	var orders models.CombinedOrderDetails

	query := `select orders.user_id,users.name,users.email,users.phone,addresses.house_name,addresses.street,addresses.city,addresses.state,addresses.pin,orders.address_id,orders.payment_method_id,payment_methods.payment_name,orders.final_price from orders inner join users on orders.user_id=users.id inner join addresses on orders.address_id=addresses.id inner join payment_methods on orders.payment_method_id=payment_methods.id where orders.order_id=?  `

	err := o.DB.Raw(query, orderID).Scan(&orders).Error
	if err != nil {
		return models.CombinedOrderDetails{}, err
	}
	return orders, nil
}
func (o *OrderRepository) GetItemsByOrderId(orderID string) ([]models.ProductDetails, error) {
	var items []models.ProductDetails

	query := "select products.name,user_order_items.quantity,products.price,user_order_items.total_price from user_order_items  join products on user_order_items.product_id=products.id where user_order_items.order_id=?"

	if err := o.DB.Raw(query, orderID).Scan(&items).Error; err != nil {
		return []models.ProductDetails{}, err
	}
	return items, nil

}
