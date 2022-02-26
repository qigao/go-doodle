package user

import (
	"forum/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProfileResponse(t *testing.T) {
	type args struct {
		u *entity.User
	}
	tests := []struct {
		name string
		args args
		want *profileResponse
	}{
		{"TestNewProfileResponse", args{userFoo}, &profileResponse{Profile: *profileTypeFoo}},
		{"TestNewProfileResponseWithNull", args{userWithoutBio}, &profileResponse{Profile: *profileTypeFoo}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewProfileResponse(tt.args.u)
			if got.Profile.Bio != nil {
				assert.Equal(t, tt.args.u.Bio.String, *got.Profile.Bio)
				assert.Equal(t, tt.args.u.Image.String, *got.Profile.Image)
			} else {
				assert.Nil(t, got.Profile.Bio)
				assert.Nil(t, got.Profile.Image)
			}
		})
	}
}

func TestNewUserResponse(t *testing.T) {
	type args struct {
		u *entity.User
	}
	tests := []struct {
		name string
		args args
		want *userResponse
	}{
		{"TestNewUserResponse", args{userFoo}, &userResponse{User: *userFooResponse}},
		{"TestNewUserResponseWithNull", args{userWithoutBio}, &userResponse{User: *userFooResponseWithNull}},
	}
	for _, tt := range tests {
		got :=NewUserResponse(tt.args.u)
		t.Run(tt.name, func(t *testing.T) {
			if got.User.Bio != nil {
				assert.Equal(t, tt.args.u.Bio.String, *got.User.Bio)
				assert.Equal(t, tt.args.u.Image.String, *got.User.Image)
			}else {
				assert.Nil(t, got.User.Bio)
				assert.Nil(t, got.User.Image)
			}
		})
	}
}
