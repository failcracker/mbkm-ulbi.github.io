package services

import (
	"mbkm-ulbi-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MonthlyLogService struct {
	db *gorm.DB
}

func NewMonthlyLogService(db *gorm.DB) *MonthlyLogService {
	return &MonthlyLogService{db: db}
}

func (s *MonthlyLogService) GetMonthlyLogs(page, limit int) ([]models.MonthlyLog, int64, error) {
	var logs []models.MonthlyLog
	var total int64

	query := s.db.Model(&models.MonthlyLog{}).
		Preload("ApplyJob").
		Preload("ApplyJob.User").
		Preload("ApplyJob.User.ProfilePicture").
		Preload("ApplyJob.Job")

	// Count total
	query.Count(&total)

	// Apply pagination
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}

func (s *MonthlyLogService) GetMonthlyLogByID(id uuid.UUID) (*models.MonthlyLog, error) {
	var log models.MonthlyLog
	if err := s.db.Preload("ApplyJob").
		Preload("ApplyJob.User").
		Preload("ApplyJob.User.ProfilePicture").
		Preload("ApplyJob.Job").
		First(&log, id).Error; err != nil {
		return nil, err
	}
	return &log, nil
}

func (s *MonthlyLogService) CreateMonthlyLog(logData map[string]interface{}) (*models.MonthlyLog, error) {
	log := models.MonthlyLog{
		ApplyJobID: uuid.MustParse(logData["apply_job_id"].(string)),
		Content:    logData["content"].(string),
		Hasil:      logData["hasil"].(string),
		Status:     "Draft",
	}

	// Parse dates if provided
	if startDate, ok := logData["start_date"].(string); ok {
		// Parse date string to time.Time
		_ = startDate
	}
	if endDate, ok := logData["end_date"].(string); ok {
		// Parse date string to time.Time
		_ = endDate
	}

	if err := s.db.Create(&log).Error; err != nil {
		return nil, err
	}

	return &log, nil
}

func (s *MonthlyLogService) UpdateMonthlyLog(logData map[string]interface{}) error {
	id := uuid.MustParse(logData["id"].(string))
	
	updates := map[string]interface{}{
		"content": logData["content"].(string),
		"hasil":   logData["hasil"].(string),
	}

	return s.db.Model(&models.MonthlyLog{}).Where("id = ?", id).Updates(updates).Error
}

func (s *MonthlyLogService) ApproveMonthlyLog(id uuid.UUID) error {
	return s.db.Model(&models.MonthlyLog{}).Where("id = ?", id).Update("status", "Disetujui").Error
}

func (s *MonthlyLogService) RevisionMonthlyLog(id uuid.UUID, feedback string) error {
	updates := map[string]interface{}{
		"status":   "Revisi",
		"feedback": feedback,
	}
	return s.db.Model(&models.MonthlyLog{}).Where("id = ?", id).Updates(updates).Error
}