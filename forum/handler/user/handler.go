package user

import (
	"forum/repository"
)

type Handler struct {
	user repository.User
}

func NewUserHandler(us repository.User) *Handler {
	return &Handler{
		user: us,
	}
}
