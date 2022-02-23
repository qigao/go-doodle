package mysql

import (
	mysqlC "containers/db"
	"fmt"
	"forum/entity"
	repo "forum/repository/mysql"
	"forum/utils"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/volatiletech/null/v8"
)

var userRepo *repo.UserRepo
var articleRepo *repo.ArticleRepo

func TestMain(m *testing.M) {
	path := "../../../"
	config := utils.ReadDBConfigFromToml(path)
	container := mysqlC.NewMysqlContainer(config.User, config.Pass, config.DbName)
	container.CreateContainer()
	defer container.CloseContainer()
	host, port, _ := container.GetConnHostAndPort()
	dsn := utils.BuildDSNFromDbConfig(config, host, port)
	db := utils.SqlDbManager(dsn)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	mg, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s/db/sql", path), "mysql", driver)
	if err != nil {
		println(err.Error())
	}
	mg.Up()
	userRepo = repo.NewUserRepo(db)
	articleRepo = repo.NewArticleRepo(db)
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
		Slug:        "foo-slug",
	}
)
