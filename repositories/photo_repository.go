package repositories

import (
	"errors"
	"project2-golang/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	CreatePhoto(db *gorm.DB, Photo models.Photo) (models.Photo, error)
	GetPhotos(db *gorm.DB) ([]models.Photo, error)
	UpdatePhoto(db *gorm.DB, Photo models.Photo, PhotoId int) (models.Photo, error)
	DeletePhoto(db *gorm.DB, Photo models.Photo) error
	GetPhotoById(db *gorm.DB, PhotoId int) (models.Photo, error)
}

type PhotoRepositoryImpl struct {
}

func NewPhotoRepository() PhotoRepository {
	return &PhotoRepositoryImpl{}
}

func (r *PhotoRepositoryImpl) CreatePhoto(db *gorm.DB, Photo models.Photo) (models.Photo, error) {
	err := db.Create(&Photo).Error
	if err != nil {
		return Photo, errors.New(err.Error())
	}

	PhotoCreated := models.Photo{
		ID:        Photo.ID,
		Title:     Photo.Title,
		Caption:   Photo.Caption,
		PhotoURL:  Photo.PhotoURL,
		UserID:    Photo.UserID,
		CreatedAt: Photo.CreatedAt,
	}

	return PhotoCreated, nil
}

func (r *PhotoRepositoryImpl) GetPhotos(db *gorm.DB) ([]models.Photo, error) {
	photo := []models.Photo{}

	result := db.Table("photos").Scan(&photo)
	if result.RowsAffected == 0 {
		return photo, errors.New("Photo not found")
	}

	for i, p := range photo {
		user := models.User{}
		err := db.Table("users").Select([]string{"id", "email", "username"}).Where("id = ?", p.UserID).Scan(&user).Error
		if err != nil {
			continue
		}
		photo[i].User = &user
	}
	return photo, nil
}

func (r *PhotoRepositoryImpl) GetPhotoById(db *gorm.DB, PhotoId int) (models.Photo, error) {
	Photo := models.Photo{}

	result := db.Table("photos").Select([]string{"id", "user_id"}).Where("id = ?", PhotoId).Scan(&Photo)

	if result.RowsAffected == 0 {
		return Photo, errors.New("Photo not found")
	}

	return Photo, nil
}

func (r *PhotoRepositoryImpl) UpdatePhoto(db *gorm.DB, Photo models.Photo, PhotoId int) (models.Photo, error) {
	requestPhoto := Photo

	result := db.Where("id = ?", PhotoId).First(&Photo)

	if result.RowsAffected == 0 {
		return Photo, errors.New("Photo not found")
	}

	err := db.Model(&Photo).Where("id = ?", PhotoId).Updates(models.Photo{
		Title:    requestPhoto.Title,
		Caption:  requestPhoto.Caption,
		PhotoURL: requestPhoto.PhotoURL,
	}).Error

	if err != nil {
		return Photo, errors.New(err.Error())
	}

	PhotoUpdate := models.Photo{
		ID:        Photo.ID,
		Title:     Photo.Title,
		Caption:   Photo.Caption,
		PhotoURL:  Photo.PhotoURL,
		UserID:    Photo.UserID,
		UpdatedAt: Photo.UpdatedAt,
	}

	return PhotoUpdate, nil
}

func (r *PhotoRepositoryImpl) DeletePhoto(db *gorm.DB, Photo models.Photo) error {
	err := db.Delete(&Photo).Error
	if err != nil {
		return err
	}

	return nil
}
