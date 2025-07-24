package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Report struct {
	ID                  uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ApplyJobID          uuid.UUID  `json:"apply_job_id" gorm:"type:uuid;not null"`
	FileLaporanID       *uuid.UUID `json:"file_laporan_id"`
	FileLaporan         *File      `json:"file_laporan" gorm:"foreignKey:FileLaporanID"`
	Status              string     `json:"status" gorm:"default:'Draft'"`
	CompanyCheckedID    *uuid.UUID `json:"company_checked_id"`
	CompanyChecked      *User      `json:"company_checked" gorm:"foreignKey:CompanyCheckedID"`
	LecturerCheckedID   *uuid.UUID `json:"lecturer_checked_id"`
	LecturerChecked     *User      `json:"lecturer_checked" gorm:"foreignKey:LecturerCheckedID"`
	ExaminerCheckedID   *uuid.UUID `json:"examiner_checked_id"`
	ExaminerChecked     *User      `json:"examiner_checked" gorm:"foreignKey:ExaminerCheckedID"`
	ProdiCheckedID      *uuid.UUID `json:"prodi_checked_id"`
	ProdiChecked        *User      `json:"prodi_checked" gorm:"foreignKey:ProdiCheckedID"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	ApplyJob ApplyJob `json:"apply_job" gorm:"foreignKey:ApplyJobID"`
}