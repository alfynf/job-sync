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

func GetCompanyByUUID(uuid string) (*models.Company, error) {
	var data models.Company
	res := configs.DB.Preload("Employees.Position").Preload("Employees").Preload("JobVacancies").Where("uuid = ?", uuid).First(&data)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get from database: %v", res.Error)
	}
	return &data, nil
}

func GetCompanyByName(name string) (*models.Company, error) {
	var data models.Company
	res := configs.DB.Where("name = ?", name).First(&data)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get from database: %v", res.Error)
	}
	return &data, nil
}

func UpdateCompany(company *models.Company) error {
	res := configs.DB.Save(company)
	if res.Error != nil {
		return fmt.Errorf("failed to create on database: %v", res.Error)
	}
	return nil
}
