package interfaces

import (
	"eatry/pkg/domain"
	"eatry/pkg/utils/models"
	"time"
)

type AdminRepository interface {
	CheckAdminAvailability(admin models.AdminSignUp) bool
	CreateAdmin(admin models.AdminSignUp) (models.AdminDetailsResponse, error)
	LoginHandler(adminDetails models.AdminLogin) (domain.AdminDetails, error)
	GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error)
	GetUserById(userID int) (domain.Users, error)
	UpdateBlockUserByID(user domain.Users) error
	UserDetailsDashboard() (models.DashboardUsers, error)
	ProductDetailsDashboard() (models.DashboardProducts, error)
	OrderDetailsDashboard() (models.DashboardOrders, error)
	AmountDetails() (models.DashboardAmount, error)
	TotalRevenue() (models.DashboardRevenue, error)

	FilterSalesReport(startTime time.Time, endTime time.Time) (models.SalesReport, error)
}
