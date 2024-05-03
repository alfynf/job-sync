package user_controllers

import (
	"io"
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils"
	"jobsync-be/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateUser struct {
	FirstName       string  `form:"first_name" json:"first_name" validate:"required,max=50"`
	LastName        string  `form:"last_name" json:"last_name" validate:"required,max=50"`
	Username        string  `form:"username" json:"username" validate:"required,max=255"`
	Email           string  `form:"email" json:"email" validate:"required,email"`
	Password        string  `form:"password" json:"password" validate:"required,min=8"`
	ConfirmPassword string  `form:"confirm_password" json:"confirm_password" validate:"required,eqfield=Password"`
	Birthdate       string  `form:"birthdate" json:"birthdate" validate:"required"`
	Gender          *int    `form:"gender" json:"gender" validate:"required"`
	Phone           *string `form:"phone" json:"phone"`
	ProfilePicture  *string `form:"profile_picture" json:"profile_picture"`
}

func Create(c *gin.Context) {
	body := CreateUser{}
	c.ShouldBind(&body)

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseBadRequest("Validation errors", err))
		return
	}

	user := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Username:  body.Username,
		Email:     body.Email,
		Password:  body.Password,
		Birthdate: body.Birthdate,
		Gender:    body.Gender,
		Phone:     body.Phone,
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseBadRequest("Failed get multipart form", err))
		return
	}

	var profilePictureName string

	files := form.File["profile_picture"]
	if files != nil {
		for _, file := range files {
			mConfig := utils.Init()
			src, err := file.Open()
			if err != nil {
				c.JSON(http.StatusBadRequest, utils.ResponseBadRequest("Failed to open file", err))
				return
			}

			profilePictureName = "/tmp/" + file.Filename
			defer src.Close()

			temp, err := os.Create(profilePictureName)
			if err != nil {
				c.JSON(http.StatusBadRequest, utils.ResponseBadRequest("Failed to create temporary file", err))
			}
			defer os.Remove(temp.Name())

			if _, err := io.Copy(temp, src); err != nil {
				c.JSON(http.StatusBadRequest, utils.ResponseBadRequest("Failed to copy file", err))
				return
			}
			if err := temp.Close(); err != nil {
				c.JSON(http.StatusBadRequest, utils.ResponseBadRequest("Failed to close file", err))
				return
			}

			err = mConfig.Store(temp.Name())
			if err != nil {
				c.JSON(http.StatusBadRequest, utils.ResponseBadRequest("Failed to store file", err))
				return
			}
		}
		user.ProfilePicture = &profilePictureName
	}

	err = q.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseBadRequest("Failed to save data", err))
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Create Data",
	})
}
