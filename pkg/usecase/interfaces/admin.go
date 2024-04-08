package interfaces

import (
	"eatry/pkg/domain"
	"eatry/pkg/utils/models"
)

type AdminUseCase interface {
	CreateAdmin(admin models.AdminSignUp) (domain.TokenAdmin, error)
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
	GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error)
	BlockUser(userID int) error
	UnBlockUser(userID int) error
	Dashboard() (models.DashBoardTotal, error)
	FilterSalesReport(timePeriod string) (models.SalesReport, error)
	ExecuteSalesReportByDate(startDate, endDate string) (models.SalesReport, error)
}
