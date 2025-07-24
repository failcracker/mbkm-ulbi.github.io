package routes

import (
	"mbkm-ulbi-backend/config"
	"mbkm-ulbi-backend/controllers"
	"mbkm-ulbi-backend/middleware"
	"mbkm-ulbi-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	// Initialize services
	userService := services.NewUserService(db)
	authService := services.NewAuthService(db, cfg)
	jobService := services.NewJobService(db)
	applyJobService := services.NewApplyJobService(db)
	reportService := services.NewReportService(db)
	evaluationService := services.NewEvaluationService(db)
	monthlyLogService := services.NewMonthlyLogService(db)
	academicService := services.NewAcademicService(db)
	dashboardService := services.NewDashboardService(db)
	fileService := services.NewFileService(cfg)

	// Initialize controllers
	authController := controllers.NewAuthController(authService, userService, fileService)
	userController := controllers.NewUserController(userService)
	jobController := controllers.NewJobController(jobService, fileService)
	applyJobController := controllers.NewApplyJobController(applyJobService, fileService)
	reportController := controllers.NewReportController(reportService, fileService)
	evaluationController := controllers.NewEvaluationController(evaluationService)
	monthlyLogController := controllers.NewMonthlyLogController(monthlyLogService)
	academicController := controllers.NewAcademicController(academicService)
	dashboardController := controllers.NewDashboardController(dashboardService)

	// Public routes
	api := r.Group("/api/v1")
	{
		api.POST("/login", authController.Login)
		api.POST("/register", authController.Register)
	}

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware(cfg.JWT.Secret))
	{
		// Profile routes
		protected.GET("/profile", userController.GetProfile)
		protected.PUT("/profile", userController.UpdateProfile)

		// User management routes
		protected.GET("/users/lecturer", userController.GetLecturers)
		protected.GET("/users/student", userController.GetStudents)

		// Role routes
		protected.GET("/roles", userController.GetRoles)
		protected.POST("/roles/assign", userController.AssignRole)

		// Company routes
		protected.GET("/companies", jobController.GetCompanies)

		// Job routes
		protected.GET("/jobs", jobController.GetJobs)
		protected.GET("/jobs/:id", jobController.GetJobByID)
		protected.POST("/jobs", jobController.CreateJob)
		protected.PUT("/jobs/:id", jobController.UpdateJob)
		protected.DELETE("/jobs/:id", jobController.DeleteJob)
		protected.POST("/jobs/:id/approve", jobController.ApproveJob)
		protected.POST("/jobs/:id/reject", jobController.RejectJob)
		protected.GET("/jobs/:id/list", jobController.GetJobCandidates)

		// Apply job routes
		protected.GET("/apply-jobs", applyJobController.GetApplyJobs)
		protected.GET("/apply-jobs/:id", applyJobController.GetApplyJobByID)
		protected.GET("/apply-jobs/user/:user_id", applyJobController.GetApplyJobsByUser)
		protected.GET("/apply-jobs/user/:user_id/last", applyJobController.GetLastApplyJobByUser)
		protected.POST("/apply-jobs", applyJobController.CreateApplyJob)
		protected.POST("/apply-jobs/:id/approve", applyJobController.ApproveApplyJob)
		protected.POST("/apply-jobs/:id/reject", applyJobController.RejectApplyJob)
		protected.POST("/apply-jobs/:id/activate", applyJobController.ActivateApplyJob)
		protected.POST("/apply-jobs/:id/done", applyJobController.DoneApplyJob)
		protected.POST("/apply-jobs/:id/set-lecturer", applyJobController.SetLecturer)

		// Monthly logs routes
		protected.GET("/apply-jobs/monthly-logs", monthlyLogController.GetMonthlyLogs)
		protected.GET("/apply-jobs/monthly-logs/:id", monthlyLogController.GetMonthlyLogByID)
		protected.POST("/apply-jobs/monthly-logs", monthlyLogController.CreateMonthlyLog)
		protected.PUT("/apply-jobs/monthly-logs/update", monthlyLogController.UpdateMonthlyLog)
		protected.POST("/apply-jobs/monthly-logs/:id/approve", monthlyLogController.ApproveMonthlyLog)
		protected.POST("/apply-jobs/monthly-logs/:id/revision", monthlyLogController.RevisionMonthlyLog)

		// Report routes
		protected.GET("/reports", reportController.GetReports)
		protected.GET("/reports/:id", reportController.GetReportByID)
		protected.POST("/reports", reportController.CreateReport)
		protected.POST("/reports/:id/check", reportController.CheckReport)

		// Evaluation routes
		protected.GET("/evaluations", evaluationController.GetEvaluations)
		protected.GET("/evaluations/:id", evaluationController.GetEvaluationByID)
		protected.POST("/evaluations", evaluationController.CreateEvaluation)

		// Academic routes
		protected.GET("/program-studi", academicController.GetProgramStudi)
		protected.GET("/mata-kuliah", academicController.GetMataKuliah)
		protected.GET("/konversi-nilai", academicController.GetKonversiNilai)
		protected.GET("/konversi-nilai/:id", academicController.GetKonversiNilaiByID)
		protected.POST("/konversi-nilai", academicController.CreateKonversiNilai)

		// Settings routes
		protected.GET("/settings/bobot-nilai", academicController.GetBobotNilai)
		protected.POST("/settings/bobot-nilai", academicController.UpdateBobotNilai)

		// Dashboard routes
		protected.GET("/dashboard/overview", dashboardController.GetOverview)
	}
}