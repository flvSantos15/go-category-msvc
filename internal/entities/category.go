package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCategory(name string) (*Category, error) {
	category := &Category{
		ID:   uuid.New().String(),
		Name: name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	return category, category.IsValid()
}

func (c *Category) IsValid() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if len(c.Name) < 5 {
		return errors.New("name must be at least 5 characters long")
	}
	return nil
}

func (c *Category) Update(name string) error {
	c.Name = name
	c.UpdatedAt = time.Now()
	return c.IsValid()
}

