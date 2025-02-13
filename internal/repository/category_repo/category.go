package category_repo

import (
	"errors"
	"gorm.io/gorm"
	"log"

	"leather-shop/internal/models"
	"leather-shop/internal/repository"
)

type categoryRepository struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) repository.Category {
	return &categoryRepository{
		DB: DB,
	}
}

func (cr *categoryRepository) GetCategory(id int) (*models.Category, error) {
	var category models.Category
	//query := "SELECT * FROM product_category WHERE id = $1"
	err := cr.DB.Table("product_category").Where("id = ?", id).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Категория не найдена")
		}
		log.Println(err)
		return nil, errors.New("Ошибка запроса в базу")
	}
	return &category, nil
}

func (cr *categoryRepository) GetCategories() ([]*models.Category, error) {
	var categories []*models.Category
	err := cr.DB.Table("product_category").Find(&categories).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Категории не найдены")
		}
		log.Println(err)
		return nil, errors.New("Ошибка запроса в базу")
	}
	return categories, nil
}

func (cr *categoryRepository) CreateCategory(category *models.Category) error {
	var err error
	err = cr.DB.Table("product_category").Create(category).Error
	if err != nil {
		log.Println(err)
		return errors.New("Ошибка создания категории")
	}
	return nil
}

func (cr *categoryRepository) DeleteCategory(id int) error {
	err := cr.DB.Table("product_category").Where("id = ?", id).Delete(&models.Category{}).Error
	if err != nil {
		log.Println(err)
		return errors.New("Ошибка удаления категории")
	}
	return err
}

func (cr *categoryRepository) EditCategory(category *models.Category) error {
	err := cr.DB.Table("product_category").Where("id = ?", category.Id).Updates(category).Error
	if err != nil {
		log.Println(err)
		return errors.New("Ошибка редактирования категории")
	}
	return nil
}
