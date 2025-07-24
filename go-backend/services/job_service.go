package services

import (
	"mbkm-ulbi-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobService struct {
	db *gorm.DB
}

func NewJobService(db *gorm.DB) *JobService {
	return &JobService{db: db}
}

func (s *JobService) GetJobs(page, limit int) ([]models.Job, int64, error) {
	var jobs []models.Job
	var total int64

	query := s.db.Model(&models.Job{}).
		Preload("JobVacancyImage").
		Preload("CompanyRef")

	// Count total
	query.Count(&total)

	// Apply pagination
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&jobs).Error; err != nil {
		return nil, 0, err
	}

	return jobs, total, nil
}

func (s *JobService) GetJobByID(id uuid.UUID) (*models.Job, error) {
	var job models.Job
	if err := s.db.Preload("JobVacancyImage").Preload("CompanyRef").First(&job, id).Error; err != nil {
		return nil, err
	}
	return &job, nil
}

func (s *JobService) CreateJob(jobData map[string]interface{}) (*models.Job, error) {
	job := models.Job{
		Title:       jobData["title"].(string),
		Company:     jobData["company"].(string),
		Location:    jobData["location"].(string),
		Description: jobData["description"].(string),
		Duration:    jobData["duration"].(string),
		JobType:     jobData["job_type"].(string),
		Benefits:    jobData["benefits"].(string),
		Status:      jobData["status"].(string),
		VacancyType: jobData["vacancy_type"].(string),
	}

	// Parse deadline if provided
	if deadline, ok := jobData["deadline"].(string); ok {
		// Parse deadline string to time.Time
		// You might want to add proper date parsing here
		_ = deadline
	}

	if err := s.db.Create(&job).Error; err != nil {
		return nil, err
	}

	return &job, nil
}

func (s *JobService) UpdateJob(id uuid.UUID, updates map[string]interface{}) error {
	return s.db.Model(&models.Job{}).Where("id = ?", id).Updates(updates).Error
}

func (s *JobService) DeleteJob(id uuid.UUID) error {
	return s.db.Delete(&models.Job{}, id).Error
}

func (s *JobService) ApproveJob(id uuid.UUID) error {
	return s.db.Model(&models.Job{}).Where("id = ?", id).Update("status", "Tersedia").Error
}

func (s *JobService) RejectJob(id uuid.UUID) error {
	return s.db.Model(&models.Job{}).Where("id = ?", id).Update("status", "Ditolak").Error
}

func (s *JobService) GetJobCandidates(jobID uuid.UUID) ([]models.User, error) {
	var users []models.User
	
	if err := s.db.Table("users").
		Select("users.*, apply_jobs.*").
		Joins("JOIN apply_jobs ON users.id = apply_jobs.user_id").
		Where("apply_jobs.job_id = ?", jobID).
		Preload("ProfilePicture").
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (s *JobService) GetCompanies() ([]models.Company, error) {
	var companies []models.Company
	if err := s.db.Preload("CompanyLogo").Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}