package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"leather-shop/internal/models"
)

// Структура для хранения секретного ключа для подписи и верификации JWT-токенов
type Helper struct {
	Secret     string
	AccessTTL  int
	RefreshTTL int
}

// Создание нового экземпляра Helper
func NewHelper(secret string, accessTTL, refreshTTL int) *Helper {
	return &Helper{
		Secret:     secret,
		AccessTTL:  accessTTL,
		RefreshTTL: refreshTTL,
	}
}

// Генерация токена
func (h *Helper) GenerateToken(userId int, username string, ttl time.Duration) (string, error) {
	// Создаем claims (утверждения) для токена, включающие идентификатор пользователя, имя пользователя и срок действия токена
	claims := &jwt.MapClaims{
		"user_id":  userId,
		"username": username,
		"exp":      time.Now().Add(ttl).Unix(),
	}
	// Создаем новый JWT-токен с заданными claims и методом подписи HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Подписываем токен с использованием секретного ключа и возвращаем токен в виде строки
	tokenString, err := token.SignedString([]byte(h.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Функция генерации access / refresh токенов
func (h *Helper) GenerateAccessAndRefreshTokens(userId int, username string) (string, string, error) {
	accessToken, err := h.GenerateToken(userId, username, time.Duration(h.AccessTTL)*time.Minute)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := h.GenerateToken(userId, username, time.Duration(h.RefreshTTL)*time.Minute)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// Парсинг токена
func (h *Helper) ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, err
	} else {
		if err = claims.Valid(); err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		return claims, nil
	}
}

func (h *Helper) ParseMapClaims(mapClaims jwt.MapClaims) *models.JWTPayload {
	userIdFloat, ok := mapClaims["user_id"].(float64)
	if !ok {
		return nil
	}

	return &models.JWTPayload{
		Id:    int(userIdFloat),
		Login: mapClaims["username"].(string),
	}
}

type TokenHandler interface {
	GenerateToken(userId int, username string, ttl time.Duration) (string, error)
	ParseToken(tokenString string) (jwt.MapClaims, error)
	ParseMapClaims(mapClaims jwt.MapClaims) *models.JWTPayload
}
