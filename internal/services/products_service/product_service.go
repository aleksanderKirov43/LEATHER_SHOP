package products_service

import (
	"leather-shop/internal/models"
	"leather-shop/internal/repository"
	"leather-shop/internal/services"

	"encoding/json"
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

func (ps *productService) GetProduct(id int) (*models.Products, error) {
	product, err := ps.productRepository.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *productService) GetProducts() ([]*models.Products, error) {
	products, err := ps.productRepository.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *productService) CreateProduct(product *models.Products) error {
	return ps.productRepository.CreateProduct(product)
}

func (ps *productService) DeleteProduct(id int) error {
	return ps.productRepository.DeleteProduct(id)
}

func (ps *productService) EditProduct(product *models.Products) error {
	return ps.productRepository.EditProduct(product)
}
