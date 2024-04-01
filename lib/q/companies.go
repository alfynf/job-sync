package q

import (
	"fmt"
	"jobsync-be/configs"
	"jobsync-be/models"
)

func CreateCompany(company models.Company) error {
	res := configs.DB.Create(&company)
	if res.Error != nil {
		return fmt.Errorf("failed to create on database: %v", res.Error)
	}
	return nil
}
