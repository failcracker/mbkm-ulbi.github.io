package services

import (
	"mbkm-ulbi-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApplyJobService struct {
	db *gorm.DB
}

func NewApplyJobService(db *gorm.DB) *ApplyJobService {
	return &ApplyJobService{db: db}
}

func (s *ApplyJobService) GetApplyJobs(page, limit int) ([]models.ApplyJob, int64, error) {
	var applyJobs []models.ApplyJob
	var total int64

	query := s.db.Model(&models.ApplyJob{}).
		Preload("User").
		Preload("User.ProfilePicture").
		Preload("Job").
		Preload("Job.JobVacancyImage").
		Preload("ResponsibleLecturer").
		Preload("ExaminerLecturer")

	// Count total
	query.Count(&total)

	// Apply pagination
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&applyJobs).Error; err != nil {
		return nil, 0, err
	}

	return applyJobs, total, nil
}

func (s *ApplyJobService) GetApplyJobByID(id uuid.UUID) (*models.ApplyJob, error) {
	var applyJob models.ApplyJob
	if err := s.db.Preload("User").
		Preload("User.ProfilePicture").
		Preload("Job").
		Preload("Job.JobVacancyImage").
		Preload("ResponsibleLecturer").
		Preload("ExaminerLecturer").
		Preload("DHS").
		Preload("CV").
		Preload("SuratLamaran").
		Preload("SuratRekomendasiProdi").
		First(&applyJob, id).Error; err != nil {
		return nil, err
	}
	return &applyJob, nil
}

func (s *ApplyJobService) GetApplyJobsByUser(userID uuid.UUID) ([]models.ApplyJob, error) {
	var applyJobs []models.ApplyJob
	if err := s.db.Where("user_id = ?", userID).
		Preload("User").
		Preload("User.ProfilePicture").
		Preload("Job").
		Preload("Job.JobVacancyImage").
		Find(&applyJobs).Error; err != nil {
		return nil, err
	}
	return applyJobs, nil
}

func (s *ApplyJobService) GetLastApplyJobByUser(userID uuid.UUID) (*models.ApplyJob, error) {
	var applyJob models.ApplyJob
	if err := s.db.Where("user_id = ? AND status = ?", userID, "Aktif").
		Preload("User").
		Preload("User.ProfilePicture").
		Preload("Job").
		Preload("Job.JobVacancyImage").
		Preload("Reports").
		Preload("Reports.FileLaporan").
		Preload("Evaluations").
		Preload("Evaluations.BobotNilai").
		First(&applyJob).Error; err != nil {
		return nil, err
	}
	return &applyJob, nil
}

func (s *ApplyJobService) CreateApplyJob(applyJobData map[string]interface{}) (*models.ApplyJob, error) {
	applyJob := models.ApplyJob{
		UserID:      uuid.MustParse(applyJobData["user_id"].(string)),
		JobID:       uuid.MustParse(applyJobData["job_id"].(string)),
		Email:       applyJobData["email"].(string),
		PhoneNumber: applyJobData["phone_number"].(string),
		Address:     applyJobData["address"].(string),
		Status:      "Melamar",
	}

	if err := s.db.Create(&applyJob).Error; err != nil {
		return nil, err
	}

	return &applyJob, nil
}

func (s *ApplyJobService) ApproveApplyJob(id uuid.UUID) error {
	return s.db.Model(&models.ApplyJob{}).Where("id = ?", id).Update("status", "Disetujui").Error
}

func (s *ApplyJobService) RejectApplyJob(id uuid.UUID) error {
	return s.db.Model(&models.ApplyJob{}).Where("id = ?", id).Update("status", "Ditolak").Error
}

func (s *ApplyJobService) ActivateApplyJob(id uuid.UUID) error {
	return s.db.Model(&models.ApplyJob{}).Where("id = ?", id).Update("status", "Aktif").Error
}

func (s *ApplyJobService) DoneApplyJob(id uuid.UUID) error {
	return s.db.Model(&models.ApplyJob{}).Where("id = ?", id).Update("status", "Selesai").Error
}

func (s *ApplyJobService) SetLecturer(id uuid.UUID, lecturerID, examinerID *uuid.UUID) error {
	updates := make(map[string]interface{})
	if lecturerID != nil {
		updates["responsible_lecturer_id"] = *lecturerID
	}
	if examinerID != nil {
		updates["examiner_lecturer_id"] = *examinerID
	}
	return s.db.Model(&models.ApplyJob{}).Where("id = ?", id).Updates(updates).Error
}