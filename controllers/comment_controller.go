package controllers

import (
	"net/http"
	"project2-golang/helpers"
	"project2-golang/params"
	"project2-golang/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentsController interface {
	Create(ctx *gin.Context)
	GetAllData(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type CommentController struct {
	CommentService services.CommentService
}

func NewCommentController(service services.CommentService) CommentsController {
	return &CommentController{
		CommentService: service,
	}
}

func (c *CommentController) Create(ctx *gin.Context) {
	request := params.CreateComments{}
	requestValid := helpers.ReadFromRequestBody(ctx, &request)
	if !requestValid {
		return
	}

	userId := ctx.MustGet("id").(float64)
	request.UserID = uint(userId)

	Comment, err := c.CommentService.Create(request)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, Comment)
}

func (c *CommentController) GetAllData(ctx *gin.Context) {
	Comment, err := c.CommentService.GetAllData()
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Comment": Comment,
	})
}

func (c *CommentController) Update(ctx *gin.Context) {
	request := params.CreateComments{}
	requestValid := helpers.ReadFromRequestBody(ctx, &request)
	if !requestValid {
		return
	}

	CommentId, err := strconv.Atoi(ctx.Param("CommentId"))
	if err != nil {
		helpers.FailedMessageResponse(ctx, "invalid parameter Comment id")
		return
	}

	Comment, err := c.CommentService.Update(request, CommentId)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, Comment)
}

func (c *CommentController) Delete(ctx *gin.Context) {
	CommentId, err := strconv.Atoi(ctx.Param("CommentId"))
	if err != nil {
		helpers.FailedMessageResponse(ctx, "invalid parameter Photo id")
		return
	}

	err = c.CommentService.Delete(CommentId)

	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	helpers.SuccessMessageResponse(ctx, "Your Comment has been successfully deleted")
}
