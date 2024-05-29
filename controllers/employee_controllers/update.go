package employee_controllers

import (
	"io"
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils"
	"jobsync-be/lib/utils/responses"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UpdateEmployee struct {
	FirstName       *string `form:"first_name" json:"first_name" validate:"omitempty,max=50"`
	LastName        *string `form:"last_name" json:"last_name" validate:"omitempty,max=50"`
	Username        *string `form:"username" json:"username" validate:"omitempty,max=255"`
	Email           *string `form:"email" json:"email" validate:"omitempty,email"`
	Password        *string `form:"password" json:"password" validate:"omitempty,min=8"`
	ConfirmPassword *string `form:"confirm_password" json:"confirm_password" validate:"omitempty,eqfield=Password"`
	ProfilePicture  *string `form:"profile_picture" json:"profile_picture" validate:"omitempty"`
}

func Update(c *gin.Context) {
	uuid := c.MustGet("user-uuid")
	if uuid == nil {
		c.JSON(http.StatusNotFound, responses.ResponseBadRequest("User UUID not found", nil))
		return
	}

	body := UpdateEmployee{}
	c.ShouldBind(&body)

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Validation errors", err))
		return
	}

	employee, err := q.GetEmployeeByUUID(uuid.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ResponseBadRequest("Failed to fetch data", err))
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed get multipart form", err))
		return
	}

	profilePictureFile := form.File["profile_picture"]

	if profilePictureFile != nil {
		var profilePictureName string

		for _, file := range profilePictureFile {
			mConfig := utils.Init()
			src, err := file.Open()
			if err != nil {
				c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to open file", err))
				return
			}

			profilePictureName = "/tmp/" + file.Filename
			defer src.Close()

			temp, err := os.Create(profilePictureName)
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
		employee.ProfilePicture = &profilePictureName
	}

	if body.FirstName != nil && *body.FirstName != "" {
		employee.FirstName = *body.FirstName
	}
	if body.LastName != nil && *body.LastName != "" {
		employee.LastName = *body.LastName
	}
	if body.Username != nil && *body.Username != "" {
		employee.Username = *body.Username
	}
	if body.Email != nil && *body.Email != "" {
		employee.Email = *body.Email
	}
	if body.Password != nil && *body.Password != "" {
		employee.Password = *body.Password
	}

	err = q.UpdateEmployee(employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to save data", err))
		return
	}

	c.JSON(http.StatusOK, responses.ResponseSuccess("Success Update Data"))
}
