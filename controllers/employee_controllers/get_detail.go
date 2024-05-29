package employee_controllers

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
	Username       string `json:"username"`
	ProfilePicture string `json:"profile_picture"`
	Email          string `json:"email"`
	Position       string `json:"position"`
}

func GetDetail(c *gin.Context) {
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

	res := detailResult{
		UUID:      employee.UUID.String(),
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
		Username:  employee.Username,
		Email:     employee.Email,
		Position:  employee.Position.Name,
	}

	if employee.ProfilePicture != nil {
		mConfig := utils.Init()
		employeeProfilePictureUrl := mConfig.GetPresignedUrl(*employee.ProfilePicture)
		res.ProfilePicture = employeeProfilePictureUrl
	}

	c.JSON(http.StatusCreated, responses.ResponseSuccessWithData("", res))
}
