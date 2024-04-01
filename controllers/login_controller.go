package controllers

import (
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginBody struct {
	Email    string `form:"email" json:"email" validate:"required,email"`
	Password string `form:"password" json:"password" validate:"required,min=8"`
}

func LoginUserController(c *gin.Context) {
	body := LoginBody{}
	c.ShouldBind(&body)

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseBadRequest("Validation errors", err))
		return
	}

	data, err := q.GetUserByEmail(body.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseBadRequest("Failed get user", err))
		return
	}

	if data.Password != body.Password {
		c.JSON(http.StatusBadRequest, utils.ResponseBadRequest("Password or Email not match", err))
		return
	}

	token, err := utils.GenerateJwtToken(data.UUID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseBadRequest("Error generate token", err))
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Login Success",
		"data": map[string]interface{}{
			"token":      token,
			"expires_in": os.Getenv("TOKEN_EXPIRED_TIME"),
		},
	})
}
