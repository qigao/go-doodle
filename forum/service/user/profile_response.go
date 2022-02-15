package user

import (
	"forum/entity"
	"forum/model"
	"forum/repository"
)

type profileResponse struct {
	Profile model.ProfileType `json:"profile"`
}

func NewProfileResponse(us repository.User, userID uint, u *entity.User) *profileResponse {
	r := new(profileResponse)
	r.Profile.Username = u.Username
	if u.Bio.Valid {
		r.Profile.Bio = &(u.Bio.String)
	}
	if u.Image.Valid {
		r.Profile.Image = &(u.Image.String)
	}
	//r.Profile.Following, _ = us.IsFollower(u.ID, userID)
	return r
}
