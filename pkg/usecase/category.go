package usecase

import (
	"eatry/pkg/domain"
	interfaces "eatry/pkg/repository/interfaces"
	services "eatry/pkg/usecase/interfaces"
)

type CategoryUseCase struct {
	CategoryRepository interfaces.CategoryRepository
}

func NewCategoryUseCase(usecase interfaces.CategoryRepository) services.CategoryUseCase {
	return &CategoryUseCase{
		CategoryRepository: usecase,
	}

}
func (cat *CategoryUseCase) AddCategory(category domain.Category) (domain.Category, error) {
	productresponse, err := cat.CategoryRepository.AddCategory(category)
	if err != nil {
		return domain.Category{}, err
	}
	return productresponse, nil
}
func (cat *CategoryUseCase) DeleteCategory(CategoryID int) error {

	err := cat.CategoryRepository.DeleteCategory(CategoryID)
	if err != nil {
		return err
	}
	return nil
}
func (cat *CategoryUseCase) GetCategory(page int, count int) ([]domain.Category, error) {
	category, err := cat.CategoryRepository.GetCategory(page, count)

	if err != nil {
		return []domain.Category{}, err
	}
	return category, nil
}
