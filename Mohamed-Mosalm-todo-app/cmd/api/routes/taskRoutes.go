package routes

import (
	"github.com/MohamedMosalm/Todo-App/cmd/api/handlers"
	"github.com/MohamedMosalm/Todo-App/utils/middleware"
	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.Engine, taskHandler *handlers.TaskHandler, jwtSecret string) {
	taskRoutes := router.Group("/api/tasks")
	taskRoutes.Use(middleware.AuthMiddleware(jwtSecret))
	{
		taskRoutes.POST("", taskHandler.CreateTask)
		taskRoutes.GET("", taskHandler.GetTasks)
		taskRoutes.PUT("/:id", taskHandler.UpdateTask)
		taskRoutes.DELETE("/:id", taskHandler.DeleteTask)
	}
}
