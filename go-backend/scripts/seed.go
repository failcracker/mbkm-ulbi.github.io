package main

import (
	"log"
	"mbkm-ulbi-backend/config"
	"mbkm-ulbi-backend/database"
	"mbkm-ulbi-backend/models"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize configuration
	cfg := config.Load()

	// Initialize database
	db := database.Init(cfg)

	// Seed roles
	roles := []models.Role{
		{Title: "superadmin", Description: "Super Administrator"},
		{Title: "cdc", Description: "Career Development Center"},
		{Title: "prodi", Description: "Program Studi"},
		{Title: "dosen", Description: "Dosen/Lecturer"},
		{Title: "mahasiswa", Description: "Student"},
		{Title: "mitra", Description: "Company Partner"},
	}

	for _, role := range roles {
		db.FirstOrCreate(&role, models.Role{Title: role.Title})
	}

	// Seed default admin user
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	admin := models.User{
		Username: "admin",
		Email:    "admin@ulbi.ac.id",
		Password: string(hashedPassword),
		Name:     "Administrator",
		Status:   "Aktif",
	}

	var existingAdmin models.User
	if err := db.Where("username = ?", admin.Username).First(&existingAdmin).Error; err != nil {
		db.Create(&admin)
		
		// Assign superadmin role
		var superadminRole models.Role
		db.Where("title = ?", "superadmin").First(&superadminRole)
		
		userRole := models.UserRole{
			UserID: admin.ID,
			RoleID: superadminRole.ID,
		}
		db.Create(&userRole)
	}

	// Seed program studi
	programStudi := []models.ProgramStudi{
		{KodeProgramStudi: "D3-AK", NamaProgramStudi: "D3 Akuntansi", Fakultas: "Ekonomi", Jenjang: "D3"},
		{KodeProgramStudi: "D3-MB", NamaProgramStudi: "D3 Manajemen Bisnis", Fakultas: "Ekonomi", Jenjang: "D3"},
		{KodeProgramStudi: "D3-ML", NamaProgramStudi: "D3 Manajemen Logistik", Fakultas: "Logistik", Jenjang: "D3"},
		{KodeProgramStudi: "D3-SI", NamaProgramStudi: "D3 Sistem Informasi", Fakultas: "Teknologi", Jenjang: "D3"},
		{KodeProgramStudi: "D3-TI", NamaProgramStudi: "D3 Teknik Informatika", Fakultas: "Teknologi", Jenjang: "D3"},
		{KodeProgramStudi: "S1-SD", NamaProgramStudi: "S1 Sains Data", Fakultas: "Teknologi", Jenjang: "S1"},
		{KodeProgramStudi: "S1-MT", NamaProgramStudi: "S1 Manajemen Transportasi", Fakultas: "Logistik", Jenjang: "S1"},
	}

	for _, prodi := range programStudi {
		db.FirstOrCreate(&prodi, models.ProgramStudi{KodeProgramStudi: prodi.KodeProgramStudi})
	}

	// Seed default bobot nilai
	bobotNilai := models.BobotNilai{
		BobotNilaiPerusahaan: 40,
		BobotNilaiPembimbing: 30,
		BobotNilaiPenguji:    30,
	}
	db.FirstOrCreate(&bobotNilai)

	log.Println("Database seeded successfully!")
}