package services

import (
	"errors"
	"mbkm-ulbi-backend/config"
	"mbkm-ulbi-backend/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewAuthService(db *gorm.DB, cfg *config.Config) *AuthService {
	return &AuthService{db: db, cfg: cfg}
}

func (s *AuthService) Login(username, password string) (*models.User, string, error) {
	var user models.User
	if err := s.db.Preload("Roles").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    s.getUserRole(&user),
		"exp":     time.Now().Add(time.Hour * time.Duration(s.cfg.JWT.ExpireHours)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.cfg.JWT.Secret))
	if err != nil {
		return nil, "", err
	}

	return &user, tokenString, nil
}

func (s *AuthService) Register(userData map[string]interface{}) (*models.User, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData["password"].(string)), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Username: userData["username"].(string),
		Email:    userData["email"].(string),
		Password: string(hashedPassword),
	}

	// Set user fields based on role
	if role, ok := userData["role"].(string); ok {
		switch role {
		case "student":
			user.Name = userData["name"].(string)
			user.NIM = userData["nim"].(string)
			user.ProgramStudy = userData["program_study"].(string)
			user.Faculty = userData["faculty"].(string)
			if semester, ok := userData["semester"].(string); ok {
				// Convert semester string to int if needed
				user.Semester = 1 // Default or parse from string
			}
			user.SocialMedia = userData["social_media"].(string)
			user.PhoneNumber = userData["phone_number"].(string)
			user.ProfileDescription = userData["deskripsi"].(string)
		case "company":
			user.Name = userData["name"].(string)
			user.Position = userData["position"].(string)
			user.CompanyName = userData["company_name"].(string)
			user.BusinessField = userData["business_field"].(string)
			user.CompanySize = userData["company_size"].(string)
			user.CompanyAddress = userData["company_address"].(string)
			user.CompanyWebsite = userData["company_website"].(string)
			user.CompanyPhoneNumber = userData["company_phone_number"].(string)
			user.CompanyDescription = userData["company_description"].(string)
		}
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	// Assign default role
	if role, ok := userData["role"].(string); ok {
		var roleModel models.Role
		if err := s.db.Where("title = ?", role).First(&roleModel).Error; err == nil {
			userRole := models.UserRole{
				UserID: user.ID,
				RoleID: roleModel.ID,
			}
			s.db.Create(&userRole)
		}
	}

	return &user, nil
}

func (s *AuthService) getUserRole(user *models.User) string {
	if len(user.Roles) > 0 {
		return user.Roles[0].Title
	}
	return "user"
}