package database

import (
	"fmt"
	"log"
	"mbkm-ulbi-backend/config"
	"mbkm-ulbi-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")
	return db
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.UserRole{},
		&models.Company{},
		&models.Job{},
		&models.ApplyJob{},
		&models.Report{},
		&models.Evaluation{},
		&models.MonthlyLog{},
		&models.ProgramStudi{},
		&models.MataKuliah{},
		&models.KonversiNilai{},
		&models.BobotNilai{},
		&models.File{},
	)
}