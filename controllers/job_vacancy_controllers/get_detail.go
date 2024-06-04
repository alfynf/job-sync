package job_vacancy_controllers

import (
	"fmt"
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils"
	"jobsync-be/lib/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type detailResult struct {
	UUID        string            `json:"uuid"`
	Title       string            `json:"title"`
	Location    string            `json:"location"`
	CreatedAt   string            `json:"created_at"`
	Requirement string            `json:"requirement"`
	JobType     int               `json:"job_type"`
	WorkModel   int               `json:"work_model"`
	EndDate     string            `json:"end_date"`
	CompanyUUID string            `json:"company_uuid"`
	CompanyName string            `json:"company_name"`
	CompanyLogo string            `json:"company_logo"`
	Applicants  []applicantResult `json:"applicants"`
}

type applicantResult struct {
	UUID           string  `json:"uuid"`
	Name           string  `json:"name"`
	CV             *string `json:"cv"`
	ProfilePicture *string `json:"profile_picture"`
	Email          string  `json:"emial"`
	Phone          *string `json:"phone"`
}

func GetDetail(c *gin.Context) {
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

	res := detailResult{
		UUID:        jobVacancy.UUID.String(),
		Title:       jobVacancy.Title,
		Location:    jobVacancy.Location,
		Requirement: jobVacancy.Requirement,
		JobType:     jobVacancy.JobType,
		WorkModel:   jobVacancy.WorkModel,
		EndDate:     jobVacancy.EndDate,
		CompanyUUID: jobVacancy.CompanyUUID.String(),
		CompanyName: jobVacancy.Company.Name,
		CreatedAt:   jobVacancy.CreatedAt.Format("02-01-2006 15:03:03"),
	}

	if jobVacancy.Applicants != nil {
		fmt.Println("masuk sini")
		for _, v := range jobVacancy.Applicants {
			applicant := applicantResult{
				UUID:  v.UUID.String(),
				Name:  v.User.FirstName + v.User.LastName,
				Email: v.User.Email,
				Phone: v.User.Phone,
			}

			mConfig := utils.Init()
			cvUrl := mConfig.GetPresignedUrl(v.CV)
			applicant.CV = &cvUrl

			if v.User.ProfilePicture != nil {
				mConfig := utils.Init()
				userProfilePictureUrl := mConfig.GetPresignedUrl(*v.User.ProfilePicture)
				applicant.ProfilePicture = &userProfilePictureUrl
			}

			res.Applicants = append(res.Applicants, applicant)
		}
	}

	if jobVacancy.Company.Logo != nil {
		mConfig := utils.Init()
		companyLogoUrl := mConfig.GetPresignedUrl(*jobVacancy.Company.Logo)
		res.CompanyLogo = companyLogoUrl
	}

	c.JSON(http.StatusOK, responses.ResponseSuccessWithData("", res))
}
