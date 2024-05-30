package job_vacancy_controllers

import (
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UpdateJobVacancy struct {
	Title       *string `form:"title" json:"title" validate:"omitempty,max=255"`
	Location    *string `form:"location" json:"location" validate:"omitempty,max=50"`
	Requirement *string `form:"requirement" json:"requirement" validate:"omitempty"`
	JobType     *int    `form:"job_type" json:"job_type" validate:"omitempty,max=2"`
	WorkModel   *int    `form:"work_model" json:"work_model" validate:"omitempty,max=2"`
	EndDate     *string `form:"end_date" json:"end_date" validate:"omitempty"`
	Status      *int    `form:"status" json:"status" validate:"max=1"`
}

func Update(c *gin.Context) {
	body := UpdateJobVacancy{}
	c.ShouldBind(&body)

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Validation errors", err))
		return
	}

	uuid := c.Param("job_vacancy_uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Data not found", nil))
		return
	}

	jobVacancy, err := q.GetJobVacancyByUUID(uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ResponseBadRequest("Failed to fetch data", err))
		return
	}

	if body.Title != nil && *body.Title != "" {
		jobVacancy.Title = *body.Title
	}
	if body.Location != nil && *body.Location != "" {
		jobVacancy.Location = *body.Location
	}
	if body.Requirement != nil && *body.Requirement != "" {
		jobVacancy.Requirement = *body.Requirement
	}
	if body.JobType != nil {
		jobVacancy.JobType = *body.JobType
	}
	if body.WorkModel != nil {
		jobVacancy.WorkModel = *body.WorkModel
	}
	if body.EndDate != nil {
		jobVacancy.EndDate = *body.EndDate
	}
	if body.Status != nil {
		jobVacancy.Status = *body.Status
	}

	err = q.UpdateJobVacancy(jobVacancy)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to save data", err))
		return
	}

	c.JSON(http.StatusOK, responses.ResponseCreated("Success update data"))
}
