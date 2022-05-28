package user

import (
	"schema/entity"

	"forum/model"

	"github.com/volatiletech/null/v8"
)

var (
	userFoo = &entity.User{
		Username: "foo",
		Email:    "foo@foo.com",
		Password: "123456",
		Bio:      null.StringFrom("foo bio"),
		Image:    null.StringFrom("foo image"),
	}
	userWithoutBio = &entity.User{
		Username: "foo",
		Email:    "foo@foo.com",
		Password: "123456",
		Bio:      null.StringFromPtr(nil),
		Image:    null.StringFromPtr(nil),
	}
	userBio        = "foo bio"
	userImage      = "http://foo.com/foo.png"
	profileTypeFoo = &model.ProfileType{
		Username: "foo",
		Email:    "foo@foo.com",
		Bio:      &userBio,
		Image:    &userImage,
	}

	userFooResponse = &model.Response{
		Username: "foo",
		Email:    "foo@foo.com",
		Bio:      &userBio,
		Image:    &userImage,
	}
	userFooResponseWithNull = &model.Response{
		Username: "foo",
		Email:    "foo@foo.com",
		Bio:      nil,
		Image:    nil,
	}
)
