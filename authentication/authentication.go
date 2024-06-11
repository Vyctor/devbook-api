package authentication

import (
	"devbook-api/src/config"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

func CreateToken(userId uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.AppSecret)
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	if tokenString == "" {
		return errors.New("token inválido")
	}
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("token inválido")
}

func extractToken(r *http.Request) string {
	rawToken := r.Header.Get("Authorization")

	if len(strings.Split(rawToken, " ")) == 2 {
		return strings.Split(rawToken, " ")[1]
	}

	return ""
}

func ExtractUserId(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := uint64(claims["userId"].(float64))
		return userId, nil
	}
	return 0, errors.New("token inválido")
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.AppSecret, nil
}
