package repository

import (
	models "forum/entities"
)

type User interface {
	FindByID(u2 uint) (*models.User, error)
	FindByEmail(s string) (*models.User, error)
	FindByUsername(s string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	AddFollower(user *models.User, follower *models.User) error
	RemoveFollower(user *models.User, follower *models.User) error
	IsFollower(user, follower *models.User) (bool, error)
}
