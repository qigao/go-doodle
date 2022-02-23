package user

import (
	"forum/entity"
	"forum/model"
	"forum/repository"
	"forum/service"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type RegisterRequest struct {
	User model.RegisterUser `json:"user"`
	Repo repository.User
}

func (r *RegisterRequest) Bind(c echo.Context) error {
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

func (r *RegisterRequest) CreateUser() error {
	var u *entity.User
	u.Username = r.User.Username
	u.Email = r.User.Email
	h, err := service.HashPassword(r.User.Password)
	if err != nil {
		log.Error().Err(err).Msg("Hash password error")
		return err
	}
	u.Password = h
	return r.Repo.CreateUser(u)
}
