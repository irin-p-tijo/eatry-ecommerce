package repository

import (
	"eatry/pkg/domain"
	"eatry/pkg/helper"
	interfaces "eatry/pkg/repository/interfaces"
	"eatry/pkg/utils/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &AdminRepository{
		DB: DB,
	}
}
func (ad *AdminRepository) CheckAdminAvailability(admin models.AdminSignUp) bool {

	var count int
	if err := ad.DB.Raw("select count(*) from admin_details where email = ?", admin.Email).Scan(&count).Error; err != nil {
		return false
	}

	return count > 0

}
func (ad *AdminRepository) CreateAdmin(admin models.AdminSignUp) (models.AdminDetailsResponse, error) {

	var adminDetails models.AdminDetailsResponse
	if err := ad.DB.Raw("insert into admin_details (name,email,password) values (?, ?, ?) returning id, name, email", admin.Name, admin.Email, admin.Password).Scan(&adminDetails).Error; err != nil {
		return models.AdminDetailsResponse{}, err
	}

	return adminDetails, nil

}
func (ad *AdminRepository) LoginHandler(adminDetails models.AdminLogin) (domain.AdminDetails, error) {

	var adminCompareDetails domain.AdminDetails
	if err := ad.DB.Raw("select * from admin_details where email = ? ", adminDetails.Email).Scan(&adminCompareDetails).Error; err != nil {
		return domain.AdminDetails{}, err
	}

	return adminCompareDetails, nil
}
func (ad *AdminRepository) GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error) {

	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var userDetails []models.UserDetailsAtAdmin

	if err := ad.DB.Raw("select id,name,email,phone,blocked from users limit ? offset ?", count, offset).Scan(&userDetails).Error; err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil

}
func (ad *AdminRepository) GetUserById(userID int) (domain.Users, error) {
	query := fmt.Sprintf("select * from users where id='%d'", userID)
	var userDetails domain.Users
	if err := ad.DB.Raw(query).Scan(&userDetails).Error; err != nil {
		return domain.Users{}, err
	}
	return userDetails, nil
}
func (ad *AdminRepository) UpdateBlockUserByID(user domain.Users) error {
	err := ad.DB.Exec("update users set blocked = ? where id = ? ", user.Blocked, user.ID).Error
	if err != nil {
		return err
	}
	return nil
}
func (ad *AdminRepository) UserDetailsDashboard() (models.DashboardUsers, error) {
	var userdetails models.DashboardUsers

	err := ad.DB.Raw("select count(*) from users").Scan(&userdetails.TotalUsers).Error
	if err != nil {
		return models.DashboardUsers{}, nil
	}
	err = ad.DB.Raw("select count(*) from users where blocked=true").Scan(&userdetails.BlockedUsers).Error
	if err != nil {
		return models.DashboardUsers{}, nil
	}
	return userdetails, nil
}
func (ad *AdminRepository) ProductDetailsDashboard() (models.DashboardProducts, error) {
	var productdetails models.DashboardProducts

	err := ad.DB.Raw("select count(*) from products").Scan(&productdetails.TotalProducts).Error
	if err != nil {
		return models.DashboardProducts{}, nil
	}
	err = ad.DB.Raw("select count(*) from products where quantity=0").Scan(&productdetails.OutOfStockProducts).Error
	if err != nil {
		return models.DashboardProducts{}, nil
	}

	return productdetails, nil
}
func (ad *AdminRepository) OrderDetailsDashboard() (models.DashboardOrders, error) {
	var orderdetails models.DashboardOrders

	err := ad.DB.Raw("select count(*) from orders").Scan(&orderdetails.TotalOrders).Error
	if err != nil {
		return models.DashboardOrders{}, nil
	}

	err = ad.DB.Raw("select count(*) from orders where shipment_status='cancelled'").Scan(&orderdetails.CancelledOrders).Error
	if err != nil {
		return models.DashboardOrders{}, nil
	}
	err = ad.DB.Raw("select count(*) from orders where shipment_status='pending' or shipment_status='processing'").Scan(&orderdetails.PendingOrders).Error
	if err != nil {
		return models.DashboardOrders{}, nil
	}
	err = ad.DB.Raw("select count(*) from orders where payment_status='paid' ").Scan(&orderdetails.CompletedOrders).Error
	if err != nil {
		return models.DashboardOrders{}, nil
	}
	err = ad.DB.Raw("select sum(quantity) from user_order_items").Scan(&orderdetails.TotalOrderItems).Error
	if err != nil {
		return models.DashboardOrders{}, nil
	}
	return orderdetails, nil
}
func (ad *AdminRepository) AmountDetails() (models.DashboardAmount, error) {
	var amountDetails models.DashboardAmount

	err := ad.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'paid'").Scan(&amountDetails.CreditedAmount).Error
	if err != nil {
		return models.DashboardAmount{}, err
	}
	err = ad.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'not paid' and shipment_status = 'processing' or shipment_status = 'pending' or shipment_status = 'order placed' ").Scan(&amountDetails.PendingAmount).Error
	if err != nil {
		return models.DashboardAmount{}, nil
	}
	return amountDetails, nil
}
func (ad *AdminRepository) TotalRevenue() (models.DashboardRevenue, error) {
	var revenueDetails models.DashboardRevenue
	startTime := time.Now().AddDate(0, 0, -1)
	endTime := time.Now()
	err := ad.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'paid'  and created_at >= ? and created_at <= ?", startTime, endTime).Scan(&revenueDetails.TodayRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}
	startTime, endTime = helper.GetTimeFromPeriod("month")
	err = ad.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'paid'  and created_at >= ? and created_at <= ?", startTime, endTime).Scan(&revenueDetails.MonthRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}
	startTime, endTime = helper.GetTimeFromPeriod("year")
	err = ad.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'paid'  and created_at >= ? and created_at <= ?", startTime, endTime).Scan(&revenueDetails.YearRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}
	return revenueDetails, nil

}

