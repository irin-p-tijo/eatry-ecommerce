package repository

import (
	"eatry/pkg/domain"
	interfaces "eatry/pkg/repository/interfaces"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) interfaces.CategoryRepository {
	return &CategoryRepository{DB: DB}
}
func (ca *CategoryRepository) AddCategory(category domain.Category) (domain.Category, error) {
	var cat string
	err := ca.DB.Raw("insert into categories (category) values (?) returning category", category.Category).Scan(&cat).Error
	if err != nil {
		return domain.Category{}, err
	}
	var categoryresponse domain.Category

	err = ca.DB.Raw("select p.id,p.category from categories p where p.category=?", cat).Scan(&categoryresponse).Error
	if err != nil {
		return domain.Category{}, err
	}
	return categoryresponse, nil
}
func (ca *CategoryRepository) DeleteCategory(CategoryID int) error {

	err := ca.DB.Exec("delete from categories where id=?", CategoryID)

	if err.RowsAffected < 1 {
		return errors.New("the id is not existing")
	}
	return nil
}

// list category
func (ca *CategoryRepository) GetCategory(page int, count int) ([]domain.Category, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var categories []domain.Category
	if err := ca.DB.Raw("select * from categories limit ? offset ?", count, offset).Scan(&categories).Error; err != nil {
		return []domain.Category{}, err
	}
	return categories, nil
}
