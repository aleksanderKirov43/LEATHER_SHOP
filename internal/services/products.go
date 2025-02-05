package services

import "leather-shop/internal/models"

type Products interface {
	GetProduct(id int) (*models.Products, error)
	GetProducts() ([]*models.Products, error)
	CreateProduct(product *models.Products) error
	DeleteProduct(id int) error
	EditProduct(product *models.Products) error
	SetImages(image string) (string, error)
	GetImages(image string) ([]string, error)
}
