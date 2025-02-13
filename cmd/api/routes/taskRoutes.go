package routes

import (
	"github.com/MohamedMosalm/To-Do-List/cmd/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.Engine, taskHandler *handlers.TaskHandler) {
	taskRoutes := router.Group("/api/tasks")
	{
		taskRoutes.POST("", taskHandler.CreateTask)
		taskRoutes.GET("", taskHandler.GetTasks)
		taskRoutes.PUT("/:id", taskHandler.UpdateTask)
		taskRoutes.DELETE("/:id", taskHandler.DeleteTask)
	}
}
