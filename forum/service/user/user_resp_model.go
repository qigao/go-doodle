package user

import "forum/model"

type profileResponse struct {
	Profile model.ProfileType `json:"profile"`
}
type userResponse struct {
	User model.Response `json:"user"`
}
