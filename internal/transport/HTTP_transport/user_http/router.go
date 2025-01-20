package user_http

import (
	"github.com/gin-gonic/gin"

	"leather-shop/internal/transport/HTTP_transport"
	"leather-shop/internal/transport/middlewares"
)

// Создаём группы маршрутов для пользователей
func NewRouter(engine *gin.RouterGroup, controller HTTP_transport.UserController) {
	// Группа для маршрутов авторизации
	authGroup := engine.Group("/auth")
	{
		authGroup.POST("/login", controller.Login)
		authGroup.POST("/refresh", controller.RefreshToken)
	}

	// Группа для маршрутов создания пользователей без авторизации
	engine.POST("/users", controller.CreateUser)

	// Группа маршрутов для пользователей, требующих авторизации
	usersGroup := engine.Group("/users")
	usersGroup.Use(middlewares.JwtMiddleware())
	{
		usersGroup.GET("/:id", controller.GetUser)
		usersGroup.GET("", controller.GetUsers)
		usersGroup.DELETE("/:id", controller.DeleteUser)
		usersGroup.PUT("/:id", controller.EditUser)
	}
}
