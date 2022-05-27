package mysql

import (
	"schema/entity"

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
	articleBar = &entity.Article{
		Title:       "foo Title",
		Description: null.NewString("foo Description", false),
		Body:        null.NewString("foo Body", false),
		Slug:        "foo-slug",
	}
	commentFoo = &entity.Comment{
		Body:      null.StringFrom("foo Body"),
		ArticleID: null.Uint64From(1),
		UserID:    null.Uint64From(1),
	}
	tagFoo = &entity.Tag{
		Tag: null.StringFrom("foo"),
	}
	tagBar = &entity.Tag{
		Tag: null.StringFrom("bar"),
	}
)
