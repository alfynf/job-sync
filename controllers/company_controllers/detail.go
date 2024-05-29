package company_controllers

import (
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils"
	"jobsync-be/lib/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type detailResult struct {
	UUID        string           `json:"uuid"`
	Name        string           `json:"name"`
	EstablishAt string           `json:"establish_at"`
	UpdatedAt   string           `json:"updated_at"`
	Location    string           `json:"location"`
	Logo        string           `json:"logo"`
	Description string           `json:"description"`
	Address     string           `json:"address"`
	Email       string           `json:"email"`
	Phone       string           `json:"phone"`
	Website     string           `json:"website"`
	Employees   []resultEmployee `json:"employees"`
}

type resultEmployee struct {
	UUID           string `json:"uuid"`
	Name           string `json:"name"`
	Position       string `json:"position"`
	ProfilePicture string `json:"profile_picture"`
}

func Detail(c *gin.Context) {
	uuid := c.Param("company_uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Data not found", nil))
		return
	}

	company, err := q.GetCompanyByUUID(uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to fetch data", err))
		return
	}

	res := detailResult{
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

	var listEmployees []resultEmployee
	for _, employee := range company.Employees {
		v := resultEmployee{
			UUID:     employee.UUID.String(),
			Name:     employee.FirstName + employee.LastName,
			Position: employee.Position.Name,
		}

		if employee.ProfilePicture != nil {
			mConfig := utils.Init()
			employeeLogoUrl := mConfig.GetPresignedUrl(*employee.ProfilePicture)
			v.ProfilePicture = employeeLogoUrl
		}

		listEmployees = append(listEmployees, v)
	}

	res.Employees = listEmployees

	if company.Logo != nil {
		mConfig := utils.Init()
		companyLogoUrl := mConfig.GetPresignedUrl(*company.Logo)
		res.Logo = companyLogoUrl
	}

	// bikin get url dari minio

	c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Create Data",
		"data":    res,
	})
}
