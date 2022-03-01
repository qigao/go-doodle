// DONT EDIT: Auto generated

package service

import (
	"forum/entity"
	"forum/model"
)

// ServiceUser ...
type ServiceUser interface {
	CheckUser(user *model.LoginUser) error
	CreateUser(user *model.RegisterUser) error
	FllowUserByUserName(uid uint, userName string) error
	GetUserByID(uid uint) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUserByUserName(username string) (*entity.User, error)
	UnFollowUserByUserName(uid uint, userName string) error
	GetFollowersByUserID(uid uint) ([]*entity.User, error)
	GetFollowingUser(uid uint) ([]*entity.User, error)
}
