package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Evaluation struct {
	ID                        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ApplyJobID                uuid.UUID  `json:"apply_job_id" gorm:"type:uuid;not null"`
	CompanyGrade              string     `json:"company_grade"`
	CompanyGradeScore         int        `json:"company_grade_score"`
	CompanyGradeDescription   string     `json:"company_grade_description"`
	CompanyGradeDate          *time.Time `json:"company_grade_date"`
	LecturerGrade             string     `json:"lecturer_grade"`
	LecturerGradeScore        int        `json:"lecturer_grade_score"`
	LecturerGradeDescription  string     `json:"lecturer_grade_description"`
	LecturerGradeDate         *time.Time `json:"lecturer_grade_date"`
	ExaminerGrade             string     `json:"examiner_grade"`
	ExaminerGradeScore        int        `json:"examiner_grade_score"`
	ExaminerGradeDescription  string     `json:"examiner_grade_description"`
	ExaminerGradeDate         *time.Time `json:"examiner_grade_date"`
	ProdiGrade                string     `json:"prodi_grade"`
	ProdiGradeScore           int        `json:"prodi_grade_score"`
	ProdiGradeDescription     string     `json:"prodi_grade_description"`
	ProdiGradeDate            *time.Time `json:"prodi_grade_date"`
	Grade                     string     `json:"grade"`
	TotalScore                int        `json:"total_score"`
	PerhitunganNilaiText      string     `json:"perhitungan_nilai_text"`
	BobotNilaiID              *uuid.UUID `json:"bobot_nilai_id"`
	BobotNilai                *BobotNilai `json:"bobot_nilai" gorm:"foreignKey:BobotNilaiID"`
	CreatedAt                 time.Time  `json:"created_at"`
	UpdatedAt                 time.Time  `json:"updated_at"`
	DeletedAt                 gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	ApplyJob ApplyJob `json:"apply_job" gorm:"foreignKey:ApplyJobID"`
}