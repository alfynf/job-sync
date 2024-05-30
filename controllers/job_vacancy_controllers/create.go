package job_vacancy_controllers

import (
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils/responses"
	"jobsync-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateJobVacancy struct {
	Title       string `form:"title" json:"title" validate:"required,max=255"`
	Location    string `form:"location" json:"location" validate:"required,max=50"`
	Requirement string `form:"requirement" json:"requirement" validate:"required"`
	JobType     int    `form:"job_type" json:"job_type" validate:"required,max=2"`
	WorkModel   int    `form:"work_model" json:"work_model" validate:"required,max=2"`
	EndDate     string `form:"end_date" json:"end_date" validate:"required"`
	Status      int    `form:"status" json:"status" validate:"max=1"`
}

func Create(c *gin.Context) {
	body := CreateJobVacancy{}
	c.ShouldBind(&body)

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Validation errors", err))
		return
	}

	jobVacancy := models.JobVacancy{
		Title:       body.Title,
		Location:    body.Location,
		Requirement: body.Requirement,
		JobType:     body.JobType,
		WorkModel:   body.WorkModel,
		EndDate:     body.EndDate,
		Status:      body.Status,
	}

	// save current user company
	uuid := c.MustGet("user-uuid")
	if uuid == nil {
		c.JSON(http.StatusNotFound, responses.ResponseBadRequest("User UUID not found", nil))
		return
	}

	employee, err := q.GetEmployeeByUUID(uuid.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ResponseBadRequest("Failed to fetch data", err))
		return
	}

	jobVacancy.EmployeeUUID = employee.UUID
	jobVacancy.CompanyUUID = employee.CompanyUUID

	err = q.CreateJobVacancy(jobVacancy)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to save data", err))
		return
	}

	c.JSON(http.StatusCreated, responses.ResponseCreated("Success create data"))
}
