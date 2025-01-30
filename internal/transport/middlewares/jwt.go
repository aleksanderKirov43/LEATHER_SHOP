package middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"leather-shop/config"
	"leather-shop/internal/models"
	"leather-shop/pkg/consts"
	"leather-shop/pkg/jwt"
	"net/http"
	"strings"
	"time"
)

type Helper struct {
	Secret     string
	AccessTTL  int
	RefreshTTL int
}

func GenerateTokenMiddleware(tokenHandler jwt.TokenHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, _ := ctx.Get("userId")
		username, _ := ctx.Get("username")

		accessToken, err := tokenHandler.GenerateToken(userId.(int), username.(string), time.Duration(15)*time.Minute)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Ошибка": "Не удалось сгенерировать access токен"})
			ctx.Abort()
			return
		}

		refreshToken, err := tokenHandler.GenerateToken(userId.(int), username.(string), time.Duration(60)*time.Minute)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Ошибка": "Не удалось сгенерировать refresh токен"})
			ctx.Abort()
			return
		}

		ctx.Set("accessToken", accessToken)
		ctx.Set("refreshToken", refreshToken)
		ctx.Next()
	}
}

// Проверяем наличие и валидность JWT-токена в заголовках запроса
func JwtMiddleware(tokenHandler jwt.TokenHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var authErrorCode int
		var authHeader string

		_, jwtPayloadErr := GetJWTPayload(ctx)
		if jwtPayloadErr != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Ошибка": jwtPayloadErr.Error()})
			return
		}

		authHeaderRefresh := ctx.GetHeader("Authorization-Refresh")
		authHeaderAccess := ctx.GetHeader("Authorization")

		if authHeaderRefresh != "" {
			authErrorCode = 400
			authHeader = authHeaderRefresh
		} else {
			authErrorCode = 401
			authHeader = authHeaderAccess
		}

		if authHeader == "" {
			ctx.JSON(authErrorCode, gin.H{
				"message": "токен не найден",
			})
			ctx.Abort()
			return
		}

		headersPair := strings.Split(authHeader, " ")
		if len(headersPair) != 2 {
			ctx.JSON(authErrorCode, gin.H{
				"message": "неверный формат заголовка",
			})
			ctx.Abort()
			return
		}

		if headersPair[0] != "Bearer" {
			ctx.JSON(authErrorCode, gin.H{
				"message": "неверный формат заголовка",
			})
			ctx.Abort()
			return
		}

		appConfig := config.GetConfig()
		helper := jwt.NewHelper(appConfig.Jwt.Secret, appConfig.Jwt.AccessTTL, appConfig.Jwt.RefreshTTL)

		payload, err := helper.ParseToken(headersPair[1])
		if err != nil {
			fmt.Println(err)
			ctx.JSON(authErrorCode, gin.H{
				"message": "невалидный токен",
			})
			ctx.Abort()
			return
		}

		session := helper.ParseMapClaims(payload)
		ctx.Set(consts.ContextUserSession, session)
		ctx.Next()
	}
}

// Функция извлечения закодированной JWT информации, для аутиндефикации и авторизайции пользователей
func GetJWTPayload(c *gin.Context) (*models.JWTPayload, error) {
	ctx := c.Value(consts.ContextUserSession)
	if ctx == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "ошибка верификации сессии"})
		return nil, errors.New("ошибка верификации сессии")
	}
	jwtPayload := ctx.(*models.JWTPayload)
	return jwtPayload, nil
}
