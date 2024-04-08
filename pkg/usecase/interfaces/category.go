package interfaces

import "eatry/pkg/domain"

type CategoryUseCase interface {
	AddCategory(category domain.Category) (domain.Category, error)
	DeleteCategory(CategoryID int) error
	GetCategory(page int, count int) ([]domain.Category, error)
}
