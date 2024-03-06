package routes

import (
	"jobsync-be/controllers"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", controllers.PingHandler)

	return r
}
