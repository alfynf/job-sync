package q

import (
	"fmt"
	"jobsync-be/configs"
	"jobsync-be/models"
)

func CreateCompanyPosition(company_position models.CompanyPosition) error {
	res := configs.DB.Create(&company_position)
	if res.Error != nil {
		return fmt.Errorf("failed to create on database: %v", res.Error)
	}
	return nil
}

func GetCompanyByUUID(uuid string) (*models.Company, error) {
	var data models.Company
	res := configs.DB.Where("uuid = ?", uuid).First(&data)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get from database: %v", res.Error)
	}
	return &data, nil
}
