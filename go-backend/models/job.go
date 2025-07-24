package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Job struct {
	ID                uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title             string     `json:"title" gorm:"not null"`
	Company           string     `json:"company" gorm:"not null"`
	Location          string     `json:"location"`
	Description       string     `json:"description"`
	Duration          string     `json:"duration"`
	JobType           string     `json:"job_type"`
	Benefits          string     `json:"benefits"`
	Status            string     `json:"status" gorm:"default:'Perlu Ditinjau'"`
	VacancyType       string     `json:"vacancy_type"`
	Deadline          time.Time  `json:"deadline"`
	JobVacancyImageID *uuid.UUID `json:"job_vacancy_image_id"`
	JobVacancyImage   *File      `json:"job_vacancy_image" gorm:"foreignKey:JobVacancyImageID"`
	CompanyID         *uuid.UUID `json:"company_id"`
	CompanyRef        *Company   `json:"company_ref" gorm:"foreignKey:CompanyID"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	ApplyJobs []ApplyJob `json:"apply_jobs" gorm:"foreignKey:JobID"`
}