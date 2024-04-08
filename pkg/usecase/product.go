package usecase

import (
	interfaces "eatry/pkg/repository/interfaces"
	services "eatry/pkg/usecase/interfaces"
	"eatry/pkg/utils/models"
	"fmt"
	"strings"
)

type ProductUseCase struct {
	productRepository interfaces.ProductRepository
}

func NewProductUseCase(usecase interfaces.ProductRepository) services.ProductUseCase {
	return &ProductUseCase{
		productRepository: usecase,
	}
}
func (pr *ProductUseCase) AddProduct(product models.AddProduct) (models.ProductResponse, error) {

	products, err := pr.productRepository.AddProduct(product)
	if err != nil {
		return models.ProductResponse{}, err
	}
	return products, nil
}
func (pr *ProductUseCase) DeleteProduct(ProductID int) error {
	err := pr.productRepository.DeleteProduct(ProductID)
	if err != nil {
		return err
	}
	return nil

}
func (pr *ProductUseCase) UpdateProduct(Products models.ProductResponse, productID int) (models.ProductResponse, error) {
	updatedproduct, err := pr.productRepository.UpdateProduct(Products, productID)
	if err != nil {
		return models.ProductResponse{}, err
	}
	return updatedproduct, nil
}
func (pr *ProductUseCase) ListProducts(page int, count int) ([]models.AddProduct, error) {
	products, err := pr.productRepository.ListProducts(page, count)
	if err != nil {
		return []models.AddProduct{}, err
	}
	return products, nil
}
func (pr *ProductUseCase) FilterCategory(data map[string]int) ([]models.ProductBrief, error) {

	err := pr.productRepository.CheckValidityOfCategory(data)
	if err != nil {
		return []models.ProductBrief{}, err
	}

	var productFromCategory []models.ProductBrief
	for _, id := range data {

		product, err := pr.productRepository.GetProductFromCategory(id)
		if err != nil {
			return []models.ProductBrief{}, err
		}
		for _, product := range product {

			quantity, err := pr.productRepository.GetQuantityFromProductID(product.ID)
			if err != nil {
				return []models.ProductBrief{}, err
			}
			if quantity == 0 {
				product.ProductStatus = "out of stock"
			} else {
				product.ProductStatus = "in stock"
			}
			if product.ID != 0 {
				productFromCategory = append(productFromCategory, product)
			}
		}

		// if a product exist for that genre. Then only append it

	}
	return productFromCategory, nil

}
func (pr *ProductUseCase) SearchItemBasedOnPrefix(prefix string) ([]models.ProductBrief, error) {

	productsBrief, lengthOfPrefix, err := pr.productRepository.SearchItemBasedOnPrefix(prefix)
	if err != nil {
		return []models.ProductBrief{}, err
	}

	// Create a slice to add the products which have the given prefix
	var filteredProductBrief []models.ProductBrief
	for _, p := range productsBrief {
		length := len(p.Name)
		if length >= lengthOfPrefix {
			moviePrefix := p.Name[:lengthOfPrefix]
			if strings.EqualFold(prefix, moviePrefix) {
				filteredProductBrief = append(filteredProductBrief, p)
			}
		}
	}

	for i := range filteredProductBrief {
		fmt.Println("the code reached here")
		p := &filteredProductBrief[i]
		if p.Quantity == 0 {
			p.ProductStatus = "out of stock"
		} else {
			p.ProductStatus = "in stock"
		}
	}

	return filteredProductBrief, nil
}
