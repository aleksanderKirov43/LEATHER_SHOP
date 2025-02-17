package repository

import (
	"context"

	"leather-shop/internal/models"
)

type Category interface {
	GetCategory(ctx context.Context, id int) (*models.Category, error)
	GetCategories(ctx context.Context) ([]*models.Category, error)
	CreateCategory(ctx context.Context, category *models.Category) error
	DeleteCategory(ctx context.Context, id int) error
	EditCategory(ctx context.Context, category *models.Category) error
}
