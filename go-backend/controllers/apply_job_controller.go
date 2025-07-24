package controllers

import (
	"net/http"
	"strconv"
	"mbkm-ulbi-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ApplyJobController struct {
	applyJobService *services.ApplyJobService
	fileService     *services.FileService
}

func NewApplyJobController(applyJobService *services.ApplyJobService, fileService *services.FileService) *ApplyJobController {
	return &ApplyJobController{
		applyJobService: applyJobService,
		fileService:     fileService,
	}
}

func (ctrl *ApplyJobController) GetApplyJobs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	applyJobs, total, err := ctrl.applyJobService.GetApplyJobs(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Apply jobs retrieved successfully",
		"data": gin.H{
			"data":  applyJobs,
			"count": total,
			"page":  page,
			"limit": limit,
		},
	})
}

func (ctrl *ApplyJobController) GetApplyJobByID(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))

	applyJob, err := ctrl.applyJobService.GetApplyJobByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Apply job not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Apply job retrieved successfully",
		"data":    applyJob,
	})
}

func (ctrl *ApplyJobController) GetApplyJobsByUser(c *gin.Context) {
	userID := uuid.MustParse(c.Param("user_id"))

	applyJobs, err := ctrl.applyJobService.GetApplyJobsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User apply jobs retrieved successfully",
		"data":    applyJobs,
	})
}

func (ctrl *ApplyJobController) GetLastApplyJobByUser(c *gin.Context) {
	userID := uuid.MustParse(c.Param("user_id"))

	applyJob, err := ctrl.applyJobService.GetLastApplyJobByUser(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No active apply job found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Last apply job retrieved successfully",
		"data":    applyJob,
	})
}

func (ctrl *ApplyJobController) CreateApplyJob(c *gin.Context) {
	applyJobData := make(map[string]interface{})
	
	// Parse form data
	applyJobData["user_id"] = c.PostForm("users[]")
	applyJobData["job_id"] = c.PostForm("jobs[]")
	applyJobData["email"] = c.PostForm("email")
	applyJobData["phone_number"] = c.PostForm("telepon")
	applyJobData["address"] = c.PostForm("alamat")

	applyJob, err := ctrl.applyJobService.CreateApplyJob(applyJobData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Apply job created successfully",
		"data":    applyJob,
	})
}

func (ctrl *ApplyJobController) ApproveApplyJob(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))

	if err := ctrl.applyJobService.ApproveApplyJob(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Apply job approved successfully"})
}

func (ctrl *ApplyJobController) RejectApplyJob(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))

	if err := ctrl.applyJobService.RejectApplyJob(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Apply job rejected successfully"})
}

func (ctrl *ApplyJobController) ActivateApplyJob(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))

	if err := ctrl.applyJobService.ActivateApplyJob(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Apply job activated successfully"})
}

func (ctrl *ApplyJobController) DoneApplyJob(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))

	if err := ctrl.applyJobService.DoneApplyJob(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Apply job completed successfully"})
}

func (ctrl *ApplyJobController) SetLecturer(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))
	
	var lecturerID, examinerID *uuid.UUID
	
	if lecturer := c.PostForm("lecturer_id"); lecturer != "" {
		lid := uuid.MustParse(lecturer)
		lecturerID = &lid
	}
	
	if examiner := c.PostForm("examiner_id"); examiner != "" {
		eid := uuid.MustParse(examiner)
		examinerID = &eid
	}

	if err := ctrl.applyJobService.SetLecturer(id, lecturerID, examinerID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lecturer assigned successfully"})
}