func (ad *AdminRepository) FilterSalesReport(startTime time.Time, endTime time.Time) (models.SalesReport, error) {
	var salesReport models.SalesReport

	result := ad.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'paid'  and created_at >= ? and created_at <= ?", startTime, endTime).Scan(&salesReport.TotalSales)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	result = ad.DB.Raw("select count(*) from orders  where created_at >= ? and created_at <= ? ", startTime, endTime).Scan(&salesReport.TotalOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	result = ad.DB.Raw("select count(*) from orders where payment_status = 'paid'  and  created_at >= ? and created_at <= ?", startTime, endTime).Scan(&salesReport.CompletedOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	result = ad.DB.Raw("select count(*) from orders where shipment_status = 'processing' and approval = false and  created_at >= ? and created_at <= ?", startTime, endTime).Scan(&salesReport.PendingOrders)
	if result.Error != nil {
		return models.SalesReport{}, nil
	}
	var productIDs []int

	result = ad.DB.Raw("select product_id from user_order_items group by product_id order by sum(quantity) desc limit 10").Scan(&productIDs)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	var bestsellingProducts []string
	for _, productID := range productIDs {
		var productName string
		result = ad.DB.Raw("select name from products where id = ?", productID).Scan(&productName)
		if result.Error != nil {
			return models.SalesReport{}, result.Error
		}
		bestsellingProducts = append(bestsellingProducts, productName)
	}

	salesReport.TopTenBestsellingProducts = bestsellingProducts

	var categoryIDs []int
	result = ad.DB.Raw("SELECT category_id FROM products GROUP BY category_id ORDER BY SUM(quantity) DESC LIMIT 10 ").Scan(&categoryIDs)

	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	var bestSellingCategory []string
	for _, categoryID := range categoryIDs {
		var categoryName string
		result = ad.DB.Raw("SELECT category FROM categories WHERE id = ?", categoryID).Scan(&categoryName)
		if result.Error != nil {
			return models.SalesReport{}, result.Error
		}
		bestSellingCategory = append(bestSellingCategory, categoryName)
	}
	salesReport.TopTenBestsellingCategories = bestSellingCategory

	return salesReport, nil
}
