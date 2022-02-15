package user

import (
	"forum/model"
	"forum/repository"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type LoginRequest struct {
	User model.LoginUser `json:"user"`
	Repo repository.User
}

func (r *LoginRequest) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		log.Error().Err(err).Msg("Bind error")
		return err
	}
	if err := c.Validate(r); err != nil {
		log.Error().Err(err).Msg("Validate error")
		return err
	}
	return nil
}
