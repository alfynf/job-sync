package q

import (
	"fmt"
	"jobsync-be/configs"
	"jobsync-be/models"
)

func CreateEmployee(employee models.Employee) error {
	res := configs.DB.Create(&employee)
	if res.Error != nil {
		return fmt.Errorf("failed to create on database: %v", res.Error)
	}
	return nil
}

func GetEmployeeByEmail(email string) (*models.Employee, error) {
	var data models.Employee
	res := configs.DB.Where("email = ?", email).First(&data)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get from database: %v", res.Error)
	}
	return &data, nil
}

func GetEmployeeByUUID(uuid string) (*models.Employee, error) {
	var data models.Employee
	res := configs.DB.Preload("Position").Where("uuid = ?", uuid).First(&data)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get from database: %v", res.Error)
	}
	return &data, nil
}
