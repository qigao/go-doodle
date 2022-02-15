package user

import (
	"forum/entity"
	"forum/model"
	"forum/service"
	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	User model.RegisterUser `json:"user"`
}

func (r *RegisterRequest) Bind(c echo.Context, u *entity.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Username = r.User.Username
	u.Email = r.User.Email
	h, err := service.HashPassword(r.User.Password)
	if err != nil {
		return err
	}
	u.Password = h
	return nil
}
