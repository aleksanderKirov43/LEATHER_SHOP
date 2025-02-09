package category_service

import (
	"leather-shop/internal/models"
	"leather-shop/internal/repository"
	"leather-shop/internal/services"
)

type categoryService struct {
	categoryRepository repository.Category
}

func New(categoryRepository repository.Category) services.Category {
	return &categoryService{
		categoryRepository: categoryRepository,
	}
}

func (cs *categoryService) GetCategory(id int) (*models.Category, error) {
	category, err := cs.categoryRepository.GetCategory(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (cs *categoryService) GetCategories() ([]*models.Category, error) {
	categories, err := cs.categoryRepository.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (cs *categoryService) CreateCategory(category *models.Category) error {
	return cs.categoryRepository.CreateCategory(category)
}

func (cs *categoryService) DeleteCategory(id int) error {
	return cs.categoryRepository.DeleteCategory(id)
}

func (cs *categoryService) EditCategory(category *models.Category) error {
	return cs.categoryRepository.EditCategory(category)
}
