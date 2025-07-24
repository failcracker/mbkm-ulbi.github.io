package controllers

import (
	"net/http"
	"strconv"
	"mbkm-ulbi-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MonthlyLogController struct {
	monthlyLogService *services.MonthlyLogService
}

func NewMonthlyLogController(monthlyLogService *services.MonthlyLogService) *MonthlyLogController {
	return &MonthlyLogController{monthlyLogService: monthlyLogService}
}

func (ctrl *MonthlyLogController) GetMonthlyLogs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	logs, total, err := ctrl.monthlyLogService.GetMonthlyLogs(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Monthly logs retrieved successfully",
		"data": gin.H{
			"data":  logs,
			"count": total,
			"page":  page,
			"limit": limit,
		},
	})
}

func (ctrl *MonthlyLogController) GetMonthlyLogByID(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))

	log, err := ctrl.monthlyLogService.GetMonthlyLogByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monthly log not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Monthly log retrieved successfully",
		"data":    log,
	})
}

func (ctrl *MonthlyLogController) CreateMonthlyLog(c *gin.Context) {
	logData := make(map[string]interface{})
	
	// Parse form data
	logData["apply_job_id"] = c.PostForm("apply_job_id")
	logData["start_date"] = c.PostForm("start_date")
	logData["end_date"] = c.PostForm("end_date")
	logData["content"] = c.PostForm("content")
	logData["hasil"] = c.PostForm("hasil")

	log, err := ctrl.monthlyLogService.CreateMonthlyLog(logData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Monthly log created successfully",
		"data":    log,
	})
}

func (ctrl *MonthlyLogController) UpdateMonthlyLog(c *gin.Context) {
	logData := make(map[string]interface{})
	
	// Parse form data
	logData["id"] = c.PostForm("id")
	logData["start_date"] = c.PostForm("start_date")
	logData["end_date"] = c.PostForm("end_date")
	logData["content"] = c.PostForm("content")
	logData["hasil"] = c.PostForm("hasil")

	if err := ctrl.monthlyLogService.UpdateMonthlyLog(logData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Monthly log updated successfully"})
}

func (ctrl *MonthlyLogController) ApproveMonthlyLog(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))

	if err := ctrl.monthlyLogService.ApproveMonthlyLog(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Monthly log approved successfully"})
}

func (ctrl *MonthlyLogController) RevisionMonthlyLog(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))
	feedback := c.PostForm("feedback")

	if err := ctrl.monthlyLogService.RevisionMonthlyLog(id, feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Monthly log revision sent successfully"})
}