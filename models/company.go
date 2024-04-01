package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	UUID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // Standard field for the primary key
	EstablishAt string    `gorm:"size:4;not null"`
	Location    string    `gorm:"size:50;not null"`
	Description string    `gorm:"not null"`
	Address     string    `gorm:"size:255;not null"`
	Email       string    `gorm:"unique;size:255;not null"`
	Phone       *string   `gorm:"size:20"`
	Logo        *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Employees   []Employee        `gorm:"foreignKey:CompanyUUID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
	Positions   []CompanyPosition `gorm:"foreignKey:CompanyUUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CreateCompany struct {
	FirstName string `form:"first_name" json:"first_name" validate:"required,max=50"`
	LastName  string `form:"last_name" json:"last_name" validate:"required,max=50"`
}
