package models

import (
	"github.com/golang-jwt/jwt/v4"
	"leather-shop/config"
	"time"
)

// Создаём структуру для хранения токенов доступа
type JWT struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Создаём структуру для хранения данных пользователя
type JWTPayload struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	Type  string `json:"type"`
	jwt.RegisteredClaims
}

// Генерируем токены доступа и обновления на основе JWT
func (p *JWTPayload) Register(jwtConfig *config.JWTConfig) (*JWT, error) {
	accessToken, aErr := p.generate(time.Duration(jwtConfig.AccessTTL)*time.Minute, jwtConfig.Secret)
	if aErr != nil {
		return nil, aErr
	}

	refreshToken, rErr := p.generate(time.Duration(jwtConfig.RefreshTTL)*time.Minute, jwtConfig.Secret)
	if rErr != nil {
		return nil, rErr
	}

	return &JWT{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}, nil
}

// Генирируем JWT-токен
func (p *JWTPayload) generate(ttl time.Duration, secret string) (*string, error) {
	p.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, p)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &tokenStr, nil
}
