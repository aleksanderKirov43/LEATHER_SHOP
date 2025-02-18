package app

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"leather-shop/config"
	"leather-shop/internal/repository"
	"leather-shop/internal/repository/category_repo"
	"leather-shop/internal/repository/products_repo"
	"leather-shop/internal/repository/user_repo"
	"leather-shop/internal/services"
	"leather-shop/internal/services/category_service"
	"leather-shop/internal/services/products_service"
	"leather-shop/internal/services/users_service"
	"leather-shop/internal/transport/HTTP_transport/category_http"
	"leather-shop/internal/transport/HTTP_transport/product_http"
	"leather-shop/internal/transport/HTTP_transport/user_http"
	"leather-shop/internal/transport/middlewares"
	"leather-shop/pkg/db_connect"
	"leather-shop/pkg/jwt"
)

type App struct {
	userRepository repository.User  // Репозиторий для работы с пользователями
	userService    services.User    // Сервис для бизнес-логики пользователей
	tokenHandler   jwt.TokenHandler // Помощник для работы с JWT

	productRepository repository.Products //Репозиторий для товаров
	productService    services.Products   // Сервис бизес-логики для товаров
	cfg               *config.Config

	categoryRepository repository.Category
	categoryService    services.Category
}

func New(ctx context.Context, config *config.Config) *App {
	DB := db_connect.InitDB(config.DBLeather) //Инициализация базы данных

	// Инициализация репозитория и сервиса для пользователей
	userRepository := user_repo.New(DB)
	userService := users_service.New(userRepository)

	productRepository := products_repo.New(DB)
	productService := products_service.New(productRepository)

	categoryRepository := category_repo.New(DB)
	categoryService := category_service.New(categoryRepository)

	// Инициализация jwtHelper
	tokenHandler := jwt.NewHelper(config.Jwt.Secret, config.Jwt.AccessTTL, config.Jwt.RefreshTTL)

	return &App{

		userRepository: userRepository,
		userService:    userService,
		tokenHandler:   tokenHandler,

		cfg: config,

		productRepository: productRepository,
		productService:    productService,

		categoryRepository: categoryRepository,
		categoryService:    categoryService,
	}
}

func (a *App) Run(ctx context.Context) {
	a.startHttp(ctx)
}

// Создание экземпляра Gin и добавляем middleware для CORS.
func (a *App) startHttp(ctx context.Context) {
	engine := gin.Default()
	engine.Use(middlewares.CORSMiddleware())

	// Создаём группу маршрутов для /api
	apiGroup := engine.Group("/api")
	// Инициализируем контроллер для пользователей
	userController := user_http.New(a.userService, a.tokenHandler)
	user_http.NewRouter(apiGroup, userController)

	productGroup := engine.Group("/products")
	// Инициализируем контроллер для товаров
	productController := product_http.New(a.productService)
	product_http.NewRouterProduct(productGroup, productController)

	categoryGroup := engine.Group("/category")
	categoryController := category_http.New(a.categoryService)
	category_http.NewRouterCategory(categoryGroup, categoryController)

	if err := engine.Run(fmt.Sprintf(":%s", a.cfg.ApplicationPort)); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("HTTP Сервер прослушивает порт :%s \r\n", a.cfg.DBLeather.Port)
}
