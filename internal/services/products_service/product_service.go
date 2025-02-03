package products_service

import (
	"leather-shop/internal/models"
	"leather-shop/internal/repository"
)

type productService struct {
	productService repository.Products
}

func (ps *productService) GetProduct(id int) (*models.Products, error) {
	product, err := ps.productService.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *productService) GetProducts() ([]*models.Products, error) {
	products, err := ps.productService.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *productService) CreateProduct(product *models.Products) error {
	return ps.productService.CreateProduct(product)
}

func (ps *productService) DeleteProduct(id int) error {
	return ps.productService.DeleteProduct(id)
}

func (ps *productService) EditProduct(product *models.Products) error {
	return ps.productService.EditProduct(product)
}
