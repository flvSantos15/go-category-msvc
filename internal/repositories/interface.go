package repositories

import "github.com/flvSantos15/go-category-msvc/internal/entities"

type ICategoryRepository interface {
	Save(entity *entities.Category) error
	List() ([]*entities.Category, error)
}
