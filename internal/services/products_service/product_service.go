package products_service

import (
	"context"
	"encoding/json"

	"leather-shop/internal/models"
	"leather-shop/internal/repository"
	"leather-shop/internal/services"
)

type productService struct {
	productRepository repository.Products
}

func (pr *productService) SetImages(image string) (string, error) {
	bytes, err := json.Marshal(image)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (pr *productService) GetImages(image string) ([]string, error) {
	var images []string
	err := json.Unmarshal([]byte(image), &images)
	if err != nil {
		return nil, err
	}
	return images, nil
}

// Создаём новый экземпляр сервиса для товаров
func New(productRepository repository.Products) services.Products {
	return &productService{
		productRepository: productRepository,
	}
}

func (ps *productService) GetProduct(ctx context.Context, id int) (*models.Products, error) {
	product, err := ps.productRepository.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *productService) GetProducts(ctx context.Context) ([]*models.Products, error) {
	products, err := ps.productRepository.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *productService) CreateProduct(ctx context.Context, product *models.Products) error {
	return ps.productRepository.CreateProduct(ctx, product)
}

func (ps *productService) DeleteProduct(ctx context.Context, id int) error {
	return ps.productRepository.DeleteProduct(ctx, id)
}

func (ps *productService) EditProduct(ctx context.Context, product *models.Products) error {
	return ps.productRepository.EditProduct(ctx, product)
}
