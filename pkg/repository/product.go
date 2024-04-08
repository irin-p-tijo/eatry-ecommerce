package repository

import (
	interfaces "eatry/pkg/repository/interfaces"
	"eatry/pkg/utils/models"
	"errors"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &ProductRepository{
		DB: DB,
	}
}

func (pro *ProductRepository) AddProduct(addproduct models.AddProduct) (models.ProductResponse, error) {
	var products models.ProductResponse

	err := pro.DB.Raw("insert into products (category_id,name,quantity,stock,price) values (?,?,?,?,?) returning id,category_id,name,quantity,stock,price", addproduct.CategoryID, addproduct.Name, addproduct.Quantity, addproduct.Stock, addproduct.Price).Scan(&products).Error
	if err != nil {
		return models.ProductResponse{}, err
	}
	return products, err
}
func (pro *ProductRepository) DeleteProduct(productID int) error {
	err := pro.DB.Exec("delete from products where id=?", productID)
	if err.RowsAffected < 1 {
		return errors.New("no records with that ID exist")
	}

	return nil
}
func (pro *ProductRepository) UpdateProduct(Products models.ProductResponse, productID int) (models.ProductResponse, error) {
	var updatedproduct models.ProductResponse
	err := pro.DB.Raw("update products set category_id=?,name=?,quantity=?,stock=?,price=? where id=? ", Products.CategoryID, Products.Name, Products.Quantity, Products.Stock, Products.Price, productID).Scan(&updatedproduct).Error
	if err != nil {
		return models.ProductResponse{}, err
	}
	return updatedproduct, nil
}
func (pro *ProductRepository) ListProducts(page int, count int) ([]models.AddProduct, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var products []models.AddProduct
	err := pro.DB.Raw("select id,category_id,name,quantity,stock,price from products limit ? offset ?", count, offset).Scan(&products).Error
	if err != nil {
		return []models.AddProduct{}, err
	}
	return products, nil
}
func (pro *ProductRepository) CheckProduct(productID int) (bool, error) {
	var k int
	err := pro.DB.Raw("select count(*) from products where id=?", productID).Scan(&k).Error
	if err != nil {
		return false, err
	}
	if k == 0 {
		return false, err
	}

	return true, err
}
func (pro *ProductRepository) CheckStock(productID int) (int, error) {
	var k int
	err := pro.DB.Raw("select stock from products where id=?", productID).Scan(&k).Error
	if err != nil {
		return 0, err
	}
	return k, nil
}
func (pro *ProductRepository) GetPriceofProduct(productID int) (float64, error) {

	var price float64
	err := pro.DB.Raw("select price from products where id=?", productID).Scan(&price).Error
	if err != nil {
		return 0.0, err
	}
	return price, nil
}
func (pro *ProductRepository) CheckValidityOfCategory(data map[string]int) error {

	for _, id := range data {
		var count int
		err := pro.DB.Raw("select count(*) from categories where id = ?", id).Scan(&count).Error
		if err != nil {
			return err
		}

		if count < 1 {
			return errors.New("category does not exist")
		}
	}
	return nil
}

func (pro *ProductRepository) GetProductFromCategory(id int) ([]models.ProductBrief, error) {

	var product []models.ProductBrief
	err := pro.DB.Raw(`
		SELECT *
		FROM products
		JOIN categories ON products.category_id = categories.id
		 where categories.id = ?
	`, id).Scan(&product).Error

	if err != nil {
		return []models.ProductBrief{}, err
	}
	return product, nil
}
func (pro *ProductRepository) GetQuantityFromProductID(id int) (int, error) {

	var quantity int
	err := pro.DB.Raw("select quantity from products where id = ?", id).Scan(&quantity).Error
	if err != nil {
		return 0.0, err
	}

	return quantity, nil

}
func (p *ProductRepository) SearchItemBasedOnPrefix(prefix string) ([]models.ProductBrief, int, error) {

	// find length of prefix
	lengthOfPrefix := len(prefix)
	var productsBrief []models.ProductBrief
	err := p.DB.Raw(`
		SELECT products.id, products.name,products.price,products.quantity
		FROM products
		JOIN categories ON products.category_id = categories.id
	`).Scan(&productsBrief).Error

	if err != nil {
		return nil, 0, err
	}

	return productsBrief, lengthOfPrefix, nil

}
