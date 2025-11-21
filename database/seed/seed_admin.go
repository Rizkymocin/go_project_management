package seed

import (
	"log"

	"github.com/Rizkymocin/project-management/config"
	"github.com/Rizkymocin/project-management/models"
	"github.com/Rizkymocin/project-management/utils"
	"github.com/google/uuid"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("admin123")

	admin := models.User{
		Name:     "admin",
		PublicID: uuid.New(),
		Email:    "admin@pm.dev",
		Password: password,
		Role:     "admin",
	}
	if err := config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Println("failed to seed admin", err)
	} else {
		log.Println("Admin Users seeded")
	}

}
