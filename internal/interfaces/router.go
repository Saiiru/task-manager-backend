package interfaces

import (
	"task-manager-app/backend/internal/application"
	"task-manager-app/backend/internal/infrastructure"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Initialize handlers
	taskHandler := NewTaskHandler(application.NewTaskService(infrastructure.NewTaskRepository(db)))
	userHandler := NewUserHandler(application.NewUserService(infrastructure.NewUserRepository(db)))
	authHandler := NewAuthHandler(application.NewUserService(infrastructure.NewUserRepository(db)), []byte("your_jwt_secret"))

	// Task routes
	router.POST("/tasks", taskHandler.CreateTask)
	router.GET("/tasks", taskHandler.GetTasks)
	router.GET("/tasks/:id", taskHandler.GetTaskByID) // Adicionando rota GET /tasks/:id
	router.PUT("/tasks/:id", taskHandler.UpdateTask)
	router.DELETE("/tasks/:id", taskHandler.DeleteTask)

	// User routes
	router.POST("/users", userHandler.Register)
	router.POST("/login", authHandler.Login)
	router.GET("/users", userHandler.GetUsers)
	router.GET("/users/:id", userHandler.GetUserByID)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	return router
}
