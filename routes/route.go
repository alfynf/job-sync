package routes

import (
	"jobsync-be/controllers"
	"jobsync-be/controllers/api"
	"jobsync-be/controllers/company_controllers"
	"jobsync-be/controllers/employee_controllers"
	"jobsync-be/controllers/job_vacancy_controllers"
	"jobsync-be/controllers/login_controllers"
	"jobsync-be/controllers/user_controllers"
	"jobsync-be/lib/utils"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controllers.PingHandler)
	apiRoute := r.Group("/api")
	v1Route := apiRoute.Group("/v1")

	v1Route.GET("job-vacancies", api.GetList)
	v1Route.GET("job-vacancies/:job_vacancy_uuid", api.GetJobDetail)

	userRoute := v1Route.Group("/users")
	userRoute.POST("/", user_controllers.Create)
	userRoute.POST("/login", login_controllers.LoginUser)

	employeeRoute := v1Route.Group("/employees")
	employeeRoute.POST("/", employee_controllers.Create)
	employeeRoute.POST("/login", login_controllers.LoginEmployee)

	adminRoute := r.Group("/api/admin")
	v1AdminRoute := adminRoute.Group("/v1")

	authorizedEmployeeRoute := v1AdminRoute.Group("/employees")
	authorizedEmployeeRoute.Use(utils.CheckJWT())
	authorizedEmployeeRoute.GET("/my-profile", employee_controllers.GetDetail)
	authorizedEmployeeRoute.PUT("/my-profile", employee_controllers.Update)

	authorizedCompanyRoute := v1AdminRoute.Group("/companies/:company_uuid")
	authorizedCompanyRoute.Use(utils.CheckJWT())
	authorizedCompanyRoute.GET("/", company_controllers.Detail)
	authorizedCompanyRoute.PUT("/", company_controllers.Update)

	authorizedJobVacancyRoute := authorizedCompanyRoute.Group("/job-vacancies")
	authorizedJobVacancyRoute.POST("/", job_vacancy_controllers.Create)
	authorizedJobVacancyRoute.GET("/", job_vacancy_controllers.GetList)
	authorizedJobVacancyRoute.GET("/:job_vacancy_uuid", job_vacancy_controllers.GetDetail)
	authorizedJobVacancyRoute.PUT("/:job_vacancy_uuid", job_vacancy_controllers.Update)
	authorizedJobVacancyRoute.DELETE("/:job_vacancy_uuid", job_vacancy_controllers.Delete)
	authorizedJobVacancyRoute.POST("/:job_vacancy_uuid/apply", job_vacancy_controllers.ApplyJob)
	authorizedJobVacancyRoute.PUT("/:job_vacancy_uuid/applicant/:applicant_uuid", job_vacancy_controllers.UpdateApplicant)

	authorizedUserRoute := v1AdminRoute.Group("/users")
	authorizedUserRoute.Use(utils.CheckJWT())
	authorizedUserRoute.GET("/my-profile", user_controllers.GetDetail)
	authorizedUserRoute.PUT("/my-profile", user_controllers.Update)
	authorizedUserRoute.DELETE("/my-profile", user_controllers.Delete)

	return r
}
