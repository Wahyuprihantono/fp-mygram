package middlewares

import (
	"project2-golang/helpers"
	"project2-golang/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommentAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		CommentRepository := repositories.NewCommentRepository()
		paramCommentId, err := strconv.Atoi(ctx.Param("CommentId"))
		if err != nil {
			helpers.FailedMessageResponse(ctx, "invalid parameter Comment id")
			return
		}

		tokenUserId := uint(ctx.MustGet("id").(float64))
		Comment, err := CommentRepository.GetCommentById(db, paramCommentId)
		if err != nil {
			helpers.FailedMessageResponse(ctx, err.Error())
			return
		}

		if tokenUserId != uint(Comment.UserID) {
			helpers.FailedMessageResponse(ctx, "unauthorized")
			return
		}

		ctx.Next()
	}
}
