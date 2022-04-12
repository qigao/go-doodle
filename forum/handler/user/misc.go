package user

import (
	"forum/service"
)

type Handler struct {
	Service service.ServiceUser
}

func NewUserHandler(us service.ServiceUser) *Handler {
	return &Handler{
		Service: us,
	}
}
