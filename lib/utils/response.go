package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseBadRequest(msg string, err error) map[string]interface{} {
	if msg == "" {
		msg = "Bad Request"
	}

	res := gin.H{
		"code":    http.StatusBadRequest,
		"message": "Bad Request",
		"error":   fmt.Sprintf("%v", err),
	}
	return res
}
