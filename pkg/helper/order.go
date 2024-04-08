package helper

import (
	"eatry/pkg/domain"
	"eatry/pkg/utils/models"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func CopyOrderDetails(orderDetails domain.Order, orderbody models.OrderIncoming) domain.Order {

	id := uuid.New().ID()
	str := strconv.Itoa(int(id))
	orderDetails.OrderId = str[:8]
	orderDetails.AddressID = orderbody.AddressID
	orderDetails.PaymentMethodID = orderbody.PaymentID
	orderDetails.UserID = int(orderbody.UserID)
	orderDetails.Approval = false
	orderDetails.ShipmentStatus = "processing"
	orderDetails.PaymentStatus = "not paid"

	return orderDetails

}
func GetTimeFromPeriod(timePeriod string) (time.Time, time.Time) {

	endDate := time.Now()

	if timePeriod == "week" {
		startDate := endDate.AddDate(0, 0, -6)
		return startDate, endDate
	}

	if timePeriod == "month" {
		startDate := endDate.AddDate(0, -1, 0)
		return startDate, endDate
	}

	if timePeriod == "year" {
		startDate := endDate.AddDate(0, -1, 0)
		return startDate, endDate
	}

	return endDate.AddDate(0, 0, -6), endDate

}
