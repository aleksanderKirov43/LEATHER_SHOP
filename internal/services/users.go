package services

import "leather-shop/internal/models"

// Определение методов, которые должен реализовать сервис
type User interface {
	GetUser(id int) (*models.User, error)
	GetUsers() ([]*models.User, error)
	CreateUser(user *models.User) error
	DeleteUser(id int) error
	EditUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
	CheckPassword(password, hashedPassword string) bool
}
