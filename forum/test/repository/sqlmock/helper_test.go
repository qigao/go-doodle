package sqlmock

import (
	"forum/entity"
	"forum/repository/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var userRepo *mysql.UserRepo
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
	//articleFoo = &entity.Article{
	//	Title:       "foo Title",
	//	Description: "foo Description",
	//	Body:        "foo Body",
	//	Slug:        "foo-title",
	//}
)
