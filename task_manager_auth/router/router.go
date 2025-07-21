package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/task_manager/controllers"
	"github.com/yourusername/task_manager/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	task := r.Group("/tasks")
	task.Use(middleware.AuthMiddleware())
	{
		task.GET("/", controllers.GetTasks)
		task.GET("/:id", controllers.GetTaskByID)
		task.POST("/", middleware.AdminMiddleware(), controllers.CreateTask)
		task.PUT("/:id", middleware.AdminMiddleware(), controllers.UpdateTask)
		task.DELETE("/:id", middleware.AdminMiddleware(), controllers.DeleteTask)
	}

	r.PUT("/promote/:username", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.PromoteUser)

	return r
}
