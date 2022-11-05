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

func CommentRoutes(db *gorm.DB, router *gin.Engine) {
	validate := validator.New()

	CommentRepository := repositories.NewCommentRepository()
	CommentService := services.NewCommentService(validate, CommentRepository, db)
	CommentController := controllers.NewCommentController(CommentService)

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", CommentController.Create)
		commentRouter.GET("/", CommentController.GetAllData)
		commentRouter.PUT("/:CommentId", middlewares.CommentAuthorization(db), CommentController.Update)
		commentRouter.DELETE("/:CommentId", middlewares.CommentAuthorization(db), CommentController.Delete)
	}
}
