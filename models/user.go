package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID           uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // Standard field for the primary key
	FirstName      string    `gorm:"size:50;not null"`
	LastName       string    `gorm:"size:50;not null"`
	Username       string    `gorm:"unique;size:255;not null"`
	Email          string    `gorm:"unique;size:255;not null"`
	Password       string    `gorm:"not null"`
	Birthdate      string    `gorm:"size:10"`
	Gender         *int      `gorm:"type:enum('F', 'M')"`
	Phone          *string   `gorm:"size:20"`
	ProfilePicture *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}
