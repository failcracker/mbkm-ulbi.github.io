package services

import (
	"context"
	"mbkm-ulbi-backend/config"
	"mbkm-ulbi-backend/models"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileService struct {
	cfg *config.Config
	db  *gorm.DB
	cld *cloudinary.Cloudinary
}

func NewFileService(cfg *config.Config) *FileService {
	cld, _ := cloudinary.NewFromParams(cfg.Cloudinary.CloudName, cfg.Cloudinary.APIKey, cfg.Cloudinary.APISecret)
	return &FileService{
		cfg: cfg,
		cld: cld,
	}
}

func (s *FileService) UploadFile(file *multipart.FileHeader, userID uuid.UUID) (*models.File, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// Upload to Cloudinary
	result, err := s.cld.Upload.Upload(context.Background(), src, uploader.UploadParams{
		Folder: "mbkm-ulbi",
	})
	if err != nil {
		return nil, err
	}

	// Save file info to database
	fileModel := models.File{
		Name:       file.Filename,
		URL:        result.SecureURL,
		PublicID:   result.PublicID,
		Size:       file.Size,
		MimeType:   file.Header.Get("Content-Type"),
		UploadedBy: userID,
	}

	if err := s.db.Create(&fileModel).Error; err != nil {
		return nil, err
	}

	return &fileModel, nil
}