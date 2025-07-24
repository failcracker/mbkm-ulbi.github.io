package services

import (
	"mbkm-ulbi-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := s.db.Preload("Roles").Preload("ProfilePicture").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetLecturers(page, limit int, applyJobID string) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := s.db.Model(&models.User{}).
		Joins("JOIN user_roles ON users.id = user_roles.user_id").
		Joins("JOIN roles ON user_roles.role_id = roles.id").
		Where("roles.title IN ?", []string{"dosen", "lecturer"}).
		Preload("Roles").
		Preload("ProfilePicture")

	// Count total
	query.Count(&total)

	// Apply pagination
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (s *UserService) GetStudents(page, limit int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := s.db.Model(&models.User{}).
		Joins("JOIN user_roles ON users.id = user_roles.user_id").
		Joins("JOIN roles ON user_roles.role_id = roles.id").
		Where("roles.title IN ?", []string{"mahasiswa", "student"}).
		Preload("Roles").
		Preload("ProfilePicture")

	// Count total
	query.Count(&total)

	// Apply pagination
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (s *UserService) GetRoles() ([]models.Role, error) {
	var roles []models.Role
	if err := s.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (s *UserService) AssignRole(userID, roleID uuid.UUID) error {
	// Remove existing roles
	s.db.Where("user_id = ?", userID).Delete(&models.UserRole{})

	// Assign new role
	userRole := models.UserRole{
		UserID: userID,
		RoleID: roleID,
	}
	return s.db.Create(&userRole).Error
}

func (s *UserService) UpdateProfile(userID uuid.UUID, updates map[string]interface{}) error {
	return s.db.Model(&models.User{}).Where("id = ?", userID).Updates(updates).Error
}