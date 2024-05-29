package utils

import (
	"fmt"
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils/responses"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	jwt.MapClaims
	UserUUID  string    `json:"user_uuid"`
	ExpiredAt time.Time `json:"expired_at"`
	UserType  int       `json:"user_type"`
}

func GenerateJwtToken(userUuid string, userType int) (string, error) {
	expiredIn := os.Getenv("TOKEN_EXPIRED_TIME")
	expiredAt, err := strconv.Atoi(expiredIn)
	if err != nil {
		return "", err
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")

	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		MyClaims{
			UserUUID:  userUuid,
			ExpiredAt: time.Now().Local().Add(time.Second * time.Duration(expiredAt)),
			UserType:  userType,
		})
	token, err := t.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, nil

}

func CheckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		secretKey := os.Getenv("JWT_SECRET_KEY")
		headerToken := c.Request.Header["Authorization"]
		if headerToken == nil {
			c.JSON(http.StatusUnauthorized, responses.ResponseUnauthorized("", fmt.Errorf("Token not found")))
		}
		splitToken := strings.Split(headerToken[0], "Bearer ")
		tokenString := splitToken[1]

		token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, responses.ResponseUnauthorized("", fmt.Errorf("Failed parse token")))
		} else if claims, ok := token.Claims.(*MyClaims); ok {
			if time.Now().Compare(claims.ExpiredAt) == 1 {
				c.JSON(http.StatusUnauthorized, responses.ResponseUnauthorized("", fmt.Errorf("Token Expired")))
			}
			if claims.UserType == 1 {
				user, err := q.GetUserByUUID(claims.UserUUID)
				if err != nil {
					c.JSON(http.StatusUnauthorized, responses.ResponseUnauthorized("", fmt.Errorf("User not found")))
				}
				c.Set("user-uuid", user.UUID.String())
			} else if claims.UserType == 2 {
				user, err := q.GetEmployeeByUUID(claims.UserUUID)
				if err != nil {
					c.JSON(http.StatusUnauthorized, responses.ResponseUnauthorized("", fmt.Errorf("User not found")))
				}
				c.Set("user-uuid", user.UUID.String())
			}
		} else {
			c.JSON(http.StatusUnauthorized, responses.ResponseUnauthorized("", fmt.Errorf("Unrecognizable claim, unable to proceed")))
		}
		c.Next()
	}
}
