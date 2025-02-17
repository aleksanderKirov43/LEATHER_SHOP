package db_connect

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	"leather-shop/config"
)

func InitDB(cfg config.DBLeather) *pgxpool.Pool {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow", cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Не удается разобрать DSN: %v\n", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Не удается подключиться к базе данных: %v", err)
	}

	// Настройка пула соединений
	pool.Config().MaxConns = 100
	pool.Config().MaxConnIdleTime = 10 * time.Minute
	pool.Config().MaxConnLifetime = 30 * time.Minute

	log.Println("Подключение к базе данных успешно установлено.")

	return pool
}
