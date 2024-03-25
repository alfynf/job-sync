package controllers

import (
	"fmt"
	"io"
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils"
	"jobsync-be/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateUserController(c *gin.Context) {
	body := models.CreateUser{}
	c.ShouldBind(&body)

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, utils.ResponseBadRequest(err))
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

	fmt.Println("SAMPAI SINI 2")

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseBadRequest(err))
		return
	}

	var profilePictureName string

	files := form.File["profile_picture"]
	if files != nil {
		for _, file := range files {
			mConfig := utils.Init()
			src, err := file.Open()
			if err != nil {
				c.JSON(http.StatusBadRequest, utils.ResponseBadRequest(err))
				return
			}

			profilePictureName = "/tmp/" + file.Filename
			defer src.Close()

			temp, err := os.Create(profilePictureName)
			if err != nil {
				log.Fatal(err)
			}
			defer os.Remove(temp.Name())

			if _, err := io.Copy(temp, src); err != nil {
				c.JSON(http.StatusBadRequest, utils.ResponseBadRequest(err))
				return
			}
			if err := temp.Close(); err != nil {
				c.JSON(http.StatusBadRequest, utils.ResponseBadRequest(err))
				return
			}

			err = mConfig.Store(temp.Name())
			if err != nil {
				c.JSON(http.StatusBadRequest, utils.ResponseBadRequest(err))
				return
			}
		}
		user.ProfilePicture = &profilePictureName
	}

	fmt.Println(user)

	err = q.Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseBadRequest(err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success create new users",
	})
}
