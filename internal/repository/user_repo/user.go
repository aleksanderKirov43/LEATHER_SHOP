package user_repo

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"

	"leather-shop/internal/models"
	"leather-shop/internal/repository"
)

type userRepository struct {
	pool *pgxpool.Pool
}

// Создание нового экземпляра репозитория пользователей
func New(pool *pgxpool.Pool) repository.User {
	return &userRepository{
		pool: pool,
	}
}

func (ur *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE username = $1"
	row := ur.pool.QueryRow(context.Background(), query, username)

	err := row.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Username, &user.Type, &user.Email, &user.Password, &user.Phone, &user.Wishlist, &user.Cart)
	if err != nil {
		if err.Error() == "Таких столбцов нет в таблице" {
			return nil, errors.New("Пользователь не найден")
		}
		log.Println(err)
		return nil, errors.New("Ошибка запроса в базу")
	}
	return &user, nil
}

// Методы репозитория для пользователя
func (ur *userRepository) GetUser(id int) (*models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE id = $1"
	row := ur.pool.QueryRow(context.Background(), query, id)

	err := row.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Username, &user.Type, &user.Email, &user.Password, &user.Phone, &user.Wishlist, &user.Cart)
	if err != nil {
		if err.Error() == "Таких столбцов нет в таблице" {
			return nil, errors.New("Пользователь не найден")
		}
		log.Println(err)
		return nil, errors.New("Ошибка запроса в базу")
	}
	return &user, nil
}

func (ur *userRepository) GetUsers() ([]*models.User, error) {
	var users []*models.User
	query := "SELECT * FROM users"
	rows, err := ur.pool.Query(context.Background(), query)

	if err != nil {
		log.Println(err)
		return nil, errors.New("Ошибка запроса в базу")
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Username, &user.Type, &user.Email, &user.Password, &user.Phone, &user.Wishlist, &user.Cart)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Ошибка обработки записей")
		}
		users = append(users, &user)
	}
	return users, nil
}

func (ur *userRepository) CreateUser(user *models.User) (string, *models.User, error) {
	query := "INSERT INTO users (firstname, lastname, username, type, email, password, phone, wishlist, cart) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
	err := ur.pool.QueryRow(context.Background(), query, user.Firstname, user.Lastname, user.Username, user.Type, user.Email, user.Password, user.Phone, user.Wishlist, user.Cart).Scan(&user.Id)
	if err != nil {
		log.Println(err)
		return "", nil, errors.New("Ошибка создания пользователя")
	}
	successMessage := "Пользователь успешно создан"
	return successMessage, user, nil
}

func (ur *userRepository) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := ur.pool.Exec(context.Background(), query, id)
	if err != nil {
		log.Println(err)
		return errors.New("Ошибка удаления пользователя")
	}
	return nil
}

func (ur *userRepository) EditUser(user *models.User) error {
	query := "UPDATE users SET firstname = $1, lastname = $2, username = $3, type = $4, email = $5, password = $6, phone = $7, wishlist = $8, cart = $9 WHERE id = $10"
	_, err := ur.pool.Exec(context.Background(), query, user.Firstname, user.Lastname, user.Username, user.Type, user.Email, user.Password, user.Phone, user.Wishlist, user.Cart, user.Id)
	if err != nil {
		log.Println(err)
		return errors.New("Ошибка редактирования данных пользователя")
	}
	return nil
}
