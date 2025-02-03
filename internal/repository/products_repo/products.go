package products_repo

import (
	"errors"
	"gorm.io/gorm"
	"leather-shop/internal/models"
	"log"
)

type productRepository struct {
	DB *gorm.DB
}

func (pr *productRepository) GetProduct(id int) (*models.Products, error) {
	var product models.Products
	err := pr.DB.Table("products").Where("id = ?", id).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Товар не найден")
		}
		log.Println(err)
		return nil, errors.New("Ошибка запроса в базу")
	}
	return &product, nil
}

func (pr *productRepository) GetProducts() ([]*models.Products, error) {
	var products []*models.Products
	err := pr.DB.Table("products").Find(&products).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Товары не найдены")
		}
		log.Println(err)
		return nil, errors.New("Ошибка запроса в базу")
	}
	return products, nil
}

func (pr *productRepository) CreateProduct(product *models.Products) error {
	err := pr.DB.Table("products").Create(product).Error
	if err != nil {
		log.Println(err)
		return errors.New("Ошибка создания товара")
	}
	return nil
}

func (pr *productRepository) DeleteProduct(id int) error {
	err := pr.DB.Table("products").Where("id = ?", id).Delete(&models.Products{}).Error
	if err != nil {
		log.Println(err)
		return errors.New("Ошибка удаления товара")
	}
	return nil
}

func (pr *productRepository) EditProduct(product *models.Products) error {
	err := pr.DB.Table("products").Where("id = ?", product.Id).Updates(product).Error
	if err != nil {
		log.Println(err)
		return errors.New("Ошибка редактирования товара")
	}
	return nil
}
