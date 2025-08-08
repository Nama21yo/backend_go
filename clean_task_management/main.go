package cleantaskmanagement

import (
	"log"
	"os"
	"time"
	"clean_task_management/Delivery/controllers"
	"clean_task_management/Delivery/routers"
	"clean_task_management/Infrastructure"
	"clean_task_management/Repositories"
	"clean_task_management/Usecases"

	"github.com/joho/godotenv"
)

func main() {
	// load .env if present
	_ = godotenv.Load()

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret-change-me"
	}

	// initialize infrastructure
	jwtSvc := infrastructure.NewJWTService(jwtSecret)
	pwdSvc := infrastructure.NewBcryptPasswordService(0)

	// initialize repositories (in-memory for demo)
	userRepo := repositories.NewInMemoryUserRepository()
	taskRepo := repositories.NewInMemoryTaskRepository()

	// initialize usecases
	userUC := usecases.NewUserUsecase(userRepo, pwdSvc, jwtSvc, time.Hour*24)
	taskUC := usecases.NewTaskUsecase(taskRepo)

	// controllers
	userCtrl := controllers.NewUserController(userUC, jwtSvc)
	taskCtrl := controllers.NewTaskController(taskUC)

	// router
	r := routers.SetupRouter(userCtrl, taskCtrl, jwtSvc)
	log.Println("Starting server at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
