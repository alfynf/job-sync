package api

import (
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils"
	"jobsync-be/lib/utils/responses"
	"jobsync-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
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

type searchQuery struct {
	Title    string `form:"title" json:"title"`
	Location int    `form:"location" json:"location"`
}

func GetJobList(c *gin.Context) {

	search := &models.JobVacancy{}

	jobVacancy, err := q.GetJobVacancy(search)
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
