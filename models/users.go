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

type CreateUser struct {
	FirstName       string  `form:"first_name" json:"first_name" validate:"required,max=50"`
	LastName        string  `form:"last_name" json:"last_name" validate:"required,max=50"`
	Username        string  `form:"username" json:"username" validate:"required,max=255"`
	Email           string  `form:"email" json:"email" validate:"required,email"`
	Password        string  `form:"password" json:"password" validate:"required,min=8"`
	ConfirmPassword string  `form:"confirm_password" json:"confirm_password" validate:"required,eqfield=Password"`
	Birthdate       string  `form:"birthdate" json:"birthdate" validate:"required"`
	Gender          *int    `form:"gender" json:"gender" validate:"required"`
	Phone           *string `form:"phone" json:"phone"`
	ProfilePicture  *string `form:"profile_picture" json:"profile_picture"`
}
