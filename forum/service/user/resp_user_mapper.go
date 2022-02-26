package user

import (
	"forum/entity"
	"forum/utils"
)

func NewProfileResponse(u *entity.User) *profileResponse {
	r := new(profileResponse)
	r.Profile.Username = u.Username
	if u.Bio.Valid {
		r.Profile.Bio = &(u.Bio.String)
	}
	if u.Image.Valid {
		r.Profile.Image = &(u.Image.String)
	}
	return r
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
