package category_service

import (
	"context"

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

func (cs *categoryService) GetCategory(ctx context.Context, id int) (*models.Category, error) {
	category, err := cs.categoryRepository.GetCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (cs *categoryService) GetCategories(ctx context.Context) ([]*models.Category, error) {
	categories, err := cs.categoryRepository.GetCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (cs *categoryService) CreateCategory(ctx context.Context, category *models.Category) error {
	return cs.categoryRepository.CreateCategory(ctx, category)
}

func (cs *categoryService) DeleteCategory(ctx context.Context, id int) error {
	return cs.categoryRepository.DeleteCategory(ctx, id)
}

func (cs *categoryService) EditCategory(ctx context.Context, category *models.Category) error {
	return cs.categoryRepository.EditCategory(ctx, category)
}
