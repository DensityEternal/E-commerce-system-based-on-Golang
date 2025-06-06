package utils

import (
	"E_commerce_System/models"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func ParseToken(tokenString string) (*models.Claims, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET env variable not set")
	}
	// 解析并验证 Token
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	// 解析失败
	if err != nil {
		return nil, err
	}
	// 断言为自定义的 Claims 类型
	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {

		return claims, nil
	}
	return nil, errors.New("invalid token")
}
