package controllers

import (
	"net/http"
	"strconv"
	"mbkm-ulbi-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type JobController struct {
	jobService  *services.JobService
	fileService *services.FileService
}

func NewJobController(jobService *services.JobService, fileService *services.FileService) *JobController {
	return &JobController{
		jobService:  jobService,
		fileService: fileService,
	}
}

func (ctrl *JobController) GetJobs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	jobs, total, err := ctrl.jobService.GetJobs(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Jobs retrieved successfully",
		"data": gin.H{
			"data":  jobs,
			"count": total,
			"page":  page,
			"limit": limit,
		},
	})
}

func (ctrl *JobController) GetJobByID(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))

	job, err := ctrl.jobService.GetJobByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Job retrieved successfully",
		"data":    job,
	})
}

func (ctrl *JobController) CreateJob(c *gin.Context) {
	jobData := make(map[string]interface{})
	
	// Parse form data
	jobData["title"] = c.PostForm("title")
	jobData["company"] = c.PostForm("company")
	jobData["location"] = c.PostForm("location")
	jobData["description"] = c.PostForm("description")
	jobData["duration"] = c.PostForm("duration")
	jobData["job_type"] = c.PostForm("job_type")
	jobData["benefits"] = c.PostForm("benefits")
	jobData["status"] = c.PostForm("status")
	jobData["vacancy_type"] = c.PostForm("vacancy_type")
	jobData["deadline"] = c.PostForm("deadline")

	job, err := ctrl.jobService.CreateJob(jobData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Job created successfully",
		"data":    job,
	})
}

func (ctrl *JobController) UpdateJob(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))
	
	updates := make(map[string]interface{})
	// Parse form data for updates
	if title := c.PostForm("title"); title != "" {
		updates["title"] = title
	}
	if company := c.PostForm("company"); company != "" {
		updates["company"] = company
	}
	// Add more fields as needed

	if err := ctrl.jobService.UpdateJob(id, updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job updated successfully"})
}

func (ctrl *JobController) DeleteJob(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))

	if err := ctrl.jobService.DeleteJob(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job deleted successfully"})
}

func (ctrl *JobController) ApproveJob(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))

	if err := ctrl.jobService.ApproveJob(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job approved successfully"})
}

func (ctrl *JobController) RejectJob(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))

	if err := ctrl.jobService.RejectJob(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job rejected successfully"})
}

func (ctrl *JobController) GetJobCandidates(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))

	candidates, err := ctrl.jobService.GetJobCandidates(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Job candidates retrieved successfully",
		"data":    candidates,
	})
}

func (ctrl *JobController) GetCompanies(c *gin.Context) {
	companies, err := ctrl.jobService.GetCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Companies retrieved successfully",
		"data": gin.H{
			"data": companies,
		},
	})
}