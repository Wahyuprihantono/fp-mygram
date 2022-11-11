package routers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) *gin.Engine {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	router := gin.Default()

	UserRoutes(db, router)
	SocialMediaRoutes(db, router)
	PhotoRoutes(db, router)
	CommentRoutes(db, router)

	router.Use(gin.Recovery())

	return router
}
