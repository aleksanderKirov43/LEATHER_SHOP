package user_repo

import (
	"errors"
	"gorm.io/gorm"
	"log"

	"leather-shop/internal/models"
	"leather-shop/internal/repository"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository struct {
	DB *gorm.DB
}

func (ur *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := ur.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Создание нового экземпляра репозитория пользователей
func New(DB *gorm.DB) repository.User {
	return &userRepository{
		DB: DB,
	}
}

// Методы репозитория для пользователя
func (ur *userRepository) GetUser(id int) (*models.User, error) {
	var user models.User
	err := ur.DB.Table("users").Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Пользователь не найден")
		}
		log.Println(err)
		return nil, errors.New("Ошибка запроса в базу")
	}
	return &user, nil
}

func (ur *userRepository) GetUsers() ([]*models.User, error) {
	var users []*models.User
	err := ur.DB.Table("users").Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Пользователи не найдены")
		}
		log.Println(err)
		return nil, errors.New("Ошибка запроса в базу")
	}
	return users, nil
}

func (ur *userRepository) CreateUser(user *models.User) error {
	err := ur.DB.Table("users").Create(user).Error
	if err != nil {
		log.Println(err)
		return errors.New("Ошибка создания пользователя")
	}
	return nil
}

func (ur *userRepository) DeleteUser(id int) error {
	err := ur.DB.Table("users").Where("id = ?", id).Delete(&models.User{}).Error
	if err != nil {
		log.Println(err)
		return errors.New("Ошибка удаления пользователя")
	}
	return nil
}

func (ur *userRepository) EditUser(user *models.User) error {
	err := ur.DB.Table("users").Where("id = ?", user.Id).Updates(user).Error
	if err != nil {
		log.Println(err)
		return errors.New("Ошибка редактирования данных пользователя")
	}
	return nil
}

func (ur *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := ur.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
