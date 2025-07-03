package usecases

import (
	"github.com/flvSantos15/go-category-msvc/internal/entities"
	"github.com/flvSantos15/go-category-msvc/internal/repositories"
)

type createCategoryUseCase struct {
	categoryRepository repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(categoryRepository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{categoryRepository}
}

func (usecase *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)
	if err != nil {
		return err
	}

	err = usecase.categoryRepository.Save(category)
	if err != nil {
		return err
	}

	return nil
}