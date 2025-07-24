package services

import (
	"mbkm-ulbi-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReportService struct {
	db *gorm.DB
}

func NewReportService(db *gorm.DB) *ReportService {
	return &ReportService{db: db}
}

func (s *ReportService) GetReports(page, limit int) ([]models.Report, int64, error) {
	var reports []models.Report
	var total int64

	query := s.db.Model(&models.Report{}).
		Preload("ApplyJob").
		Preload("ApplyJob.User").
		Preload("ApplyJob.User.ProfilePicture").
		Preload("ApplyJob.Job").
		Preload("ApplyJob.Job.JobVacancyImage").
		Preload("FileLaporan")

	// Count total
	query.Count(&total)

	// Apply pagination
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&reports).Error; err != nil {
		return nil, 0, err
	}

	return reports, total, nil
}

func (s *ReportService) GetReportByID(id uuid.UUID) (*models.Report, error) {
	var report models.Report
	if err := s.db.Preload("ApplyJob").
		Preload("ApplyJob.User").
		Preload("ApplyJob.User.ProfilePicture").
		Preload("ApplyJob.Job").
		Preload("ApplyJob.Job.JobVacancyImage").
		Preload("FileLaporan").
		Preload("CompanyChecked").
		Preload("LecturerChecked").
		Preload("ExaminerChecked").
		Preload("ProdiChecked").
		First(&report, id).Error; err != nil {
		return nil, err
	}
	return &report, nil
}

func (s *ReportService) CreateReport(reportData map[string]interface{}) (*models.Report, error) {
	report := models.Report{
		ApplyJobID: uuid.MustParse(reportData["apply_job_id"].(string)),
		Status:     "Pending",
	}

	if err := s.db.Create(&report).Error; err != nil {
		return nil, err
	}

	return &report, nil
}

func (s *ReportService) CheckReport(id uuid.UUID, checkerID uuid.UUID, role string) error {
	updates := make(map[string]interface{})
	
	switch role {
	case "mitra":
		updates["company_checked_id"] = checkerID
	case "dosen":
		updates["lecturer_checked_id"] = checkerID
	case "examiner":
		updates["examiner_checked_id"] = checkerID
	case "prodi":
		updates["prodi_checked_id"] = checkerID
	}

	return s.db.Model(&models.Report{}).Where("apply_job_id = ?", id).Updates(updates).Error
}