package usecase

import (
	"eatry/pkg/domain"
	"eatry/pkg/helper"
	interfaces "eatry/pkg/repository/interfaces"
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/models"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase struct {
	adminRepository interfaces.AdminRepository
}

func NewAdminUseCase(adminrepository interfaces.AdminRepository) services.AdminUseCase {
	return &AdminUseCase{
		adminRepository: adminrepository,
	}
}

func (ad *AdminUseCase) CreateAdmin(admin models.AdminSignUp) (domain.TokenAdmin, error) {

	if err := validator.New().Struct(admin); err != nil {
		return domain.TokenAdmin{}, err
	}

	userExist := ad.adminRepository.CheckAdminAvailability(admin)
	if userExist {
		return domain.TokenAdmin{}, errors.New("admin already exist, sign in")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return domain.TokenAdmin{}, errors.New("internal server error")
	}

	admin.Password = string(hashedPassword)

	adminDetails, err := ad.adminRepository.CreateAdmin(admin)
	if err != nil {
		return domain.TokenAdmin{}, err
	}
	tokenString, err := helper.GenerateTokenAdmin(adminDetails)

	if err != nil {
		return domain.TokenAdmin{}, err
	}

	return domain.TokenAdmin{
		Admin: adminDetails,
		Token: tokenString,
	}, nil

}
func (ad *AdminUseCase) LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error) {

	adminCompareDetails, err := ad.adminRepository.LoginHandler(adminDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	var adminDetailsResponse models.AdminDetailsResponse

	err = copier.Copy(&adminDetailsResponse, &adminCompareDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	tokenString, err := helper.GenerateTokenAdmin(adminDetailsResponse)

	if err != nil {
		return domain.TokenAdmin{}, err
	}

	return domain.TokenAdmin{
		Admin: adminDetailsResponse,
		Token: tokenString,
	}, nil

}
func (ad *AdminUseCase) GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error) {

	userDetails, err := ad.adminRepository.GetUsers(page, count)
	if err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil

}
func (ad *AdminUseCase) BlockUser(userID int) error {
	user, err := ad.adminRepository.GetUserById(userID)
	if err != nil {
		return err
	}
	if user.Blocked {
		return errors.New("already blocked")
	} else {
		user.Blocked = true
	}
	err = ad.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}

	return nil

}
func (ad *AdminUseCase) UnBlockUser(userID int) error {
	user, err := ad.adminRepository.GetUserById(userID)
	if err != nil {
		return err
	}
	if user.Blocked {
		user.Blocked = false
	} else {
		return errors.New("already unblocked")
	}
	err = ad.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}

	return nil
}
func (ad *AdminUseCase) Dashboard() (models.DashBoardTotal, error) {
	userdetails, err := ad.adminRepository.UserDetailsDashboard()
	if err != nil {
		return models.DashBoardTotal{}, err
	}
	productdetails, err := ad.adminRepository.ProductDetailsDashboard()
	if err != nil {
		return models.DashBoardTotal{}, err
	}
	orderDetails, err := ad.adminRepository.OrderDetailsDashboard()
	if err != nil {
		return models.DashBoardTotal{}, err
	}
	amountDetails, err := ad.adminRepository.AmountDetails()
	if err != nil {
		return models.DashBoardTotal{}, err
	}
	totalRevenue, err := ad.adminRepository.TotalRevenue()
	if err != nil {
		return models.DashBoardTotal{}, err
	}

	return models.DashBoardTotal{
		Users:    userdetails,
		Products: productdetails,
		Orders:   orderDetails,
		Amount:   amountDetails,
		Revenue:  totalRevenue,
	}, nil
}

func (ad *AdminUseCase) FilterSalesReport(timePeriod string) (models.SalesReport, error) {

	if timePeriod == "" {
		err := errors.New("field cannot be empty")
		return models.SalesReport{}, err
	}
	if timePeriod != "week" && timePeriod != "month" && timePeriod != "year" {
		err := errors.New("invalid time period, available options : week, month & year")
		return models.SalesReport{}, err
	}
	startTime, endTime := helper.GetTimeFromPeriod(timePeriod)
	saleReport, err := ad.adminRepository.FilterSalesReport(startTime, endTime)
	if err != nil {
		return models.SalesReport{}, err
	}
	return saleReport, nil

}
func (ad *AdminUseCase) ExecuteSalesReportByDate(startDate, endDate string) (models.SalesReport, error) {
	parsedStartDate, err := time.Parse("02-01-2006", startDate)
	if err != nil {
		err := errors.New("enter in correct format")
		return models.SalesReport{}, err
	}
	isValid := !parsedStartDate.IsZero()
	if !isValid {
		err := errors.New("not in correct format")
		return models.SalesReport{}, err
	}
	parsedEndDate, err := time.Parse("02-01-2006", endDate)
	if err != nil {
		err := errors.New("enter the datas in correct format")
		return models.SalesReport{}, err
	}
	isValid = !parsedEndDate.IsZero()
	if !isValid {
		err := errors.New("enter the datas in correct format")
		return models.SalesReport{}, err
	}
	if parsedStartDate.After(parsedEndDate) {
		err := errors.New(" enter in correct format")

		return models.SalesReport{}, err
	}
	orders, err := ad.adminRepository.FilterSalesReport(parsedStartDate, parsedEndDate)
	if err != nil {
		return models.SalesReport{}, errors.New("error")
	}
	return orders, nil
}
