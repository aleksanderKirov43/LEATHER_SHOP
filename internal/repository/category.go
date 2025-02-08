package repository

import "leather-shop/internal/models"

type Category interface {
	GetCategory(id int) (*models.Category, error)
	GetCategories() ([]*models.Category, error)
	CreateCategory(category *models.Category) error
	DeleteCategory(id int) error
	EditCategory(category *models.Category) error
}
