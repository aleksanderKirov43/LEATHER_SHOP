package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"leather-shop/config"
	"leather-shop/internal/models"
	"leather-shop/pkg/consts"
	"leather-shop/pkg/jwt"
)

type Helper struct {
	Secret     string
	AccessTTL  int
	RefreshTTL int
}

// Проверяем наличие и валидность JWT-токена в заголовках запроса
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var authErrorCode int
		var authHeader string

		authHeaderRefresh := c.GetHeader("Authorization-Refresh")
		authHeaderAccess := c.GetHeader("Authorization")

		if authHeaderRefresh != "" {
			authErrorCode = 400
			authHeader = authHeaderRefresh
		} else {
			authErrorCode = 401
			authHeader = authHeaderAccess
		}

		if authHeader == "" {
			c.JSON(authErrorCode, gin.H{
				"message": "токен не найден",
			})
			c.Abort()
			return
		}

		headersPair := strings.Split(authHeader, " ")
		if len(headersPair) != 2 {
			c.JSON(authErrorCode, gin.H{
				"message": "неверный формат заголовка",
			})
			c.Abort()
			return
		}

		if headersPair[0] != "Bearer" {
			c.JSON(authErrorCode, gin.H{
				"message": "неверный формат заголовка",
			})
			c.Abort()
			return
		}

		appConfig := config.GetConfig()
		helper := jwt.NewHelper(appConfig.Jwt.Secret, appConfig.Jwt.AccessTTL, appConfig.Jwt.RefreshTTL)

		payload, err := helper.ParseToken(headersPair[1])

		if err != nil {
			fmt.Println(err)
			c.JSON(authErrorCode, gin.H{
				"message": "невалидный токен",
			})
			c.Abort()
			return
		}

		session := helper.ParseMapClaims(payload)

		c.Set(consts.ContextUserSession, session)

		c.Next()
	}
}

func GetJWTPayload(c *gin.Context) (*models.JWTPayload, error) {
	ctx := c.Value(consts.ContextUserSession)
	if ctx == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "ошибка верификации сессии"})
		return nil, errors.New("ошибка верификации сессии")
	}
	jwtPayload := ctx.(*models.JWTPayload)
	return jwtPayload, nil
}
