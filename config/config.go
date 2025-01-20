package config

import (
	"log"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	ApplicationPort string `env:"LEATHER_APP_PORT"`
	DBLeather       DBLeather
	Jwt             JWTConfig
}
type DBLeather struct {
	Host     string `env:"LEATHER_DB_HOST"`
	Port     string `env:"LEATHER_DB_PORT"`
	User     string `env:"LEATHER_DB_USER"`
	Password string `env:"LEATHER_DB_PASSWORD"`
	Database string `env:"LEATHER_DB_DATABASE"`
}

type JWTConfig struct {
	Secret     string `env:"JWT_SECRET"`
	AccessTTL  int    `env:"JWT_ACCESS_TTL"`
	RefreshTTL int    `env:"JWT_REFRESH_TTL"`
}

const localConfigPath = "./config/env/.env"
const deployConfigPath = "./.env"

var instance *Config
var once sync.Once

func GetConfig() *Config {
	deploy := os.Getenv("DEPLOY") == "true"
	once.Do(func() {
		if !deploy {
			err := godotenv.Load(localConfigPath)
			if err != nil {
				panic("Конфигурационный local .env файл не найден")
			}
		} else {
			err := godotenv.Load(deployConfigPath)
			if err != nil {
				panic("Конфигурационный deploy .env файл не найден")
			}
		}

		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			panic(err)
		}
	})

	return instance
}
