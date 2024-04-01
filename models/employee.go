package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
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

type CreateEmployee struct {
	FirstName string `form:"first_name" json:"first_name" validate:"required,max=50"`
	LastName  string `form:"last_name" json:"last_name" validate:"required,max=50"`
}
