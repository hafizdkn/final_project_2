package main

import (
	"github.com/gin-gonic/gin"

	"final_project_2/auth"
	"final_project_2/comment"
	"final_project_2/database"
	"final_project_2/handler"
	"final_project_2/middleware"
	"final_project_2/photo"
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
	authservice := auth.NewJwtService()

	photoRepository := photo.NewPhotoRepository(db)
	photoService := photo.NewServiceRepository(photoRepository)

	commentRepository := comment.NewCommentRepository(db)
	commentService := comment.NewServiceRepository(commentRepository)

	userHandler := handler.NewUserHandler(userService, authservice)
	photoHandler := handler.NewPhotoHandler(photoService)
	commentHandler := handler.NewCommentHandler(commentService)

	app := gin.Default()
	user := app.Group("/users")
	{
		user.POST("/register", userHandler.CreateUser)
		user.POST("/login", userHandler.UserLogin)
		user.GET("/users", userHandler.GetAllUsers)
		user.POST("/update", middleware.AuthMiddleware(authservice, userService), userHandler.UpdateUser)
		user.DELETE("/delete", middleware.AuthMiddleware(authservice, userService), userHandler.DeleteUser)

		user.POST("/photos", middleware.AuthMiddleware(authservice, userService), photoHandler.CreatePhoto)
		user.GET("/photos", photoHandler.GetPhotos)
		user.PUT("/photos/:photoId", middleware.AuthMiddleware(authservice, userService), photoHandler.UpdatePhoto)
		user.DELETE("/photos/:photoId", middleware.AuthMiddleware(authservice, userService), photoHandler.DeletePhoto)

		user.POST("/comments", middleware.AuthMiddleware(authservice, userService), commentHandler.CreateComment)
		user.GET("/comments", middleware.AuthMiddleware(authservice, userService), commentHandler.GetComments)
		user.PUT("/comments/:commentId", middleware.AuthMiddleware(authservice, userService), commentHandler.UpdateComment)
		user.DELETE("/comments/:commentId", middleware.AuthMiddleware(authservice, userService), commentHandler.DeleteComment)
	}

	app.Run(":8080")
}
