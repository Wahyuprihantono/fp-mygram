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

func PhotoRoutes(db *gorm.DB, router *gin.Engine) {
	validate := validator.New()

	PhotoRepository := repositories.NewPhotoRepository()
	PhotoService := services.NewPhotoService(validate, PhotoRepository, db)
	PhotoController := controllers.NewPhotoController(PhotoService)

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", PhotoController.Create)
		photoRouter.GET("/", PhotoController.GetAllData)
		photoRouter.PUT("/:PhotoId", middlewares.PhotoAuthorization(db), PhotoController.Update)
		photoRouter.DELETE("/:PhotoId", middlewares.PhotoAuthorization(db), PhotoController.Delete)
	}
}
