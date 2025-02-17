package user_http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"leather-shop/internal/models"
	"leather-shop/internal/services"
	"leather-shop/internal/transport/HTTP_transport"
	"leather-shop/internal/transport/middlewares"
	"leather-shop/pkg/jwt"
)

type userController struct {
	usersService services.User
	tokenHandler jwt.TokenHandler
}

// Создаём новый экземпляр колнтроллера пользователей
func New(usersService services.User, tokenHandler jwt.TokenHandler) HTTP_transport.UserController {
	return &userController{
		usersService: usersService,
		tokenHandler: tokenHandler,
	}
}

func (uc *userController) Login(ctx *gin.Context) {
	var auth struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&auth); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Не корректный запрос"})
		return
	}

	user, err := uc.usersService.GetUserByUsername(auth.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Ошибка сервера!"})
		return
	}

	if ok := uc.usersService.CheckPassword(auth.Password, user.Password); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Ошибка": "Неверный пароль"})
		return
	}

	ctx.Set("userId", user.Id)
	ctx.Set("username", user.Username)

	middlewares.GenerateTokenMiddleware(uc.tokenHandler)(ctx)

	accessToken := ctx.Value("accessToken").(string)
	refreshToken := ctx.Value("refreshToken").(string)

	ctx.JSON(http.StatusOK, gin.H{"Access токен": accessToken, "Refresh токен": refreshToken})
}

// Перенести в мидлвар
func (uc *userController) RefreshToken(ctx *gin.Context) {
	fmt.Println(ctx)
	jwtPayload, jwtPayloadErr := middlewares.GetJWTPayload(ctx)
	if jwtPayloadErr != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Ошибка": jwtPayloadErr.Error()})
		return
	}

	ctx.Set("userId", jwtPayload.Id)
	ctx.Set("username", jwtPayload.Login)

	middlewares.GenerateTokenMiddleware(uc.tokenHandler)(ctx)

	accessToken := ctx.Value("accessToken").(string)

	ctx.JSON(http.StatusOK, gin.H{"Аккес токен": accessToken})
}

// Получаем пользователя по ID из параметра запроса
func (uc *userController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := strconv.Atoi(id) //Преобразуем строковый id в целое число
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}
	user, err := uc.usersService.GetUser(userId) // вызов метода сервиса GetUsers
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// Получаем всех пользователей, вызывая метода сервиса GetUsers
func (uc *userController) GetUsers(ctx *gin.Context) {
	users, err := uc.usersService.GetUsers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// Создаем нового пользователя
func (uc *userController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return
	}

	message, createdUser, err := uc.usersService.CreateUser(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": message,
		"user":    createdUser,
	})
}

// Метод для удаления пользователя
func (uc *userController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}
	if err := uc.usersService.DeleteUser(userId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно удалён"})
}

// Метод для редактирования данных пользователя
func (uc *userController) EditUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}
	user.Id = userId

	if err := uc.usersService.EditUser(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
