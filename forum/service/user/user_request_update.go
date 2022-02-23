package user

import (
	"forum/entity"
	"forum/model"
	"forum/repository"
	"forum/service"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type UpdateRequest struct {
	User model.UpdateUser `json:"user"`
	Repo repository.User
}

func (r *UpdateRequest) Bind(c echo.Context) error {
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
func (r *UpdateRequest) populate(u *entity.User) error {
	if r.User.Password != "" {
		h, err := service.HashPassword(r.User.Password)
		if err != nil {
			log.Error().Err(err).Msg("Hash password error")
			return err
		}
		u.Password = h
	}
	if u.Bio.Valid {
		r.User.Bio = u.Bio.String
	}
	if u.Image.Valid {
		r.User.Image = u.Image.String
	}
	return nil
}

func (r *UpdateRequest) FindThenUpdateUser(uid uint) error {
	u, err := r.Repo.FindByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindByID error")
		return err
	}
	if err = r.populate(u); err != nil {
		log.Error().Err(err).Msg("error changing user info")
		return err
	}
	return r.Repo.UpdateUser(u)
}
