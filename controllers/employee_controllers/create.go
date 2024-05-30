package employee_controllers

import (
	"io"
	"jobsync-be/lib/q"
	"jobsync-be/lib/utils"
	"jobsync-be/lib/utils/responses"
	"jobsync-be/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateEmployee struct {
	FirstName           string  `form:"first_name" json:"first_name" validate:"required,max=50"`
	LastName            string  `form:"last_name" json:"last_name" validate:"required,max=50"`
	Username            string  `form:"username" json:"username" validate:"required,max=255"`
	Email               string  `form:"email" json:"email" validate:"required,email"`
	Password            string  `form:"password" json:"password" validate:"required,min=8"`
	ConfirmPassword     string  `form:"confirm_password" json:"confirm_password" validate:"required,eqfield=Password"`
	IsCompanyRegistered bool    `form:"is_company_registered" json:"is_company_registered"`
	CompanyUUID         *string `form:"company_uuid" json:"company_uuid" validate:"excluded_if=IsCompanyRegistered false"`
	CompanyName         string  `form:"company_name" json:"company_name" validate:"excluded_if=IsCompanyRegistered true,omitempty,required"`
	CompanyEstablishAt  string  `form:"company_establish_at" json:"company_establish_at" validate:"excluded_if=IsCompanyRegistered true,omitempty,len=4"`
	CompanyLocation     string  `form:"company_location" json:"company_location" validate:"excluded_if=IsCompanyRegistered true,omitempty,max=50"`
	CompanyDescription  string  `form:"company_description" json:"company_description" validate:"excluded_if=IsCompanyRegistered true,omitempty,required"`
	CompanyAddress      string  `form:"company_address" json:"company_address" validate:"excluded_if=IsCompanyRegistered true,omitempty,max=255"`
	CompanyEmail        string  `form:"company_email" json:"company_email" validate:"excluded_if=IsCompanyRegistered true,omitempty,max=255"`
	CompanyPhone        *string `form:"company_phone" json:"company_phone" validate:"omitempty,max=20"`
	PositionUUID        *string `form:"position_uuid" json:"position_uuid" validate:"omitempty"`
	OtherPositionName   *string `form:"other_position_name" json:"other_position_name" validate:"omitempty"`
	CompanyLogo         *string `form:"company_logo" json:"company_logo" validate:"omitempty"`
	ProfilePicture      *string `form:"profile_picture" json:"profile_picture" validate:"omitempty"`
}

func Create(c *gin.Context) {
	body := CreateEmployee{}
	c.ShouldBind(&body)

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Validation errors", err))
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed get multipart form", err))
		return
	}

	logoFile := form.File["company_logo"]

	var company models.Company
	if body.IsCompanyRegistered == true {
		_, err := q.GetCompanyByUUID(*body.CompanyUUID)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed get company data", err))
			return
		}
	} else {
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
		company.EstablishAt = body.CompanyEstablishAt
		company.Location = body.CompanyLocation
		company.Description = body.CompanyDescription
		company.Address = body.CompanyAddress
		company.Email = body.CompanyEmail
		company.Phone = body.CompanyPhone
		company.Name = body.CompanyName
	}

	// save gambar, create employee
	err = q.CreateCompany(company)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to save data", err))
		return
	}

	user := models.Employee{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Username:  body.Username,
		Email:     body.Email,
		Password:  body.Password,
	}

	if body.CompanyUUID != nil && *body.CompanyUUID != "" {
		companyUUID, _ := uuid.Parse(*body.CompanyUUID)
		user.CompanyUUID = companyUUID
	}

	var profilePictureName string

	ppFiles := form.File["profile_picture"]
	if ppFiles != nil {

		for _, file := range ppFiles {
			mConfig := utils.Init()
			src, err := file.Open()
			if err != nil {
				c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to open file", err))
				return
			}

			profilePictureName = "/tmp/" + file.Filename
			defer src.Close()

			temp, err := os.Create(profilePictureName)
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
		user.ProfilePicture = &profilePictureName
	}

	var companyPosition *models.CompanyPosition

	// save company position
	if body.PositionUUID != nil && *body.PositionUUID != "" {
		// cek apakah positionnya ada
		companyPosition, err = q.GetCompanyPositionByUUID(*body.PositionUUID)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to get data", err))
			return
		}
		user.PositionUUID = companyPosition.UUID
	} else {
		// crete new company position
		companyPosition := models.CompanyPosition{}
		company, err := q.GetCompanyByName(body.CompanyName)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to get data", err))
			return
		}
		companyPosition.CompanyUUID = company.UUID
		companyPosition.Name = *body.OtherPositionName

		err = q.CreateCompanyPosition(companyPosition)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to save data", err))
			return
		}

		data, err := q.GetCompanyPositionByCompanyAndName(companyPosition.CompanyUUID, companyPosition.Name)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to save data", err))
			return
		}
		user.PositionUUID = data.UUID
		user.CompanyUUID = company.UUID
	}

	err = q.CreateEmployee(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ResponseBadRequest("Failed to save data", err))
		return
	}

	// update employeecompany position

	c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Create Data",
	})
}
