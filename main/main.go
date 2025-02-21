package main

import (
	"log"

	"asset_management/config"
	"asset_management/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := config.ConnectRedis(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	r := gin.Default()

	userController := controllers.NewUserController()
	resourceController := controllers.NewResourceController()

	r.POST("/users", userController.CreateUser)
	r.GET("/users", userController.GetUsers)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	r.GET("/resources", resourceController.GetResources)
	r.POST("/resources", resourceController.AddResource)
	r.PUT("/resources/:id", resourceController.UpdateResource)
	r.DELETE("/resources/:id", resourceController.DeleteResource)

	r.Run(":8080")
}
