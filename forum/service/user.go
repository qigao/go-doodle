// DON'T EDIT: This is generated code

package service

import (
	"forum/model"
	"schema/entity"
)

// UserService ...
type UserService interface {
	CheckUser(user *model.LoginUser) error
	CreateUser(user *model.RegisterUser) error
	FollowUserByUserName(uid uint, userName string) error
	GetUserByID(uid uint) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUserByUserName(username string) (*entity.User, error)
	UnFollowUserByUserName(uid uint, userName string) error
	GetFollowersByUserID(uid uint) ([]*entity.User, error)
	GetFollowingUser(uid uint) ([]*entity.User, error)
	UpdateUser(user *entity.User) error
}
