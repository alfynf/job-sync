package job_vacancy_controllers

import (
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils"
	"jobsync-be/lib/utils/responses"
	"jobsync-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type listResult struct {
	UUID        string `json:"uuid"`
	Title       string `json:"title"`
	Location    string `json:"location"`
	CreatedAt   string `json:"created_at"`
	CompanyUUID string `json:"company_uuid"`
	CompanyName string `json:"company_name"`
	CompanyLogo string `json:"company_logo"`
}

func GetList(c *gin.Context) {
	uuidParam := c.Param("company_uuid")
	if uuidParam == "" {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Data not found", nil))
		return
	}

	companyUUID, _ := uuid.Parse(uuidParam)

	jobVacancy, err := q.GetJobVacancy(&models.JobVacancy{CompanyUUID: companyUUID})
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ResponseBadRequest("Failed to fetch data", err))
		return
	}

	var res []listResult

	for _, v := range jobVacancy {
		data := listResult{
			UUID:        v.UUID.String(),
			Title:       v.Title,
			Location:    v.Location,
			CompanyUUID: v.CompanyUUID.String(),
			CompanyName: v.Company.Name,
			CreatedAt:   v.CreatedAt.Format("02-01-2006 15:03:03"),
		}

		if v.Company.Logo != nil {
			mConfig := utils.Init()
			companyLogoUrl := mConfig.GetPresignedUrl(*v.Company.Logo)
			data.CompanyLogo = companyLogoUrl
		}
		res = append(res, data)
	}

	c.JSON(http.StatusOK, responses.ResponseSuccessWithData("", res))
}
