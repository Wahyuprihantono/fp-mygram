package services

import (
	"project2-golang/models"
	"project2-golang/params"
)

type UserService interface {
	CreateUser(userParams params.CreateUser) (models.User, error)
	LoginUser(userParams params.LoginUser) (string, error)
	UpdateUser(userParams params.UpdateUser, userId int) (models.User, error)
	DeleteUserByID(userId int) error
}
