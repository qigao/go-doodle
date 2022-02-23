package user

import (
	"fmt"
	"forum/model"
	"forum/repository"
	"forum/service"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type RequestLogin struct {
	User model.LoginUser `json:"user"`
	Repo repository.User
}

func (r *RequestLogin) Bind(c echo.Context) error {
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

func (r *RequestLogin) ValidateUser() error {
	userInfo, err := r.Repo.FindByEmail(r.User.Email)
	if err != nil {
		return fmt.Errorf("user not found")
	}
	if !service.CheckPassword(userInfo.Password, r.User.Password) {
		return fmt.Errorf("password is not correct")
	}
	return nil
}
