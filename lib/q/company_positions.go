package q

import (
	"fmt"
	"jobsync-be/configs"
	"jobsync-be/models"

	"github.com/google/uuid"
)

func CreateCompanyPosition(companyPosition models.CompanyPosition) error {
	res := configs.DB.Create(&companyPosition)
	if res.Error != nil {
		return fmt.Errorf("failed to create on database: %v", res.Error)
	}
	return nil
}

func GetCompanyPositionByUUID(uuid string) (*models.CompanyPosition, error) {
	var data models.CompanyPosition
	res := configs.DB.Where("uuid = ?", uuid).First(&data)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get from database: %v", res.Error)
	}
	return &data, nil
}

func GetCompanyPositionByCompanyAndName(companyUuid uuid.UUID, name string) (*models.CompanyPosition, error) {
	var data models.CompanyPosition
	res := configs.DB.Where("company_uuid = ? AND name = ?", companyUuid, name).First(&data)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get from database: %v", res.Error)
	}
	return &data, nil
}
