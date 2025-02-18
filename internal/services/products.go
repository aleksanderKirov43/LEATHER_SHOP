package services

import (
	"context"

	"leather-shop/internal/models"
)

type Products interface {
	GetProduct(ctx context.Context, id int) (*models.Products, error)
	GetProducts(ctx context.Context) ([]*models.Products, error)
	CreateProduct(ctx context.Context, product *models.Products) error
	DeleteProduct(ctx context.Context, id int) error
	EditProduct(ctx context.Context, product *models.Products) error
	SetImages(image string) (string, error)
	GetImages(image string) ([]string, error)
}
