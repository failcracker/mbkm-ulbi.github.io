package controllers

import (
	"net/http"
	"strconv"
	"mbkm-ulbi-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReportController struct {
	reportService *services.ReportService
	fileService   *services.FileService
}

func NewReportController(reportService *services.ReportService, fileService *services.FileService) *ReportController {
	return &ReportController{
		reportService: reportService,
		fileService:   fileService,
	}
}

func (ctrl *ReportController) GetReports(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	reports, total, err := ctrl.reportService.GetReports(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Reports retrieved successfully",
		"data": gin.H{
			"data":  reports,
			"count": total,
			"page":  page,
			"limit": limit,
		},
	})
}

func (ctrl *ReportController) GetReportByID(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))

	report, err := ctrl.reportService.GetReportByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Report not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Report retrieved successfully",
		"data":    report,
	})
}

func (ctrl *ReportController) CreateReport(c *gin.Context) {
	reportData := make(map[string]interface{})
	
	// Parse form data
	reportData["apply_job_id"] = c.PostForm("apply_job_id")

	report, err := ctrl.reportService.CreateReport(reportData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Report created successfully",
		"data":    report,
	})
}

func (ctrl *ReportController) CheckReport(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")
	
	checkerID := uuid.MustParse(userID.(string))
	roleStr := role.(string)

	if err := ctrl.reportService.CheckReport(id, checkerID, roleStr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Report checked successfully"})
}