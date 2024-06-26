package interfaces

import (
	"eatry/pkg/domain"
	"eatry/pkg/utils/models"
)

type OrderRepository interface {
	DoesCartExist(userID int) (bool, error)
	AddressExist(orderBody models.OrderIncoming) (bool, error)
	GetCouponDiscountPrice(UserID int, GrandTotal float64) (float64, error)
	UpdateCouponDetails(discount_price float64, UserID int) error
	CreateOrder(orderDetails domain.Order) error
	AddOrderItems(orderItemDetails domain.UserOrderItem, UserID int, ProductID uint, Quantity float64) error
	GetBriefOrderDetails(orderID string) (domain.OrderSuccessResponse, error)
	GetOrderDetailsByOrderId(orderID string) (models.CombinedOrderDetails, error)
	GetOrders(orderID string) (domain.Order, error)
	UserOrderRelationship(orderID string, userID int) (int, error)
	GetOrderDetails(userID int, page int, count int) ([]models.FullOrderDetails, error)
	CancelOrders(orderID string) error
	GetShipmentStatus(orderID string) (string, error)
	GetProductDetailsFromOrders(orderID string) ([]models.OrderProducts, error)
	UpdateQuantityOfProduct(orderProducts []models.OrderProducts) error
	GetOrderDetailsBrief(page int) ([]models.CombinedOrderDetails, error)
	GetWalletAmount(UserID int) (float64, error)
	UpdateWalletAmount(walletAmount float64, UserID int) error
	GetPaymentStatus(orderID string) (string, error)
	GetPriceoftheproduct(orderID string) (float64, error)
	CheckOrderID(orderID string) (bool, error)
	ApproveOrder(orderID string) error
	GetOrderDetailsofAproduct(orderID string) (models.OrderDetails, error)
	GetAddressDetailsFromID(orderID string) (models.Address, error)
	GetItemsByOrderId(orderID string) ([]models.ProductDetails, error)
	GetOrderDetailsByID(orderID string) (models.CombinedOrderDetails, error)
}
