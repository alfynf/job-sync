package q

import (
	"fmt"
	"jobsync-be/configs"
	"jobsync-be/models"
)

func Create(user models.User) error {
	res := configs.DB.Create(&user)
	if res.Error != nil {
		return fmt.Errorf("failed to create on database: %v", res.Error)
	}
	return nil
}

func GetByEmail(email string) (*models.User, error) {
	var data models.User
	res := configs.DB.Where("email = ?", email).First(&data)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get from database: %v", res.Error)
	}
	return &data, nil
}
