package services

import (
	"errors"
	"project2-golang/models"
	"project2-golang/params"
	"project2-golang/repositories"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CommentService interface {
	Create(PhotoParams params.CreateComments) (models.Comment, error)
	GetAllData() ([]models.Comment, error)
	Update(CommentParam params.CreateComments, CommentId int) (models.Comment, error)
	Delete(CommentId int) error
}

type CommentServiceImpl struct {
	CommentRepository repositories.CommentRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

func NewCommentService(validate *validator.Validate, repository repositories.CommentRepository, db *gorm.DB) CommentService {
	return &CommentServiceImpl{
		Validate:          validate,
		CommentRepository: repository,
		DB:                db,
	}
}

func (s *CommentServiceImpl) Create(CommentParams params.CreateComments) (models.Comment, error) {
	Comment := models.Comment{}

	errValidate := s.Validate.Struct(CommentParams)
	if errValidate != nil {
		return Comment, errors.New(errValidate.Error())
	}

	Comment.Message = CommentParams.Message
	Comment.PhotoID = CommentParams.PhotoID
	Comment.UserID = CommentParams.UserID

	response, err := s.CommentRepository.CreateComment(s.DB, Comment)

	if err != nil {
		return Comment, errors.New(err.Error())
	}

	return response, nil
}

func (s *CommentServiceImpl) GetAllData() ([]models.Comment, error) {
	Comment := []models.Comment{}

	response, err := s.CommentRepository.GetComments(s.DB)

	if err != nil {
		return Comment, errors.New(err.Error())
	}

	return response, nil
}

func (s *CommentServiceImpl) Update(CommentParams params.CreateComments, CommentId int) (models.Comment, error) {
	Comment := models.Comment{}

	errRequest := s.Validate.Struct(CommentParams)
	if errRequest != nil {
		return Comment, errors.New(errRequest.Error())
	}

	Comment.Message = CommentParams.Message
	Comment.PhotoID = CommentParams.PhotoID
	Comment.UserID = CommentParams.UserID

	response, err := s.CommentRepository.UpdateComment(s.DB, Comment, CommentId)

	if err != nil {
		return Comment, errors.New(err.Error())
	}

	return response, nil
}

func (s *CommentServiceImpl) Delete(CommentId int) error {
	Comment := models.Comment{
		ID: uint(CommentId),
	}

	err := s.CommentRepository.DeleteComment(s.DB, Comment)
	if err != nil {
		return err
	}

	return nil
}
