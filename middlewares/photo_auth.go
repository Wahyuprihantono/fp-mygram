package middlewares

import (
	"project2-golang/helpers"
	"project2-golang/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PhotoAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		PhotoRepository := repositories.NewPhotoRepository()

		paramPhotoId, err := strconv.Atoi(ctx.Param("PhotoId"))
		if err != nil {
			helpers.FailedMessageResponse(ctx, "invalid parameter Photo id")
			return
		}

		tokenUserId := uint(ctx.MustGet("id").(float64))
		Photo, err := PhotoRepository.GetPhotoById(db, paramPhotoId)
		if err != nil {
			helpers.FailedMessageResponse(ctx, err.Error())
			return
		}

		if tokenUserId != uint(Photo.UserID) {
			helpers.FailedMessageResponse(ctx, "unauthorized")
			return
		}

		ctx.Next()
	}
}
