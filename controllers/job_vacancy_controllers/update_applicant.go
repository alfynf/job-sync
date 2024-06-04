package job_vacancy_controllers

import (
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UpdateApplicantStatus struct {
	Status int `form:"status" json:"status" validate:"required"`
}

func UpdateApplicant(c *gin.Context) {
	body := UpdateApplicantStatus{}
	c.ShouldBind(&body)

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Validation errors", err))
		return
	}

	applicantUUID := c.Param("applicant_uuid")
	if applicantUUID == "" {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("User UUID not found", nil))
		return
	}

	applicant, err := q.GetApplicantByUUID(applicantUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to fetch data", err))
		return
	}

	applicant.Status = body.Status

	err = q.UpdateApplicant(applicant)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to save data", err))
		return
	}

	c.JSON(http.StatusOK, responses.ResponseCreated("Success update applicant status"))
}
