package responses

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
		"message": msg,
		"error":   fmt.Sprintf("%v", err),
	}
	return res
}

func ResponseSuccess(msg interface{}) map[string]interface{} {
	if msg == nil {
		msg = "Success Retrieve Data"
	}

	res := gin.H{
		"code":    http.StatusOK,
		"message": msg,
	}
	return res
}

func ResponseSuccessWithData(msg interface{}, data interface{}) map[string]interface{} {
	if msg == "" {
		msg = "Success Retrieve Data"
	}

	res := gin.H{
		"code":    http.StatusOK,
		"message": msg,
		"data":    data,
	}
	return res
}

func ResponseCreated(msg string) map[string]interface{} {
	if msg == "" {
		msg = "Success Create Data"
	}

	res := gin.H{
		"code":    http.StatusCreated,
		"message": msg,
	}
	return res
}
