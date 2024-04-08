package interfaces

import "eatry/pkg/utils/models"

type ProductUseCase interface {
	AddProduct(product models.AddProduct) (models.ProductResponse, error)
	DeleteProduct(ProductID int) error
	UpdateProduct(Products models.ProductResponse, productID int) (models.ProductResponse, error)
	ListProducts(page int, count int) ([]models.AddProduct, error)
	FilterCategory(data map[string]int) ([]models.ProductBrief, error)
	SearchItemBasedOnPrefix(prefix string) ([]models.ProductBrief, error)
	//SearchItem(prefix string) ([]models.ProductResponse, error)
	//ShowIndividualProduct(productID int) (models.ProductResponse, error)
}
