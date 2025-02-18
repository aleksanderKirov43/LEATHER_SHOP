package category_repo

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"

	"leather-shop/internal/models"
	"leather-shop/internal/repository"
)

type categoryRepository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) repository.Category {
	return &categoryRepository{
		pool: pool,
	}
}

func (cr *categoryRepository) GetCategory(ctx context.Context, id int) (*models.Category, error) {
	var category models.Category
	query := "SELECT * FROM product_category WHERE id = $1"
	row := cr.pool.QueryRow(ctx, query, id)

	err := row.Scan(&category.Id, &category.Name, &category.Props)
	if err != nil {
		if err.Error() == "Таких столбцов не сущесвтует" {
			return nil, errors.New("Категория не найдена")
		}
		log.Println(err)
		return nil, errors.New("Ошибка запроса в базу")
	}
	return &category, nil
}

func (cr *categoryRepository) GetCategories(ctx context.Context) ([]*models.Category, error) {
	var categories []*models.Category
	query := "SELECT * FROM product_category"
	rows, err := cr.pool.Query(ctx, query)

	if err != nil {
		log.Println(err)
		return nil, errors.New("Ошибка запроса в базу")
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.Id, &category.Name, &category.Props)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Ошибка обработки записей")
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (cr *categoryRepository) CreateCategory(ctx context.Context, category *models.Category) error {
	query := "INSERT INTO product_category (name, props) VALUES ($1, $2)"
	_, err := cr.pool.Exec(ctx, query, category.Name, category.Props)
	if err != nil {
		log.Println(err)
		return errors.New("Ошибка создания категории")
	}
	return nil
}

func (cr *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	query := "DELETE FROM product_category WHERE id=$1"
	_, err := cr.pool.Exec(ctx, query, id)
	if err != nil {
		log.Println(err)
		return errors.New("Ошибка удаления категории")
	}
	return nil
}

func (cr *categoryRepository) EditCategory(ctx context.Context, category *models.Category) error {
	query := "UPDATE product_category SET name=$1, props=$2 WHERE id=$3"
	_, err := cr.pool.Exec(ctx, query, category.Name, category.Props, category.Id)
	if err != nil {
		log.Println(err)
		return errors.New("Ошибка редактирования категории")
	}
	return nil
}
