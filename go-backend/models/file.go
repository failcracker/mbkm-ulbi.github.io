package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"not null"`
	URL         string    `json:"url" gorm:"not null"`
	PublicID    string    `json:"public_id"`
	Size        int64     `json:"size"`
	MimeType    string    `json:"mime_type"`
	UploadedBy  uuid.UUID `json:"uploaded_by" gorm:"type:uuid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}