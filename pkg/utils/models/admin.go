package models

type AdminSignUp struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
type AdminDetailsResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name" `
	Email string `json:"email" `
}
type AdminLogin struct {
	Email    string `json:"email"  binding:"required" validate:"required"`
	Password string `json:"password"  binding:"required" validate:"min=8,max=20"`
}

type DashboardUsers struct {
	TotalUsers   int
	BlockedUsers int
}

type DashboardProducts struct {
	TotalProducts      int64
	OutOfStockProducts int64
}
type DashboardOrders struct {
	TotalOrders     int64
	CancelledOrders int64
	PendingOrders   int64
	CompletedOrders int64
	TotalOrderItems int64
}
type DashboardRevenue struct {
	TodayRevenue float64
	MonthRevenue float64
	YearRevenue  float64
}
type DashboardAmount struct {
	CreditedAmount float64
	PendingAmount  float64
}
type DashBoardTotal struct {
	Users    DashboardUsers
	Products DashboardProducts
	Orders   DashboardOrders
	Revenue  DashboardRevenue
	Amount   DashboardAmount
}
type SalesReport struct {
	TotalSales                  float64
	TotalOrders                 int
	CompletedOrders             int
	PendingOrders               int
	TopTenBestsellingProducts   []string //changes to slice
	TopTenBestsellingCategories []string
}
