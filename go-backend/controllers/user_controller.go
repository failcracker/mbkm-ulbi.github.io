package controllers

import (
	"net/http"
	"strconv"
	"mbkm-ulbi-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (ctrl *UserController) GetProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := uuid.MustParse(userID.(string))

	user, err := ctrl.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile retrieved successfully",
		"user":    user,
	})
}

func (ctrl *UserController) UpdateProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := uuid.MustParse(userID.(string))

	updates := make(map[string]interface{})
	// Parse form data for updates
	if name := c.PostForm("name"); name != "" {
		updates["name"] = name
	}
	if email := c.PostForm("email"); email != "" {
		updates["email"] = email
	}
	// Add more fields as needed

	if err := ctrl.userService.UpdateProfile(id, updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

func (ctrl *UserController) GetLecturers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	applyJobID := c.Query("apply_job_id")

	lecturers, total, err := ctrl.userService.GetLecturers(page, limit, applyJobID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lecturers retrieved successfully",
		"data": gin.H{
			"data":  lecturers,
			"count": total,
			"page":  page,
			"limit": limit,
		},
	})
}

func (ctrl *UserController) GetStudents(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	students, total, err := ctrl.userService.GetStudents(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Students retrieved successfully",
		"data": gin.H{
			"data":  students,
			"count": total,
			"page":  page,
			"limit": limit,
		},
	})
}

func (ctrl *UserController) GetRoles(c *gin.Context) {
	roles, err := ctrl.userService.GetRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Roles retrieved successfully",
		"data": gin.H{
			"data": roles,
		},
	})
}

func (ctrl *UserController) AssignRole(c *gin.Context) {
	userID := uuid.MustParse(c.PostForm("user_id"))
	roleID := uuid.MustParse(c.PostForm("role_id"))

	if err := ctrl.userService.AssignRole(userID, roleID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role assigned successfully"})
}