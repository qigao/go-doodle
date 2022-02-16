package user

import (
	"forum/entity"
	"forum/model"
	"forum/service"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
)

type UpdateRequest struct {
	User model.UpdateUser `json:"user"`
}

func (r *UpdateRequest) Populate(u *entity.User) {
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.Password = u.Password
	if u.Bio.Valid {
		r.User.Bio = u.Bio.String
	}
	if u.Image.Valid {
		r.User.Image = u.Image.String
	}
}

func (r *UpdateRequest) Bind(c echo.Context, u *entity.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Username = r.User.Username
	u.Email = r.User.Email
	if r.User.Password != u.Password {
		h, err := service.HashPassword(r.User.Password)
		if err != nil {
			return err
		}
		u.Password = h
	}
	u.Bio = null.StringFrom(r.User.Bio)
	u.Image = null.StringFrom(r.User.Image)

	return nil
}
