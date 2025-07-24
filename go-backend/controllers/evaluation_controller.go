package controllers

import (
	"net/http"
	"strconv"
	"mbkm-ulbi-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EvaluationController struct {
	evaluationService *services.EvaluationService
}

func NewEvaluationController(evaluationService *services.EvaluationService) *EvaluationController {
	return &EvaluationController{evaluationService: evaluationService}
}

func (ctrl *EvaluationController) GetEvaluations(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	evaluations, total, err := ctrl.evaluationService.GetEvaluations(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Evaluations retrieved successfully",
		"data": gin.H{
			"data":  evaluations,
			"count": total,
			"page":  page,
			"limit": limit,
		},
	})
}

func (ctrl *EvaluationController) GetEvaluationByID(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))

	evaluation, err := ctrl.evaluationService.GetEvaluationByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Evaluation not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Evaluation retrieved successfully",
		"data":    evaluation,
	})
}

func (ctrl *EvaluationController) CreateEvaluation(c *gin.Context) {
	evaluationData := make(map[string]interface{})
	
	// Parse form data
	evaluationData["apply_job_id"] = c.PostForm("apply_job_id")
	evaluationData["grade"] = c.PostForm("grade")
	evaluationData["grade_score"] = c.PostForm("grade_score")
	evaluationData["grade_description"] = c.PostForm("grade_description")
	evaluationData["is_examiner"] = c.PostForm("is_examiner")

	evaluation, err := ctrl.evaluationService.CreateEvaluation(evaluationData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Evaluation created successfully",
		"data":    evaluation,
	})
}