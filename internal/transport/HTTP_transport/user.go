package HTTP_transport

import "github.com/gin-gonic/gin"

// Определяем методы, которые должен реализовывать контроллер пользователей
type UserController interface {
	GetUser(c *gin.Context)
	GetUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	EditUser(c *gin.Context)
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
}
