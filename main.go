package main

import (
	"log"

	"github.com/adhityaf/deall-be-test/config"
	"github.com/adhityaf/deall-be-test/controllers"
	"github.com/adhityaf/deall-be-test/middlewares"
	"github.com/adhityaf/deall-be-test/repositories"
	"github.com/adhityaf/deall-be-test/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Load Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.ConnectDB()
	route := gin.Default()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	mainRouter := route.Group("/v1")
	{
		mainRouter.POST("/login", userController.Login)
		mainRouter.POST("/register", userController.Register)

		authorized := mainRouter.Group("/")
		authorized.Use(middlewares.Auth())
		{
			authorized.GET("/user", userController.GetUserProfile)

			admin := authorized.Group("/")
			admin.Use(middlewares.IsAdmin())
			{
				admin.POST("/user", userController.CreateUser)
				admin.GET("/users", userController.GetAllUsers)
				admin.GET("/user/:userId", userController.GetUserById)
				admin.PUT("/user/:userId", userController.UpdateUserById)
				admin.DELETE("/user/:userId", userController.DeleteUserById)
			}
		}
	}

	route.Run()
}
