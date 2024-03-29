package routers

import (
	"project2-golang/controllers"
	"project2-golang/middlewares"
	"project2-golang/repositories"
	"project2-golang/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB, router *gin.Engine) {
	validate := validator.New()

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository, db, validate)
	userController := controllers.NewUserController(userService)

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.CreateUser)
		userRouter.POST("/login", userController.LoginUser)
		userRouter.Use(middlewares.Authentication())
		userRouter.PUT("/:userId", middlewares.UserAuthorization(), userController.UpdateUser)
		userRouter.DELETE("/:userId", middlewares.UserAuthorization(), userController.DeleteUser)
	}
}
