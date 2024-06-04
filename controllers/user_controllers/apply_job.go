package user_controllers

import (
	"io"
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils"
	"jobsync-be/lib/utils/responses"
	"jobsync-be/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ApplyJobUser struct {
	CV    *string `form:"cv" json:"cv" validate:"omitempty"`
	Notes *string `form:"notes" json:"notes" validate:"omitempty"`
}

func ApplyJob(c *gin.Context) {
	body := ApplyJobUser{}
	c.ShouldBind(&body)

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Validation errors", err))
		return
	}

	userUUID := c.MustGet("user-uuid")
	if userUUID == nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("User UUID not found", nil))
		return
	}

	user, err := q.GetUserByUUID(userUUID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ResponseBadRequest("Failed to fetch data", err))
		return
	}

	jobVacancyUUID := c.Param("job_vacancy_uuid")
	if jobVacancyUUID == "" {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Data not found", nil))
		return
	}
	jobVacancy, err := q.GetJobVacancyByUUID(jobVacancyUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ResponseBadRequest("Failed to fetch data", err))
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed get multipart form", err))
		return
	}

	var cvName string
	applicant := models.Applicant{
		Status:         0,
		UserUUID:       user.UUID,
		JobVacancyUUID: jobVacancy.UUID,
	}

	if body.Notes != nil && *body.Notes != "" {
		applicant.Notes = *body.Notes
	}

	applicant.Status = 0

	files := form.File["cv"]
	if files != nil {
		for _, file := range files {
			mConfig := utils.Init()
			src, err := file.Open()
			if err != nil {
				c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to open file", err))
				return
			}

			cvName = "/tmp/" + file.Filename
			defer src.Close()

			temp, err := os.Create(cvName)
			if err != nil {
				c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to create temporary file", err))
			}
			defer os.Remove(temp.Name())

			if _, err := io.Copy(temp, src); err != nil {
				c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to copy file", err))
				return
			}
			if err := temp.Close(); err != nil {
				c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to close file", err))
				return
			}

			err = mConfig.Store(temp.Name())
			if err != nil {
				c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to store file", err))
				return
			}
		}
		applicant.CV = cvName
	}

	err = q.CreateApplicant(applicant)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to save data", err))
		return
	}

	c.JSON(http.StatusCreated, responses.ResponseCreated("Success apply job"))
}
