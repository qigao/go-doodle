package user

import (
	"forum/service"
)

type Handler struct {
	Service service.IServiceUser
}

func NewUserHandler(us service.IServiceUser) *Handler {
	return &Handler{
		Service: us,
	}
}
