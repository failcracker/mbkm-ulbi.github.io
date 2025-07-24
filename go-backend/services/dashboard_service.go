package services

import (
	"mbkm-ulbi-backend/models"

	"gorm.io/gorm"
)

type DashboardService struct {
	db *gorm.DB
}

type DashboardOverview struct {
	TotalCompany      int64                    `json:"total_company"`
	TotalJob          int64                    `json:"total_job"`
	TotalStudent      int64                    `json:"total_student"`
	TotalAktifMagang  int64                    `json:"total_aktif_magang"`
	ChartData         ChartData                `json:"chart_data"`
	LatestData        LatestData               `json:"latest_data"`
}

type ChartData struct {
	Labels   []string    `json:"labels"`
	Datasets []Dataset   `json:"datasets"`
}

type Dataset struct {
	Label string `json:"label"`
	Data  []int  `json:"data"`
}

type LatestData struct {
	Jobs              []models.Job      `json:"jobs"`
	Companies         []models.Company  `json:"companies"`
	ApplyJobStudents  []models.ApplyJob `json:"apply_job_students"`
}

func NewDashboardService(db *gorm.DB) *DashboardService {
	return &DashboardService{db: db}
}

func (s *DashboardService) GetOverview() (*DashboardOverview, error) {
	overview := &DashboardOverview{}

	// Count totals
	s.db.Model(&models.Company{}).Count(&overview.TotalCompany)
	s.db.Model(&models.Job{}).Count(&overview.TotalJob)
	s.db.Model(&models.User{}).
		Joins("JOIN user_roles ON users.id = user_roles.user_id").
		Joins("JOIN roles ON user_roles.role_id = roles.id").
		Where("roles.title = ?", "mahasiswa").
		Count(&overview.TotalStudent)
	s.db.Model(&models.ApplyJob{}).Where("status = ?", "Aktif").Count(&overview.TotalAktifMagang)

	// Get chart data (simplified)
	overview.ChartData = ChartData{
		Labels: []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		Datasets: []Dataset{
			{
				Label: "Melamar",
				Data:  []int{10, 20, 30, 40, 23, 60, 100, 80, 90, 22, 110, 212},
			},
			{
				Label: "Disetujui",
				Data:  []int{5, 15, 25, 15, 45, 30, 65, 75, 100, 80, 105, 200},
			},
			{
				Label: "Aktif",
				Data:  []int{3, 10, 20, 10, 35, 25, 55, 65, 85, 70, 95, 180},
			},
			{
				Label: "Selesai",
				Data:  []int{2, 8, 15, 8, 30, 20, 45, 55, 75, 60, 85, 160},
			},
		},
	}

	// Get latest data
	var jobs []models.Job
	s.db.Preload("JobVacancyImage").Limit(5).Order("created_at DESC").Find(&jobs)
	overview.LatestData.Jobs = jobs

	var companies []models.Company
	s.db.Preload("CompanyLogo").Limit(5).Order("created_at DESC").Find(&companies)
	overview.LatestData.Companies = companies

	var applyJobs []models.ApplyJob
	s.db.Preload("User").Preload("User.ProfilePicture").Limit(5).Order("created_at DESC").Find(&applyJobs)
	overview.LatestData.ApplyJobStudents = applyJobs

	return overview, nil
}