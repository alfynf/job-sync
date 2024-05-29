package company_controllers

import (
	"fmt"
	"io"
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils"
	"jobsync-be/lib/utils/responses"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UpdateCompany struct {
	Name        *string `form:"name" json:"name" validate:"omitempty"`
	EstablishAt *string `form:"establish_at" json:"establish_at" validate:"omitempty,len=4"`
	Location    *string `form:"location" json:"location" validate:"omitempty,max=50"`
	Description *string `form:"description" json:"description"`
	Address     *string `form:"address" json:"address" validate:"omitempty,max=255"`
	Email       *string `form:"email" json:"email" validate:"omitempty,max=255"`
	Phone       *string `form:"phone" json:"phone" validate:"omitempty,max=20"`
	CompanyLogo *string `form:"logo" json:"logo" validate:"omitempty"`
}

func Update(c *gin.Context) {
	uuid := c.Param("company_uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Data not found", nil))
		return
	}

	body := UpdateCompany{}
	c.ShouldBind(&body)

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Validation errors", err))
		return
	}

	company, err := q.GetCompanyByUUID(uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ResponseBadRequest("Failed to fetch data", err))
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed get multipart form", err))
		return
	}

	logoFile := form.File["logo"]

	if logoFile != nil {
		var logoName string

		for _, file := range logoFile {
			mConfig := utils.Init()
			src, err := file.Open()
			if err != nil {
				c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to open file", err))
				return
			}

			logoName = "/tmp/" + file.Filename
			defer src.Close()

			temp, err := os.Create(logoName)
			if err != nil {
				c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to create temporary file", err))
			}
			defer os.Remove(temp.Name())

			if _, err := io.Copy(temp, src); err != nil {
				c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to copy file", err))
				return
			}
			if err := temp.Close(); err != nil {
				c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to close file", err))
				return
			}

			err = mConfig.Store(temp.Name())
			if err != nil {
				c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to store file", err))
				return
			}
		}
		company.Logo = &logoName
	}

	if body.EstablishAt != nil && *body.EstablishAt != "" {
		company.EstablishAt = *body.EstablishAt
	}
	if body.Location != nil && *body.Location != "" {
		company.Location = *body.Location
	}
	if body.Description != nil && *body.Description != "" {
		company.Description = *body.Description
	}
	if body.Address != nil && *body.Address != "" {
		company.Address = *body.Address
	}
	if body.Email != nil && *body.Email != "" {
		company.Email = *body.Email
	}
	if body.Phone != nil && *body.Phone != "" {
		company.Phone = *&body.Phone
	}
	if body.Email != nil && *body.Email != "" {
		company.Email = *body.Email
	}
	if body.Name != nil && *body.Name != "" {
		company.Name = *body.Name
	}
	fmt.Println(company)

	// save gambar, create employee
	err = q.UpdateCompany(company)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to save data", err))
		return
	}

	// bikin get url dari minio

	c.JSON(http.StatusOK, responses.ResponseSuccess("Update sucess"))
}
