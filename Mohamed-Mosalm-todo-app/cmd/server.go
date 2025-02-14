package cmd

import (
	"log"

	"github.com/MohamedMosalm/Todo-App/cmd/api/handlers"
	"github.com/MohamedMosalm/Todo-App/cmd/api/routes"
	"github.com/MohamedMosalm/Todo-App/config"
	"github.com/MohamedMosalm/Todo-App/database"
	"github.com/MohamedMosalm/Todo-App/models"
	taskRepository "github.com/MohamedMosalm/Todo-App/repositories/taskRepository"
	userRepository "github.com/MohamedMosalm/Todo-App/repositories/userRepository"
	"github.com/MohamedMosalm/Todo-App/services"
	"github.com/gin-gonic/gin"
)

func StartServer(config config.AppConfig) {
	r := gin.Default()

	db, err := database.ConnectDB(config.DSN)
	if err != nil {
		log.Fatalf("could not connect to the database: %v\n", err)
	}

	if err := database.AutoMigrate(db, &models.User{}, &models.Task{}); err != nil {
		log.Fatalf("database migration failed: %v\n", err)
	}

	taskRepo := taskRepository.NewGormTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService, config)

	userRepo := userRepository.NewGormUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler, err := handlers.NewAuthHandler(userService, config)
	if err != nil {
		log.Fatalf("Failed to create auth handler: %v", err)
	}

	routes.SetupAuthRoutes(r, userHandler)
	routes.SetupTaskRoutes(r, taskHandler, config.JWTSecret)

	if err := r.Run(config.ServerPort); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
