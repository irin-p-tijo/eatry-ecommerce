package interfaces

import "eatry/pkg/utils/models"

type ProductRepository interface {
	AddProduct(addproduct models.AddProduct) (models.ProductResponse, error)
	DeleteProduct(productID int) error
	UpdateProduct(Products models.ProductResponse, productID int) (models.ProductResponse, error)
	CheckStock(productID int) (int, error)
	CheckProduct(productID int) (bool, error)
	GetPriceofProduct(productID int) (float64, error)
	ListProducts(page int, count int) ([]models.AddProduct, error)
	GetQuantityFromProductID(id int) (int, error)
	GetProductFromCategory(id int) ([]models.ProductBrief, error)
	CheckValidityOfCategory(data map[string]int) error
	SearchItemBasedOnPrefix(prefix string) ([]models.ProductBrief, int, error)
	// SearchItem(key string) ([]models.ProductResponse, error)
	// ShowIndividualProduct(productID int) (models.ProductResponse, error)
}
