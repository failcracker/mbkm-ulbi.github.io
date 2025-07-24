package controllers

import (
	"net/http"
	"mbkm-ulbi-backend/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
	userService *services.UserService
	fileService *services.FileService
}

func NewAuthController(authService *services.AuthService, userService *services.UserService, fileService *services.FileService) *AuthController {
	return &AuthController{
		authService: authService,
		userService: userService,
		fileService: fileService,
	}
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var loginData struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := ctrl.authService.Login(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Get user role
	role := "user"
	if len(user.Roles) > 0 {
		role = user.Roles[0].Title
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"data": gin.H{
			"token": token,
			"user":  user,
			"role":  role,
		},
	})
}

func (ctrl *AuthController) Register(c *gin.Context) {
	// Parse form data
	userData := make(map[string]interface{})
	
	// Basic fields
	userData["username"] = c.PostForm("username")
	userData["password"] = c.PostForm("password")
	userData["email"] = c.PostForm("email")
	userData["role"] = c.PostForm("role")
	userData["type"] = c.PostForm("type")

	// Role-specific fields
	if role := c.PostForm("role"); role == "student" {
		userData["nim"] = c.PostForm("nim")
		userData["name"] = c.PostForm("name")
		userData["program_study"] = c.PostForm("program_study")
		userData["faculty"] = c.PostForm("faculty")
		userData["semester"] = c.PostForm("semester")
		userData["social_media"] = c.PostForm("social_media")
		userData["phone_number"] = c.PostForm("phone_number")
		userData["deskripsi"] = c.PostForm("deskripsi")
	} else if role == "company" {
		userData["name"] = c.PostForm("name")
		userData["position"] = c.PostForm("position")
		userData["company_name"] = c.PostForm("company_name")
		userData["business_field"] = c.PostForm("business_field")
		userData["company_size"] = c.PostForm("company_size")
		userData["company_address"] = c.PostForm("company_address")
		userData["company_website"] = c.PostForm("company_website")
		userData["company_phone_number"] = c.PostForm("company_phone_number")
		userData["company_description"] = c.PostForm("company_description")
	}

	user, err := ctrl.authService.Register(userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful",
		"data":    user,
	})
}