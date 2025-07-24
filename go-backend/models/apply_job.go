package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApplyJob struct {
	ID                         uuid.UUID    `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID                     uuid.UUID    `json:"user_id" gorm:"type:uuid;not null"`
	JobID                      uuid.UUID    `json:"job_id" gorm:"type:uuid;not null"`
	Status                     string       `json:"status" gorm:"default:'Melamar'"`
	Email                      string       `json:"email"`
	PhoneNumber                string       `json:"phone_number"`
	Address                    string       `json:"address"`
	DHSID                      *uuid.UUID   `json:"dhs_id"`
	DHS                        *File        `json:"dhs" gorm:"foreignKey:DHSID"`
	CVID                       *uuid.UUID   `json:"cv_id"`
	CV                         *File        `json:"cv" gorm:"foreignKey:CVID"`
	SuratLamaranID             *uuid.UUID   `json:"surat_lamaran_id"`
	SuratLamaran               *File        `json:"surat_lamaran" gorm:"foreignKey:SuratLamaranID"`
	SuratRekomendasiProdiID    *uuid.UUID   `json:"surat_rekomendasi_prodi_id"`
	SuratRekomendasiProdi      *File        `json:"surat_rekomendasi_prodi" gorm:"foreignKey:SuratRekomendasiProdiID"`
	ResponsibleLecturerID      *uuid.UUID   `json:"responsible_lecturer_id"`
	ResponsibleLecturer        *User        `json:"responsible_lecturer" gorm:"foreignKey:ResponsibleLecturerID"`
	ExaminerLecturerID         *uuid.UUID   `json:"examiner_lecturer_id"`
	ExaminerLecturer           *User        `json:"examiner_lecturer" gorm:"foreignKey:ExaminerLecturerID"`
	CreatedAt                  time.Time    `json:"created_at"`
	UpdatedAt                  time.Time    `json:"updated_at"`
	DeletedAt                  gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	User        User          `json:"users" gorm:"foreignKey:UserID"`
	Job         Job           `json:"jobs" gorm:"foreignKey:JobID"`
	Reports     []Report      `json:"reports" gorm:"foreignKey:ApplyJobID"`
	Evaluations []Evaluation  `json:"evaluations" gorm:"foreignKey:ApplyJobID"`
	MonthlyLogs []MonthlyLog  `json:"monthly_logs" gorm:"foreignKey:ApplyJobID"`
	KonversiNilai []KonversiNilai `json:"konversi_nilai" gorm:"foreignKey:ApplyJobID"`
}