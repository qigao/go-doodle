package mysql

import (
	"forum/entity"

	"github.com/volatiletech/null/v8"
)

var (
	userFoo = &entity.User{
		Username: "foo",
		Email:    "foo@foo.com",
		Password: "123456",
	}
	userBar = &entity.User{
		Username: "bar",
		Email:    "bar@bar.com",
		Password: "123456",
	}
	articleFoo = &entity.Article{
		Title:       "foo Title",
		Description: null.NewString("foo Description", false),
		Body:        null.NewString("foo Body", false),
		Slug:        "foo-slug",
	}
)
