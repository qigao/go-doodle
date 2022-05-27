package user

import (
	"forum/service"
)

type Handler struct {
	Service service.UserService
}

func NewUserHandler(us service.UserService) *Handler {
	return &Handler{
		Service: us,
	}
}
