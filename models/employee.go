package models

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	UUID           uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // Standard field for the primary key
	FirstName      string    `gorm:"size:50;not null"`
	LastName       string    `gorm:"size:50;not null"`
	Username       string    `gorm:"unique;size:255;not null"`
	Email          string    `gorm:"unique;size:255;not null"`
	Password       string    `gorm:"not null"`
	CompanyUUID    string
	Company        Company `gorm:"foreignKey:CompanyUUID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
	PositionUUID   uuid.UUID
	Position       CompanyPosition `gorm:"foreignKey:PositionUUID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
	ProfilePicture *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}
