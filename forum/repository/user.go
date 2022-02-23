package repository

import (
	entity "forum/entity"
)

type User interface {
	FindByID(u uint) (*entity.User, error)
	FindByEmail(s string) (*entity.User, error)
	FindByUsername(s string) (*entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(user *entity.User) error
	AddFollower(user *entity.User, follower *entity.User) error
	RemoveFollower(user *entity.User, follower *entity.User) error
	IsFollower(user, follower *entity.User) (bool, error)
	GetFollowers(user *entity.User) ([]*entity.User, error)
	GetFollowingUsers(user *entity.User) ([]*entity.User, error)
}
