package user_controllers

import (
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
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
	err = q.DeleteUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to save data", err))
		return
	}

	c.JSON(http.StatusOK, responses.ResponseSuccess("Success delete data"))
}
