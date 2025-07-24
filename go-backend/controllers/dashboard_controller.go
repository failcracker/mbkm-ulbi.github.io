package controllers

import (
	"net/http"
	"mbkm-ulbi-backend/services"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	dashboardService *services.DashboardService
}

func NewDashboardController(dashboardService *services.DashboardService) *DashboardController {
	return &DashboardController{dashboardService: dashboardService}
}

func (ctrl *DashboardController) GetOverview(c *gin.Context) {
	overview, err := ctrl.dashboardService.GetOverview()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, overview)
}