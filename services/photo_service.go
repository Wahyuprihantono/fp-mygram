package services

import (
	"errors"
	"project2-golang/models"
	"project2-golang/params"
	"project2-golang/repositories"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type PhotoService interface {
	Create(PhotoParams params.CreatePhotos) (models.Photo, error)
	GetAllData() ([]models.Photo, error)
	Update(PhotoParam params.CreatePhotos, PhotoId int) (models.Photo, error)
	Delete(PhotoId int) error
}

type PhotoServiceImpl struct {
	PhotoRepository repositories.PhotoRepository
	DB              *gorm.DB
	Validate        *validator.Validate
}

func NewPhotoService(validate *validator.Validate, repository repositories.PhotoRepository, db *gorm.DB) PhotoService {
	return &PhotoServiceImpl{
		Validate:        validate,
		PhotoRepository: repository,
		DB:              db,
	}
}

func (s *PhotoServiceImpl) Create(PhotoParams params.CreatePhotos) (models.Photo, error) {
	Photo := models.Photo{}

	errValidate := s.Validate.Struct(PhotoParams)
	if errValidate != nil {
		return Photo, errors.New(errValidate.Error())
	}

	Photo.Title = PhotoParams.Title
	Photo.Caption = PhotoParams.Caption
	Photo.PhotoURL = PhotoParams.PhotoURL
	Photo.UserID = PhotoParams.UserID

	response, err := s.PhotoRepository.CreatePhoto(s.DB, Photo)

	if err != nil {
		return Photo, errors.New(err.Error())
	}

	return response, nil
}

func (s *PhotoServiceImpl) GetAllData() ([]models.Photo, error) {
	Photo := []models.Photo{}

	response, err := s.PhotoRepository.GetPhotos(s.DB)

	if err != nil {
		return Photo, errors.New(err.Error())
	}

	return response, nil
}

func (s *PhotoServiceImpl) Update(PhotoParams params.CreatePhotos, PhotoId int) (models.Photo, error) {
	Photo := models.Photo{}

	errRequest := s.Validate.Struct(PhotoParams)
	if errRequest != nil {
		return Photo, errors.New(errRequest.Error())
	}

	Photo.Title = PhotoParams.Title
	Photo.Caption = PhotoParams.Caption
	Photo.PhotoURL = PhotoParams.PhotoURL
	Photo.UserID = PhotoParams.UserID

	response, err := s.PhotoRepository.UpdatePhoto(s.DB, Photo, PhotoId)

	if err != nil {
		return Photo, errors.New(err.Error())
	}

	return response, nil
}

func (s *PhotoServiceImpl) Delete(PhotoId int) error {
	Photo := models.Photo{
		ID: uint(PhotoId),
	}

	err := s.PhotoRepository.DeletePhoto(s.DB, Photo)
	if err != nil {
		return err
	}

	return nil
}
