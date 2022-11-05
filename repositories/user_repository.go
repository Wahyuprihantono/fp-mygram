package repositories

import (
	"project2-golang/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(db *gorm.DB, user models.User) (models.User, error)
	LoginUser(db *gorm.DB, user models.User) (models.User, error)
	UpdateUser(db *gorm.DB, user models.User, userId int) (models.User, error)
	DeleteUserByID(db *gorm.DB, user models.User) error
}
