package util

import (
	"fmt"
	"log"
	"strings"
	"time"

	"GoToDo/config"

	"github.com/golang-jwt/jwt/v4"
)

var secret_key = []byte(config.SecretKey)

const TokenExpireDuration = time.Hour * 2

type UserClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(id uint, username string) (string, error) {
	claims := &UserClaims{
		id,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "GoToDo",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Println("secret_key:", secret_key)
	tokenString, err := token.SignedString(secret_key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*UserClaims, error) {
	if strings.HasPrefix(strings.ToLower(tokenString), "bearer ") {
		tokenString = tokenString[7:]
	}

	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return secret_key, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("parse token error:%v", err)
	}
	return userClaim, nil
}
