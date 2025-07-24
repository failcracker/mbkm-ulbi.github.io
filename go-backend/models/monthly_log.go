package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MonthlyLog struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ApplyJobID uuid.UUID `json:"apply_job_id" gorm:"type:uuid;not null"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	Content    string    `json:"content"`
	Hasil      string    `json:"hasil"`
	Status     string    `json:"status" gorm:"default:'Draft'"`
	Feedback   string    `json:"feedback"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	ApplyJob ApplyJob `json:"apply_job" gorm:"foreignKey:ApplyJobID"`
}