package services

import (
	"mbkm-ulbi-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AcademicService struct {
	db *gorm.DB
}

func NewAcademicService(db *gorm.DB) *AcademicService {
	return &AcademicService{db: db}
}

func (s *AcademicService) GetProgramStudi(page, limit int) ([]models.ProgramStudi, int64, error) {
	var programStudi []models.ProgramStudi
	var total int64

	query := s.db.Model(&models.ProgramStudi{})

	// Count total
	query.Count(&total)

	// Apply pagination
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&programStudi).Error; err != nil {
		return nil, 0, err
	}

	return programStudi, total, nil
}

func (s *AcademicService) GetMataKuliah(prodiID string, page, limit int) ([]models.MataKuliah, int64, error) {
	var mataKuliah []models.MataKuliah
	var total int64

	query := s.db.Model(&models.MataKuliah{}).Preload("ProgramStudi")
	
	if prodiID != "" {
		query = query.Where("prodi_id = ?", prodiID)
	}

	// Count total
	query.Count(&total)

	// Apply pagination
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&mataKuliah).Error; err != nil {
		return nil, 0, err
	}

	return mataKuliah, total, nil
}

func (s *AcademicService) GetKonversiNilai(page, limit int) ([]models.ApplyJob, int64, error) {
	var applyJobs []models.ApplyJob
	var total int64

	query := s.db.Model(&models.ApplyJob{}).
		Where("status = ?", "Selesai").
		Preload("User").
		Preload("User.ProfilePicture").
		Preload("Job").
		Preload("Job.JobVacancyImage").
		Preload("KonversiNilai").
		Preload("KonversiNilai.MataKuliah")

	// Count total
	query.Count(&total)

	// Apply pagination
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&applyJobs).Error; err != nil {
		return nil, 0, err
	}

	return applyJobs, total, nil
}

func (s *AcademicService) GetKonversiNilaiByID(id uuid.UUID) (*models.ApplyJob, error) {
	var applyJob models.ApplyJob
	if err := s.db.Preload("User").
		Preload("User.ProfilePicture").
		Preload("Job").
		Preload("Job.JobVacancyImage").
		Preload("KonversiNilai").
		Preload("KonversiNilai.MataKuliah").
		First(&applyJob, id).Error; err != nil {
		return nil, err
	}
	return &applyJob, nil
}

func (s *AcademicService) CreateKonversiNilai(konversiData map[string]interface{}) (*models.KonversiNilai, error) {
	konversi := models.KonversiNilai{
		ApplyJobID:   uuid.MustParse(konversiData["apply_job_id"].(string)),
		MataKuliahID: uuid.MustParse(konversiData["mata_kuliah_id"].(string)),
		Grade:        konversiData["grade"].(string),
		Score:        konversiData["score"].(int),
	}

	if err := s.db.Create(&konversi).Error; err != nil {
		return nil, err
	}

	return &konversi, nil
}

func (s *AcademicService) GetBobotNilai() (*models.BobotNilai, error) {
	var bobotNilai models.BobotNilai
	if err := s.db.First(&bobotNilai).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create default bobot nilai
			bobotNilai = models.BobotNilai{
				BobotNilaiPerusahaan: 40,
				BobotNilaiPembimbing: 30,
				BobotNilaiPenguji:    30,
			}
			s.db.Create(&bobotNilai)
		} else {
			return nil, err
		}
	}
	return &bobotNilai, nil
}

func (s *AcademicService) UpdateBobotNilai(bobotData map[string]interface{}) (*models.BobotNilai, error) {
	var bobotNilai models.BobotNilai
	
	// Get existing or create new
	if err := s.db.First(&bobotNilai).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			bobotNilai = models.BobotNilai{}
			s.db.Create(&bobotNilai)
		} else {
			return nil, err
		}
	}

	updates := make(map[string]interface{})
	if val, ok := bobotData["bobot_nilai_perusahaan"]; ok {
		updates["bobot_nilai_perusahaan"] = val
	}
	if val, ok := bobotData["bobot_nilai_pembimbing"]; ok {
		updates["bobot_nilai_pembimbing"] = val
	}
	if val, ok := bobotData["bobot_nilai_penguji"]; ok {
		updates["bobot_nilai_penguji"] = val
	}

	if err := s.db.Model(&bobotNilai).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &bobotNilai, nil
}