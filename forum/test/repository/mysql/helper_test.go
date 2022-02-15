package mysql

import (
	mysqlC "containers/db"
	"forum/entity"
	sql "forum/repository/mysql"
	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/volatiletech/null/v8"
	"os"
	"testing"
)

var userRepo *sql.UserRepo
var articleRepo *sql.ArticleRepo

func TestMain(m *testing.M) {
	container := mysqlC.NewMysqlContainer()
	container.CreateContainer()
	defer container.CloseContainer()
	db := container.SqlDbManager()
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	mg, err := migrate.NewWithDatabaseInstance("file://../../../db/sql", "mysql", driver)
	if err != nil {
		println(err.Error())
	}
	mg.Up()
	userRepo = sql.NewUserRepo(db)
	articleRepo = sql.NewArticleRepo(db)
	code := m.Run()
	os.Exit(code)
}

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
		Slug:        "foo-title",
	}
)
