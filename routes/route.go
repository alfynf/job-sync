package routes

import (
	"jobsync-be/controllers"
	"jobsync-be/controllers/employee_controllers"
	"jobsync-be/controllers/login_controllers"
	"jobsync-be/controllers/user_controllers"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controllers.PingHandler)
	apiRoute := r.Group("/api")
	v1Route := apiRoute.Group("/v1")

	userRoute := v1Route.Group("/users")
	userRoute.POST("/", user_controllers.Create)
	userRoute.POST("/login", login_controllers.LoginUser)

	employeeRoute := v1Route.Group("/employees")
	employeeRoute.POST("/", employee_controllers.Create)

	return r
}
