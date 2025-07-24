package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID                     uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CompanyName            string    `json:"company_name" gorm:"not null"`
	BusinessField          string    `json:"business_field"`
	CompanySize            string    `json:"company_size"`
	CompanyAddress         string    `json:"company_address"`
	CompanyWebsite         string    `json:"company_website"`
	CompanyPhoneNumber     string    `json:"company_phone_number"`
	CompanyDescription     string    `json:"company_description"`
	CompanyLogoID          *uuid.UUID `json:"company_logo_id"`
	CompanyLogo            *File     `json:"company_logo" gorm:"foreignKey:CompanyLogoID"`
	Status                 string    `json:"status" gorm:"default:'Pending'"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
	DeletedAt              gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	Jobs []Job `json:"jobs" gorm:"foreignKey:CompanyID"`
}