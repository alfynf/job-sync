package q

import (
	"fmt"
	"jobsync-be/configs"
	"jobsync-be/models"
)

func CreateJobVacancy(jobVacancy models.JobVacancy) error {
	res := configs.DB.Create(&jobVacancy)
	if res.Error != nil {
		return fmt.Errorf("failed to create on database: %v", res.Error)
	}
	return nil
}

func GetJobVacancyByUUID(uuid string) (*models.JobVacancy, error) {
	var data models.JobVacancy
	res := configs.DB.Preload("Employees.Position").Preload("Employees").Where("uuid = ?", uuid).First(&data)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get from database: %v", res.Error)
	}
	return &data, nil
}

func GetJobVacancyByName(name string) (*models.JobVacancy, error) {
	var data models.JobVacancy
	res := configs.DB.Where("name = ?", name).First(&data)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get from database: %v", res.Error)
	}
	return &data, nil
}

func UpdateJobVacancy(jobVacancy *models.JobVacancy) error {
	res := configs.DB.Save(jobVacancy)
	if res.Error != nil {
		return fmt.Errorf("failed to create on database: %v", res.Error)
	}
	return nil
}
