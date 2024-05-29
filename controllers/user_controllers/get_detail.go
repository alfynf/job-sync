package user_controllers

import (
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils"
	"jobsync-be/lib/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type detailResult struct {
	UUID           string `json:"uuid"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Birthdate      string `json:"birthdate"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profile_picture"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Gender         int    `json:"gender"`
}

func GetDetail(c *gin.Context) {
	uuid := c.MustGet("user-uuid")
	if uuid == nil {
		c.JSON(http.StatusNotFound, responses.ResponseBadRequest("User UUID not found", nil))
		return
	}

	user, err := q.GetUserByUUID(uuid.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ResponseBadRequest("Failed to fetch data", err))
		return
	}

	res := detailResult{
		UUID:      user.UUID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Birthdate: user.Birthdate,
	}

	if user.Phone != nil {
		res.Phone = *user.Phone
	}

	if user.Gender != nil {
		res.Gender = *user.Gender
	}

	if user.ProfilePicture != nil {
		mConfig := utils.Init()
		userProfilePictureUrl := mConfig.GetPresignedUrl(*user.ProfilePicture)
		res.ProfilePicture = userProfilePictureUrl
	}

	c.JSON(http.StatusCreated, responses.ResponseSuccessWithData("", res))
}
