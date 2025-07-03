package repositories

import "github.com/flvSantos15/go-category-msvc/internal/entities"

type InMemoryCategoryRepository struct {
	db []*entities.Category
}

func NewInMemoryCategoryRepository() *InMemoryCategoryRepository {
	return &InMemoryCategoryRepository{
		db: []*entities.Category{},
	}
}

func (repo *InMemoryCategoryRepository) Save(category *entities.Category) error {
	repo.db = append(repo.db, category)

	return nil
}

func (repo *InMemoryCategoryRepository) List() ([]*entities.Category, error) {
	return repo.db, nil
}
