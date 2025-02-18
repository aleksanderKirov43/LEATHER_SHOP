package services

import (
	"context"

	"leather-shop/internal/models"
)

// Определение методов, которые должен реализовать сервис
type User interface {
	GetUser(ctx context.Context, id int) (*models.User, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (string, *models.User, error)
	DeleteUser(ctx context.Context, id int) error
	EditUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	CheckPassword(password, hashedPassword string) bool
}
