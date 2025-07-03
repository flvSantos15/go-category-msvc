package usecases

import (
	"log"

	"github.com/flvSantos15/go-category-msvc/internal/entities"
)

type createCategoryUseCase struct {
	// db
}

func NewCreateCategoryUseCase() *createCategoryUseCase {
	return &createCategoryUseCase{}
}

func (usecase *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)
	if err != nil {
		return err
	}

	// save to db
	log.Println(category)

	return nil
}