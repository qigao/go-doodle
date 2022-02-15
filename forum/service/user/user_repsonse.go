package user

import (
	"forum/entity"
	"forum/model"
	"forum/utils"
)

type userResponse struct {
	User model.Response `json:"user"`
}

func NewUserResponse(u *entity.User) *userResponse {
	r := new(userResponse)
	r.User.Username = u.Username
	r.User.Email = u.Email
	if u.Bio.Valid {
		r.User.Bio = &(u.Bio.String)
	}
	if u.Image.Valid {
		r.User.Image = &(u.Image.String)
	}
	r.User.Token = utils.GenerateJWT(uint(u.ID))
	return r
}
