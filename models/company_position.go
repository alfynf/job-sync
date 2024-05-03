package models

import (
	"time"

	"github.com/google/uuid"
)

type CompanyPosition struct {
	UUID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // Standard field for the primary key
	Name        string    `gorm:"size:100;not null"`
	CompanyUUID uuid.UUID
	Company     Company `gorm:"foreignKey:CompanyUUID;constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type CreateCompanyPosition struct {
	FirstName string `form:"first_name" json:"first_name" validate:"required,max=50"`
	LastName  string `form:"last_name" json:"last_name" validate:"required,max=50"`
}
