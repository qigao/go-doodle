package user

import (
	"forum/repository"
)

type Handler struct {
	userRepo repository.User
}

func NewUserHandler(us repository.User) *Handler {
	return &Handler{
		userRepo: us,
	}
}
