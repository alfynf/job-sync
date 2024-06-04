package models

import (
	"time"

	"github.com/google/uuid"
)

type Applicant struct {
	UUID           uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // Standard field for the primary key
	UserUUID       uuid.UUID
	User           User `gorm:"foreignKey:UserUUID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
	JobVacancyUUID uuid.UUID
	JobVacancy     JobVacancy `gorm:"foreignKey:JobVacancyUUID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
	Notes          string
	CV             string
	Status         int `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}
