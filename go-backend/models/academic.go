package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProgramStudi struct {
	ID                uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	KodeProgramStudi  string    `json:"kode_program_studi" gorm:"unique;not null"`
	NamaProgramStudi  string    `json:"nama_program_studi" gorm:"not null"`
	Fakultas          string    `json:"fakultas"`
	Jenjang           string    `json:"jenjang"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	MataKuliah []MataKuliah `json:"mata_kuliah" gorm:"foreignKey:ProdiID"`
}

type MataKuliah struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ProdiID   uuid.UUID `json:"prodi_id" gorm:"type:uuid;not null"`
	Nama      string    `json:"nama" gorm:"not null"`
	Kode      string    `json:"kode" gorm:"unique;not null"`
	SKS       int       `json:"sks"`
	Semester  int       `json:"semester"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	ProgramStudi ProgramStudi `json:"program_studi" gorm:"foreignKey:ProdiID"`
}

type KonversiNilai struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ApplyJobID   uuid.UUID `json:"apply_job_id" gorm:"type:uuid;not null"`
	MataKuliahID uuid.UUID `json:"mata_kuliah_id" gorm:"type:uuid;not null"`
	Grade        string    `json:"grade"`
	Score        int       `json:"score"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	ApplyJob    ApplyJob    `json:"apply_job" gorm:"foreignKey:ApplyJobID"`
	MataKuliah  MataKuliah  `json:"mata_kuliah" gorm:"foreignKey:MataKuliahID"`
}

type BobotNilai struct {
	ID                    uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	BobotNilaiPerusahaan  int       `json:"bobot_nilai_perusahaan" gorm:"default:40"`
	BobotNilaiPembimbing  int       `json:"bobot_nilai_pembimbing" gorm:"default:30"`
	BobotNilaiPenguji     int       `json:"bobot_nilai_penguji" gorm:"default:30"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `json:"-" gorm:"index"`
}