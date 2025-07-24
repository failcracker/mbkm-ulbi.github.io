package controllers

import (
	"net/http"
	"strconv"
	"mbkm-ulbi-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AcademicController struct {
	academicService *services.AcademicService
}

func NewAcademicController(academicService *services.AcademicService) *AcademicController {
	return &AcademicController{academicService: academicService}
}

func (ctrl *AcademicController) GetProgramStudi(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	programStudi, total, err := ctrl.academicService.GetProgramStudi(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Program studi retrieved successfully",
		"data": gin.H{
			"data":  programStudi,
			"count": total,
			"page":  page,
			"limit": limit,
		},
	})
}

func (ctrl *AcademicController) GetMataKuliah(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	prodiID := c.Query("prodi_id")

	mataKuliah, total, err := ctrl.academicService.GetMataKuliah(prodiID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Mata kuliah retrieved successfully",
		"data": gin.H{
			"data":  mataKuliah,
			"count": total,
			"page":  page,
			"limit": limit,
		},
	})
}

func (ctrl *AcademicController) GetKonversiNilai(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	konversiNilai, total, err := ctrl.academicService.GetKonversiNilai(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Konversi nilai retrieved successfully",
		"data": gin.H{
			"data":  konversiNilai,
			"count": total,
			"page":  page,
			"limit": limit,
		},
	})
}

func (ctrl *AcademicController) GetKonversiNilaiByID(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))

	konversiNilai, err := ctrl.academicService.GetKonversiNilaiByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Konversi nilai not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Konversi nilai retrieved successfully",
		"data":    konversiNilai,
	})
}

func (ctrl *AcademicController) CreateKonversiNilai(c *gin.Context) {
	konversiData := make(map[string]interface{})
	
	// Parse form data
	konversiData["apply_job_id"] = c.PostForm("apply_job_id")
	konversiData["mata_kuliah_id"] = c.PostForm("mata_kuliah_id")
	konversiData["grade"] = c.PostForm("grade")
	
	score, _ := strconv.Atoi(c.PostForm("score"))
	konversiData["score"] = score

	konversi, err := ctrl.academicService.CreateKonversiNilai(konversiData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Konversi nilai created successfully",
		"data":    konversi,
	})
}

func (ctrl *AcademicController) GetBobotNilai(c *gin.Context) {
	bobotNilai, err := ctrl.academicService.GetBobotNilai()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Bobot nilai retrieved successfully",
		"data":    bobotNilai,
	})
}

func (ctrl *AcademicController) UpdateBobotNilai(c *gin.Context) {
	bobotData := make(map[string]interface{})
	
	// Parse form data
	if val := c.PostForm("bobot_nilai_perusahaan"); val != "" {
		if intVal, err := strconv.Atoi(val); err == nil {
			bobotData["bobot_nilai_perusahaan"] = intVal
		}
	}
	if val := c.PostForm("bobot_nilai_pembimbing"); val != "" {
		if intVal, err := strconv.Atoi(val); err == nil {
			bobotData["bobot_nilai_pembimbing"] = intVal
		}
	}
	if val := c.PostForm("bobot_nilai_penguji"); val != "" {
		if intVal, err := strconv.Atoi(val); err == nil {
			bobotData["bobot_nilai_penguji"] = intVal
		}
	}

	bobotNilai, err := ctrl.academicService.UpdateBobotNilai(bobotData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Bobot nilai updated successfully",
		"data":    bobotNilai,
	})
}