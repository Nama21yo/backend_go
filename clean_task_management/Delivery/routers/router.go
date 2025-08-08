package routers

import (
	"task-manager/Delivery/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userCtrl *controllers.UserController, taskCtrl *controllers.TaskController, jwtSvc infrastructure.JWTService) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/register", userCtrl.Register)
		api.POST("/login", userCtrl.Login)
		protected := api.Group("/")
		protected.Use(infrastructure.AuthMiddleware(jwtSvc))
		{
			protected.POST("/tasks", taskCtrl.CreateTask)
			protected.PUT("/tasks", taskCtrl.UpdateTask)
			protected.GET("/tasks", taskCtrl.ListTasks)
			protected.GET("/tasks/:id", taskCtrl.GetTask)
			protected.DELETE("/tasks/:id", taskCtrl.DeleteTask)
		}
	}
	return r
}
