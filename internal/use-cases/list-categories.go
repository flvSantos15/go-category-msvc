package usecases

import (
	"github.com/flvSantos15/go-category-msvc/internal/entities"
	"github.com/flvSantos15/go-category-msvc/internal/repositories"
)

type listCategoryUseCase struct {
	categoryRepository repositories.ICategoryRepository
}

func NewListCategoryUseCase(categoryRepository repositories.ICategoryRepository) *listCategoryUseCase {
	return &listCategoryUseCase{categoryRepository}
}

func (usecase *listCategoryUseCase) Execute() ([]*entities.Category, error) {
	categories, err := usecase.categoryRepository.List()
	if err != nil {
		return nil, err
	}

	return categories, err
}
