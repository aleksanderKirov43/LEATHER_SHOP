package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"leather-shop/config"
	"leather-shop/internal/repository"
	"leather-shop/internal/repository/user_repo"
	"leather-shop/internal/services"
	"leather-shop/internal/services/users_service"
	"leather-shop/internal/transport/HTTP_transport/user_http"
	"leather-shop/internal/transport/middlewares"
	"leather-shop/pkg/db_connect"
	"leather-shop/pkg/jwt"
)

type App struct {
	userRepository repository.User // Репозиторий для работы с пользователями
	userService    services.User   // Сервис для бизнес-логики пользователей
	jwtHelper      jwt.Helper      // Помощник для работы с JWT

	cfg *config.Config
}

func New(config *config.Config) *App {
	DB := db_connect.InitDB(config.DBLeather) //Инициализация базы данных

	// Инициализация репозитория и сервиса для пользователей
	userRepository := user_repo.New(DB)
	userService := users_service.New(userRepository)

	// Инициализация jwtHelper
	jwtHelper := jwt.NewHelper(config.Jwt.Secret, config.Jwt.AccessTTL, config.Jwt.RefreshTTL)

	return &App{

		userRepository: userRepository,
		userService:    userService,
		jwtHelper:      jwtHelper,
		cfg:            config,
	}
}

func (a *App) Run() {
	a.startHttp()
}

// Создание экземпляра Gin и добавляем middleware для CORS.
func (a *App) startHttp() {
	engine := gin.Default()
	engine.Use(middlewares.CORSMiddleware())

	// Создаём группу маршрутов для /api
	apiGroup := engine.Group("/api")
	// Инициализируем контроллер для пользователей
	userController := user_http.New(a.userService, a.jwtHelper)

	user_http.NewRouter(apiGroup, userController)

	if err := engine.Run(fmt.Sprintf(":%s", a.cfg.ApplicationPort)); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("HTTP Сервер прослушивает порт :%s \r\n", a.cfg.DBLeather.Port)
}
