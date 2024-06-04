package api

import (
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils"
	"jobsync-be/lib/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type detailCompanyResult struct {
	UUID        string            `json:"uuid"`
	Name        string            `json:"name"`
	EstablishAt string            `json:"establish_at"`
	UpdatedAt   string            `json:"updated_at"`
	Location    string            `json:"location"`
	Logo        string            `json:"logo"`
	Description string            `json:"description"`
	Address     string            `json:"address"`
	Email       string            `json:"email"`
	Phone       string            `json:"phone"`
	Website     string            `json:"website"`
	VacantJobs  []vacantJobResult `json:"vacant_jobs"`
}

type vacantJobResult struct {
	UUID      string `json:"uuid"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
}

func GetCompanyDetail(c *gin.Context) {
	uuid := c.Param("company_uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Data not found", nil))
		return
	}

	company, err := q.GetCompanyByUUID(uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ResponseBadRequest("Failed to fetch data", err))
		return
	}

	res := detailCompanyResult{
		UUID:        company.UUID.String(),
		Name:        company.Name,
		EstablishAt: company.EstablishAt,
		UpdatedAt:   company.UpdatedAt.Format("02-01-2006 15:03:03"),
		Location:    company.Location,
		Description: company.Description,
		Address:     company.Address,
		Email:       company.Email,
		Phone:       *company.Phone,
	}

	for _, job := range company.JobVacancies {
		v := vacantJobResult{
			UUID:      job.UUID.String(),
			Title:     job.Title,
			CreatedAt: job.CreatedAt.Format("02-01-2006 15:03:03"),
		}

		res.VacantJobs = append(res.VacantJobs, v)
	}

	if company.Logo != nil {
		mConfig := utils.Init()
		companyLogoUrl := mConfig.GetPresignedUrl(*company.Logo)
		res.Logo = companyLogoUrl
	}

	c.JSON(http.StatusOK, responses.ResponseSuccessWithData("", res))

}
