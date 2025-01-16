package user_http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"leather-shop/internal/models"
	"leather-shop/internal/services"
	"leather-shop/internal/transport/HTTP_transport"
	"leather-shop/internal/transport/middlewares"
	"leather-shop/pkg/jwt"
	"net/http"
	"strconv"
	"time"
)

type userController struct {
	usersService services.User
	jwtHelper    jwt.Helper
}

// Создаём новый экземпляр колнтроллера пользователей
func New(usersService services.User, jwtHelper jwt.Helper) HTTP_transport.UserController {
	return &userController{
		usersService: usersService,
		jwtHelper:    jwtHelper,
	}
}

func (uc *userController) Login(ctx *gin.Context) {
	var auth struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&auth); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Ошибка": "Ошибка ввода"})
	}

	user, err := uc.usersService.GetUserByUsername(auth.Username)
	if err != nil || !uc.usersService.CheckPassword(auth.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Ошибка": "Неверный логин или пароль"})
		return
	}
	accessToken, err := uc.jwtHelper.GenerateToken(user.Id, user.Username, time.Duration(uc.jwtHelper.AccessTTL)*time.Minute)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Ошибка": "Не удалось сгенерировать аккес токен"})
		return
	}

	refreshToken, err := uc.jwtHelper.GenerateToken(user.Id, user.Username, time.Duration(uc.jwtHelper.RefreshTTL)*time.Minute)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Ошибка": "Не удалось сгенерировать рефреш токен"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Аккес токен": accessToken, "Рефреш токен": refreshToken})
}

func (uc *userController) RefreshToken(ctx *gin.Context) {
	jwtPayload, jwtPayloadErr := middlewares.GetJWTPayload(ctx)
	if jwtPayloadErr != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Ошибка": jwtPayloadErr.Error()})
		return
	}

	accessToken, err := uc.jwtHelper.GenerateToken(jwtPayload.Id, jwtPayload.Login, time.Duration(uc.jwtHelper.AccessTTL)*time.Minute)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Ошибка": "Не удалось сгенеритровать аккес токен"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Аккес токен": accessToken})
}

// Получаем пользователя по ID из параметра запроса
func (uc *userController) GetUser(ctx *gin.Context) {
	jwtPayload, jwtPayloadErr := middlewares.GetJWTPayload(ctx)
	if jwtPayloadErr != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Ошибка": jwtPayloadErr.Error()})
		return
	}

	fmt.Println("JWT Payload:", jwtPayload)

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
	jwtPayload, jwtPayloadErr := middlewares.GetJWTPayload(ctx)
	if jwtPayloadErr != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Ошибка": jwtPayloadErr.Error()})
		return
	}

	fmt.Println("JWT Payload:", jwtPayload)

	users, err := uc.usersService.GetUsers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// // Создаём нового пользователя
//
//	func (uc *userController) CreateUser(ctx *gin.Context) {
//		jwtPayload, jwtPayloadErr := middlewares.GetJWTPayload(ctx)
//		if jwtPayloadErr != nil {
//			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Ошибка": jwtPayloadErr.Error()})
//			return
//		}
//
//		fmt.Println("JWT Payload:", jwtPayload)
//
//		var user models.User
//		if err := ctx.ShouldBindJSON(&user); err != nil {
//			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
//			return
//		}
//		if err := uc.usersService.CreateUser(&user); err != nil {
//			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
//			return
//		}
//		ctx.JSON(http.StatusCreated, user)
//	}
//
// Создаем нового пользователя без авторизации

func (uc *userController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return
	}
	// Хэширование пароля
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Ошибка": "Ошибка при хэшировании пароля"})
		return
	}
	user.Password = hashedPassword

	if err := uc.usersService.CreateUser(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Метод для удаления пользователя
func (uc *userController) DeleteUser(ctx *gin.Context) {
	jwtPayload, jwtPayloadErr := middlewares.GetJWTPayload(ctx)
	if jwtPayloadErr != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Ошибка": jwtPayloadErr.Error()})
		return
	}

	fmt.Println("JWT Payload:", jwtPayload)

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
	ctx.Status(http.StatusNoContent)
}

// Метод для редактирования данных пользователя
func (uc *userController) EditUser(ctx *gin.Context) {
	jwtPayload, jwtPayloadErr := middlewares.GetJWTPayload(ctx)
	if jwtPayloadErr != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Ошибка": jwtPayloadErr.Error()})
		return
	}

	fmt.Println("JWT Payload:", jwtPayload)

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

	// Хеширование пароля, если он был изменён
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Ошибка": "Не удалось захешировать пароль"})
			return
		}
		user.Password = string(hashedPassword)
	}

	if err := uc.usersService.EditUser(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
