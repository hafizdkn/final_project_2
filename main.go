package main

import (
	"github.com/gin-gonic/gin"

	"final_project_2/database"
	"final_project_2/handler"
	"final_project_2/helper"
	"final_project_2/middleware"
	"final_project_2/user"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}
	database.Migration(db)

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	authservice := helper.NewJwtService()

	userHandler := handler.NewUserHandler(userService, authservice)

	app := gin.Default()
	user := app.Group("/users")
	{
		user.POST("/register", userHandler.CreateUser)
		user.POST("/login", userHandler.UserLogin)
		user.GET("/users", userHandler.GetAllUsers)
		user.POST("/update", middleware.AuthMiddleware(authservice, userService), userHandler.UpdateUser)
		user.GET("/delete", middleware.AuthMiddleware(authservice, userService), userHandler.DeleteUser)
	}

	app.Run(":8080")
}
