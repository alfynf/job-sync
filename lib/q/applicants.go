package q

import (
	"fmt"
	"jobsync-be/configs"
	"jobsync-be/models"

	"github.com/google/uuid"
)

func CreateApplicant(applicant models.Applicant) error {
	applicant.UUID = uuid.New()
	res := configs.DB.Create(&applicant)
	if res.Error != nil {
		return fmt.Errorf("failed to create on database: %v", res.Error)
	}
	return nil
}

func GetApplicantByUUID(uuid string) (*models.Applicant, error) {
	var data models.Applicant
	res := configs.DB.Where("uuid = ?", uuid).First(&data)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get from database: %v", res.Error)
	}
	return &data, nil
}

func UpdateApplicant(applicant *models.Applicant) error {
	res := configs.DB.Save(applicant)
	if res.Error != nil {
		return fmt.Errorf("failed to create on database: %v", res.Error)
	}
	return nil
}
