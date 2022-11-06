package repositories

import (
	"errors"
	"project2-golang/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(db *gorm.DB, Comment models.Comment) (models.Comment, error)
	GetComments(db *gorm.DB) ([]models.Comment, error)
	UpdateComment(db *gorm.DB, Comment models.Comment, CommentId int) (models.Comment, error)
	DeleteComment(db *gorm.DB, Comment models.Comment) error
	GetCommentById(db *gorm.DB, CommentId int) (models.Comment, error)
}

type CommentRepositoryImpl struct {
}

func NewCommentRepository() CommentRepository {
	return &CommentRepositoryImpl{}
}

func (r *CommentRepositoryImpl) CreateComment(db *gorm.DB, Comment models.Comment) (models.Comment, error) {
	err := db.Create(&Comment).Error
	if err != nil {
		return Comment, errors.New(err.Error())
	}

	CommentCreated := models.Comment{
		ID:        Comment.ID,
		Message:   Comment.Message,
		PhotoID:   Comment.PhotoID,
		UserID:    Comment.UserID,
		CreatedAt: Comment.CreatedAt,
	}

	return CommentCreated, nil
}

func (r *CommentRepositoryImpl) GetComments(db *gorm.DB) ([]models.Comment, error) {
	comment := []models.Comment{}

	result := db.Table("comments").Scan(&comment)
	if result.RowsAffected == 0 {
		return comment, errors.New("comment not found")
	}

	for i, p := range comment {
		user := models.User{}
		err := db.Table("users").Select([]string{"id", "email", "username"}).Where("id = ?", p.UserID).Scan(&user).Error
		if err != nil {
			continue
		}
		comment[i].User = &user
	}

	for i, p := range comment {
		photo := models.Photo{}
		err := db.Table("photos").Select([]string{"id", "title", "caption", "photo_url", "user_id"}).Where("id = ?", p.PhotoID).Scan(&photo).Error
		if err != nil {
			continue
		}
		comment[i].Photo = &photo
	}
	return comment, nil
}

func (r *CommentRepositoryImpl) GetCommentById(db *gorm.DB, CommentId int) (models.Comment, error) {
	Comment := models.Comment{}

	result := db.Table("comments").Select([]string{"id", "user_id"}).Where("id = ?", CommentId).Scan(&Comment)

	if result.RowsAffected == 0 {
		return Comment, errors.New("Comment not found")
	}

	return Comment, nil
}

func (r *CommentRepositoryImpl) UpdateComment(db *gorm.DB, Comment models.Comment, CommentId int) (models.Comment, error) {
	requestComment := Comment

	result := db.Where("id = ?", CommentId).First(&Comment)

	if result.RowsAffected == 0 {
		return Comment, errors.New("Photo not found")
	}

	err := db.Model(&Comment).Where("id = ?", CommentId).Updates(models.Comment{
		Message: requestComment.Message,
	}).Error

	if err != nil {
		return Comment, errors.New(err.Error())
	}

	CommentUpdate := models.Comment{
		ID:        Comment.ID,
		Message:   Comment.Message,
		UserID:    Comment.UserID,
		PhotoID:   Comment.PhotoID,
		UpdatedAt: Comment.UpdatedAt,
	}

	return CommentUpdate, nil
}

func (r *CommentRepositoryImpl) DeleteComment(db *gorm.DB, Comment models.Comment) error {
	err := db.Delete(&Comment).Error
	if err != nil {
		return err
	}

	return nil
}
