// DON'T EDIT: This is generated code

package repository

import (
	"schema/entity"
)

// User ...
type User interface {
	FindUserByID(uid uint) (*entity.User, error)
	FindByEmail(s string) (*entity.User, error)
	FindUserByUserName(s string) (*entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(user *entity.User) error
	AddFollower(user *entity.User, follower *entity.User) error
	RemoveFollower(user *entity.User, follower *entity.User) error
	IsFollower(user, follower *entity.User) (bool, error)
	GetFollowers(user *entity.User) ([]*entity.User, error)
	GetFollowingUsers(user *entity.User) ([]*entity.User, error)
}
