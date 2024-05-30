package job_vacancy_controllers

import (
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
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

	err = q.DeleteJobVacancy(jobVacancy)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to save data", err))
		return
	}

	c.JSON(http.StatusOK, responses.ResponseSuccess("Success delete data"))
}
