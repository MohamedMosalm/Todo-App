package cmd

import (
	"log"

	"github.com/MohamedMosalm/To-Do-List/cmd/api/handlers"
	"github.com/MohamedMosalm/To-Do-List/cmd/api/routes"
	"github.com/MohamedMosalm/To-Do-List/config"
	"github.com/MohamedMosalm/To-Do-List/database"
	"github.com/MohamedMosalm/To-Do-List/models"
	taskRepository "github.com/MohamedMosalm/To-Do-List/repositories/taskRepository"
	userRepository "github.com/MohamedMosalm/To-Do-List/repositories/userRepository"
	"github.com/MohamedMosalm/To-Do-List/services"
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
	taskHandler := handlers.NewTaskHandler(taskService)

	routes.SetupTaskRoutes(r, taskHandler)

	userRepo := userRepository.NewGormUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewAuthHandler(userService)

	routes.SetupAuthRoutes(r, userHandler)

	if err := r.Run(config.ServerPort); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
