package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseBadRequest(err error) map[string]interface{} {
	res := gin.H{
		"code":    http.StatusBadRequest,
		"message": "Bad Request",
		"error":   fmt.Sprintf("%v", err),
	}
	return res
}
