package models

import (
	"time"

	"github.com/google/uuid"
)

type JobVacancy struct {
	UUID         uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // Standard field for the primary key
	Title        string    `gorm:"size:255;not null"`
	Location     string    `gorm:"size:50;not null"`
	Requirement  string    `gorm:"not null"`
	JobType      int       `gorm:"not null"`
	WorkModel    int       `gorm:"not null"`
	EndDate      string    `gorm:"not null"`
	Status       int       `gorm:"not null"`
	CompanyUUID  uuid.UUID
	Company      Company `gorm:"foreignKey:CompanyUUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADEs;"`
	EmployeeUUID uuid.UUID
	Applicants   []Applicant `gorm:"foreignKey:JobVacancyUUID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL;;"`
	CreatedBy    Employee    `gorm:"foreignKey:EmployeeUUID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
