package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	jwt.MapClaims
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Group    string `json:"Group"`
}

func GenerateJwtToken(userUuid string) (string, error) {
	expiredIn := os.Getenv("TOKEN_EXPIRED_TIME")
	expiredAt, err := strconv.Atoi(expiredIn)
	if err != nil {
		return "", err
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")

	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_uuid":  userUuid,
			"expired_at": time.Now().Local().Add(time.Second * time.Duration(expiredAt)),
		})

	token, err := t.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, nil

}
