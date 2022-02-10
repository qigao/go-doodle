package mysql

import (
	mysqlC "containers/db"
	models "forum/entities"
	sql "forum/repository/mysql"
	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	mg, err := migrate.NewWithDatabaseInstance("file://../../../sql", "mysql", driver)
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
