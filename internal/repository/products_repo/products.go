package products_repo

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"

	"leather-shop/internal/models"
	"leather-shop/internal/repository"
)

type productRepository struct {
	pool *pgxpool.Pool
}

// Создание нового репозитория для товаров
func New(pool *pgxpool.Pool) repository.Products {
	return &productRepository{
		pool: pool,
	}
}

func (pr *productRepository) GetProduct(ctx context.Context, id int) (*models.Products, error) {
	var product models.Products
	query := "SELECT * FROM products WHERE id = $1"
	row := pr.pool.QueryRow(ctx, query, id)

	err := row.Scan(&product.Id, &product.Name, &product.Description, &product.Quantity, &product.Image, &product.Sale, &product.Price, &product.Status, &product.Category, &product.Property)
	if err != nil {
		if err.Error() == "Таких столбцов не сущесвтует" {
			return nil, errors.New("Товар не найден")
		}
		log.Println(err)
		return nil, errors.New("Ошибка запроса в базу")
	}
	return &product, nil
}

func (pr *productRepository) GetProducts(ctx context.Context) ([]*models.Products, error) {
	var products []*models.Products
	query := "SELECT * FROM products"
	rows, err := pr.pool.Query(ctx, query)

	if err != nil {
		log.Println(err)
		return nil, errors.New("Ошибка запроса в базу")
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Products
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Quantity, &product.Image, &product.Sale, &product.Price, &product.Status, &product.Category, &product.Property)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Ошибка обработки записей")
		}
		products = append(products, &product)
	}

	return products, nil
}

func (pr *productRepository) CreateProduct(ctx context.Context, product *models.Products) error {
	query := "INSERT INTO products (name, description, quantity, image, sale, price, status, category, property) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	_, err := pr.pool.Exec(ctx, query, product.Name, product.Description, product.Quantity, product.Image, product.Sale, product.Price, product.Status, product.Category, product.Property)
	if err != nil {
		log.Println(err)
		return errors.New("Ошибка создания товара")
	}
	return nil
}

func (pr *productRepository) DeleteProduct(ctx context.Context, id int) error {
	query := "DELETE FROM products WHERE id=$1"
	_, err := pr.pool.Exec(ctx, query, id)
	if err != nil {
		log.Println(err)
		return errors.New("Ошибка удаления товара")
	}
	return nil
}

func (pr *productRepository) EditProduct(ctx context.Context, product *models.Products) error {
	query := "UPDATE products SET name=$1, description=$2, quantity=$3, image=$4, sale=$5, price=$6, status=$7, category=$8, property=$9 WHERE id=$10"
	_, err := pr.pool.Exec(ctx, query, product.Name, product.Description, product.Quantity, product.Image, product.Sale, product.Price, product.Status, product.Category, product.Property, product.Id)
	if err != nil {
		log.Println(err)
		return errors.New("Ошибка редактирования товара")
	}
	return nil
}
