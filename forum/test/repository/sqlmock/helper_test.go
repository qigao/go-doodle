package sqlmock

import (
	models "forum/entities"
	"forum/repository/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var userRepo *mysql.UserRepo
var (
	userFoo = &models.User{
		Username: "foo",
		Email:    "foo@foo.com",
		Password: "123456",
	}
	userBar = &models.User{
		Username: "bar",
		Email:    "bar@bar.com",
		Password: "123456",
	}
	//articleFoo = &models.Article{
	//	Title:       "foo Title",
	//	Description: "foo Description",
	//	Body:        "foo Body",
	//	Slug:        "foo-title",
	//}
)
