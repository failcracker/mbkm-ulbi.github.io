package services

import (
	"mbkm-ulbi-backend/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EvaluationService struct {
	db *gorm.DB
}

func NewEvaluationService(db *gorm.DB) *EvaluationService {
	return &EvaluationService{db: db}
}

func (s *EvaluationService) GetEvaluations(page, limit int) ([]models.Evaluation, int64, error) {
	var evaluations []models.Evaluation
	var total int64

	query := s.db.Model(&models.Evaluation{}).
		Preload("ApplyJob").
		Preload("ApplyJob.User").
		Preload("ApplyJob.User.ProfilePicture").
		Preload("ApplyJob.Job").
		Preload("ApplyJob.Job.JobVacancyImage").
		Preload("BobotNilai")

	// Count total
	query.Count(&total)

	// Apply pagination
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&evaluations).Error; err != nil {
		return nil, 0, err
	}

	return evaluations, total, nil
}

func (s *EvaluationService) GetEvaluationByID(id uuid.UUID) (*models.Evaluation, error) {
	var evaluation models.Evaluation
	if err := s.db.Preload("ApplyJob").
		Preload("ApplyJob.User").
		Preload("ApplyJob.User.ProfilePicture").
		Preload("ApplyJob.Job").
		Preload("ApplyJob.Job.JobVacancyImage").
		Preload("ApplyJob.SuratLamaran").
		Preload("BobotNilai").
		Where("apply_job_id = ?", id).
		First(&evaluation).Error; err != nil {
		return nil, err
	}
	return &evaluation, nil
}

func (s *EvaluationService) CreateEvaluation(evaluationData map[string]interface{}) (*models.Evaluation, error) {
	applyJobID := uuid.MustParse(evaluationData["apply_job_id"].(string))
	
	// Check if evaluation exists
	var evaluation models.Evaluation
	err := s.db.Where("apply_job_id = ?", applyJobID).First(&evaluation).Error
	
	if err == gorm.ErrRecordNotFound {
		// Create new evaluation
		evaluation = models.Evaluation{
			ApplyJobID: applyJobID,
		}
		s.db.Create(&evaluation)
	}

	// Update evaluation based on role
	updates := make(map[string]interface{})
	now := time.Now()

	if grade, ok := evaluationData["grade"].(string); ok {
		if gradeScore, ok := evaluationData["grade_score"].(string); ok {
			if gradeDesc, ok := evaluationData["grade_description"].(string); ok {
				// Determine which grade to update based on context
				if isExaminer, ok := evaluationData["is_examiner"].(string); ok && isExaminer == "1" {
					updates["examiner_grade"] = grade
					updates["examiner_grade_score"] = gradeScore
					updates["examiner_grade_description"] = gradeDesc
					updates["examiner_grade_date"] = now
				} else {
					// Default to company grade for now
					updates["company_grade"] = grade
					updates["company_grade_score"] = gradeScore
					updates["company_grade_description"] = gradeDesc
					updates["company_grade_date"] = now
				}
			}
		}
	}

	if err := s.db.Model(&evaluation).Updates(updates).Error; err != nil {
		return nil, err
	}

	// Calculate final grade if all grades are present
	s.calculateFinalGrade(&evaluation)

	return &evaluation, nil
}

func (s *EvaluationService) calculateFinalGrade(evaluation *models.Evaluation) {
	// Get bobot nilai
	var bobotNilai models.BobotNilai
	s.db.First(&bobotNilai)

	// Calculate weighted average if all grades are present
	if evaluation.CompanyGradeScore > 0 && evaluation.LecturerGradeScore > 0 && evaluation.ExaminerGradeScore > 0 {
		totalScore := (evaluation.CompanyGradeScore * bobotNilai.BobotNilaiPerusahaan / 100) +
			(evaluation.LecturerGradeScore * bobotNilai.BobotNilaiPembimbing / 100) +
			(evaluation.ExaminerGradeScore * bobotNilai.BobotNilaiPenguji / 100)

		var grade string
		switch {
		case totalScore >= 90:
			grade = "A"
		case totalScore >= 80:
			grade = "B"
		case totalScore >= 70:
			grade = "C"
		case totalScore >= 60:
			grade = "D"
		default:
			grade = "E"
		}

		updates := map[string]interface{}{
			"grade":       grade,
			"total_score": totalScore,
			"perhitungan_nilai_text": "Calculated based on weighted average",
			"bobot_nilai_id": bobotNilai.ID,
		}

		s.db.Model(evaluation).Updates(updates)
	}
}