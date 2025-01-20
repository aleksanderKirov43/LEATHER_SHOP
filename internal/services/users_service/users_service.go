package users_service

import (
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

func NewUserServise(userRepo repository.User) UserService {
	return UserService{userRepo: userRepo}
}
func (us *UserService) GetUserByUsername(username string) (*models.User, error) {
	return us.userRepo.GetUserByUsername(username)
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
func (us *usersService) GetUserByUsername(username string) (*models.User, error) {
	return us.userRepository.GetUserByUsername(username)
}

// Реализация метода для проверки пароля
func (us *usersService) CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// Методы сервиса
func (us *usersService) GetUser(id int) (*models.User, error) {
	user, err := us.userRepository.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *usersService) GetUsers() ([]*models.User, error) {
	users, err := us.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *usersService) CreateUser(user *models.User) error {
	return us.userRepository.CreateUser(user)
}

func (us *usersService) DeleteUser(id int) error {
	return us.userRepository.DeleteUser(id)
}

func (us *usersService) EditUser(user *models.User) error {
	return us.userRepository.EditUser(user)
}
