package cmd

import (
	"log"

	"github.com/MohamedMosalm/To-Do-List/config"
	"github.com/MohamedMosalm/To-Do-List/models"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	r := gin.Default()

	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to the database: %v\n", err)
	}

	log.Println("Connected to the database")

	if err := db.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Fatalf("database migration failed: %v\n", err)
	}

	if err := r.Run(config.ServerPort); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}

	if err := r.Run(config.ServerPort); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
