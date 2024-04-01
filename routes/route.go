package routes

import (
	"jobsync-be/controllers"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controllers.PingHandler)
	apiRoute := r.Group("/api")
	v1Route := apiRoute.Group("/v1")
	userRoute := v1Route.Group("/users")

	userRoute.POST("/", controllers.CreateUserController)
	userRoute.POST("/login", controllers.LoginUserController)

	return r
}
