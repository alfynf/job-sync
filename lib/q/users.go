package q

import (
	"fmt"
	"jobsync-be/configs"
	"jobsync-be/models"

	"github.com/google/uuid"
)

func CreateUser(user models.User) error {
	user.UUID = uuid.New()
	res := configs.DB.Create(&user)
	if res.Error != nil {
		return fmt.Errorf("failed to create on database: %v", res.Error)
	}
	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var data models.User
	res := configs.DB.Where("email = ?", email).First(&data)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get from database: %v", res.Error)
	}
	return &data, nil
}

func GetUserByUUID(uuid string) (*models.User, error) {
	var data models.User
	res := configs.DB.Where("uuid = ?", uuid).First(&data)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get from database: %v", res.Error)
	}
	return &data, nil
}

func UpdateUser(user *models.User) error {
	res := configs.DB.Save(user)
	if res.Error != nil {
		return fmt.Errorf("failed to create on database: %v", res.Error)
	}
	return nil
}

func DeleteUser(user *models.User) error {
	res := configs.DB.Delete(user)
	if res.Error != nil {
		return fmt.Errorf("failed to create on database: %v", res.Error)
	}
	return nil
}
