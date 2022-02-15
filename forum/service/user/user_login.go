package user

import (
	"forum/model"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	User model.LoginUser `json:"user"`
}

func (r *LoginRequest) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}
