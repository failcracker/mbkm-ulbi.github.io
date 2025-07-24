package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                   uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Username             string     `json:"username" gorm:"unique;not null"`
	Email                string     `json:"email" gorm:"unique;not null"`
	Password             string     `json:"-" gorm:"not null"`
	Name                 string     `json:"name"`
	NIM                  string     `json:"nim" gorm:"unique"`
	ProgramStudy         string     `json:"program_study"`
	Faculty              string     `json:"faculty"`
	Semester             int        `json:"semester"`
	IPK                  float64    `json:"ipk"`
	PhoneNumber          string     `json:"phone_number"`
	Address              string     `json:"address"`
	SocialMedia          string     `json:"social_media"`
	EmergencyContact     string     `json:"emergency_contact"`
	ProfileDescription   string     `json:"profile_description"`
	Status               string     `json:"status" gorm:"default:'Aktif'"`
	ProfilePictureID     *uuid.UUID `json:"profile_picture_id"`
	ProfilePicture       *File      `json:"profile_picture" gorm:"foreignKey:ProfilePictureID"`
	Position             string     `json:"position"`
	CompanyName          string     `json:"company_name"`
	BusinessField        string     `json:"business_field"`
	CompanySize          string     `json:"company_size"`
	CompanyAddress       string     `json:"company_address"`
	CompanyWebsite       string     `json:"company_website"`
	CompanyPhoneNumber   string     `json:"company_phone_number"`
	CompanyDescription   string     `json:"company_description"`
	CompanyLogoID        *uuid.UUID `json:"company_logo_id"`
	CompanyLogo          *File      `json:"company_logo" gorm:"foreignKey:CompanyLogoID"`
	EmailVerifiedAt      *time.Time `json:"email_verified_at"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	Roles     []Role     `json:"roles" gorm:"many2many:user_roles;"`
	ApplyJobs []ApplyJob `json:"apply_jobs" gorm:"foreignKey:UserID"`
}

type Role struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title       string    `json:"title" gorm:"unique;not null"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserRole struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	RoleID uuid.UUID `json:"role_id" gorm:"type:uuid;not null"`
	User   User      `json:"user" gorm:"foreignKey:UserID"`
	Role   Role      `json:"role" gorm:"foreignKey:RoleID"`
}