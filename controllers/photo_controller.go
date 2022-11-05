package controllers

import (
	"net/http"
	"project2-golang/helpers"
	"project2-golang/params"
	"project2-golang/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotosController interface {
	Create(ctx *gin.Context)
	GetAllData(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type PhotoController struct {
	PhotoService services.PhotoService
}

func NewPhotoController(service services.PhotoService) PhotosController {
	return &PhotoController{
		PhotoService: service,
	}
}

func (c *PhotoController) Create(ctx *gin.Context) {
	request := params.CreatePhotos{}
	requestValid := helpers.ReadFromRequestBody(ctx, &request)
	if !requestValid {
		return
	}

	userId := ctx.MustGet("id").(float64)
	request.UserID = uint(userId)

	Photo, err := c.PhotoService.Create(request)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, Photo)
}

func (c *PhotoController) GetAllData(ctx *gin.Context) {
	Photo, err := c.PhotoService.GetAllData()
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Photo": Photo,
	})
}

func (c *PhotoController) Update(ctx *gin.Context) {
	request := params.CreatePhotos{}
	requestValid := helpers.ReadFromRequestBody(ctx, &request)
	if !requestValid {
		return
	}

	PhotoId, err := strconv.Atoi(ctx.Param("PhotoId"))
	if err != nil {
		helpers.FailedMessageResponse(ctx, "invalid parameter Photo id")
		return
	}

	Photo, err := c.PhotoService.Update(request, PhotoId)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, Photo)
}

func (c *PhotoController) Delete(ctx *gin.Context) {
	PhotoId, err := strconv.Atoi(ctx.Param("PhotoId"))
	if err != nil {
		helpers.FailedMessageResponse(ctx, "invalid parameter Photo id")
		return
	}

	err = c.PhotoService.Delete(PhotoId)

	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	helpers.SuccessMessageResponse(ctx, "Your Photo has been successfully deleted")
}
