package users_service

import (
	"context"
	"golang.org/x/crypto/bcrypt"

	"leather-shop/internal/models"
	"leather-shop/internal/repository"
	"leather-shop/internal/services"
)

// Создаём структуру для хранения экземпляра репозитория пользователей
type usersService struct {
	userRepository repository.User
}

type UserService struct {
	userRepo repository.User
}

func (us *UserService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return us.userRepo.GetUserByUsername(ctx, username)
}
func (us *UserService) CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// Создаём новый экземпляр сервиса пользователей
func New(userRepository repository.User) services.User {
	return &usersService{
		userRepository: userRepository,
	}
}

// Реализация метода для получения пользователя по имени пользователя
func (us *usersService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return us.userRepository.GetUserByUsername(ctx, username)
}

// Реализация метода для проверки пароля
func (us *usersService) CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// Методы сервиса
func (us *usersService) GetUser(ctx context.Context, id int) (*models.User, error) {
	user, err := us.userRepository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *usersService) GetUsers(ctx context.Context) ([]*models.User, error) {
	users, err := us.userRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *usersService) CreateUser(ctx context.Context, user *models.User) (string, *models.User, error) {
	// Хэширование пароля
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return "", nil, err
	}
	user.Password = hashedPassword

	message, createdUser, err := us.userRepository.CreateUser(ctx, user)
	if err != nil {
		return "", nil, err
	}
	return message, createdUser, nil
}

func (us *usersService) DeleteUser(ctx context.Context, id int) error {
	return us.userRepository.DeleteUser(ctx, id)
}

func (us *usersService) EditUser(ctx context.Context, user *models.User) error {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return us.userRepository.EditUser(ctx, user)
}

// Реализация метода для хеширования пароля
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
