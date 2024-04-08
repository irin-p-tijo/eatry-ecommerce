package interfaces

import (
	"eatry/pkg/domain"
	"eatry/pkg/utils/models"

	"github.com/jung-kurt/gofpdf"
)

type OrderUseCase interface {
	OrderItemsFromCart(orderFromCart models.OrderFromCart, userID int) (domain.OrderSuccessResponse, error)
	GetOrderDetails(userID int, page int, count int) ([]models.FullOrderDetails, error)
	CancelOrders(orderID string, userID int) error
	GetAllOrderDetailsForAdmin(page int) ([]models.CombinedOrderDetails, error)
	ApproveOrder(orderID string) error

	GenerateInvoice(orderID string, userID int) (*gofpdf.Fpdf, error)
}
