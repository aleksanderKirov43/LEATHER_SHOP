package db_connect

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"leather-shop/config"
)

func InitDB(config config.DBLeather) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow", config.Host, config.User, config.Password, config.Database, config.Port)
	// Настройка GORM
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Проверка подключения
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Ошибка получения SQL DB: %v", err)
	}

	// Настройка пула соединений
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	log.Println("Подключение к базе данных успешно установлено.")

	return DB
}
