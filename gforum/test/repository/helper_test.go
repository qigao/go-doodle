package repository

import (
	"gforum/model"
	"gforum/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	testDb "test_container/db"
	"testing"
)

var userRepo *repository.UserRepository
var articleRepo *repository.ArticleRepository

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Article{},
		&model.Tag{},
		&model.Comment{},
	)

}

func TestMain(m *testing.M) {
	container := testDb.NewMysqlContainer()
	container.CreateContainer()
	db := container.DbManager()

	userRepo = repository.NewUserRepository(db)
	articleRepo = repository.NewArticleRepository(db)
	err := AutoMigrate(db)
	if err != nil {
		return
	}
	code := m.Run()
	container.CloseContainer()
	os.Exit(code)
}

func mockMysql() (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err) // Error here
	}
	return gormDB, mock
}

var (
	userFoo = &model.User{
		Username: "foo",
		Email:    "foo@foo.com",
		Password: "123456",
	}
	userBar = &model.User{
		Username: "bar",
		Email:    "bar@bar.com",
		Password: "123456",
	}
	articleFoo = &model.Article{
		Title:          "foo Title",
		Description:    "foo Description",
		Body:           "foo Body",
		Author:         *userFoo,
		UserID:         1,
		FavoritesCount: 1,
		FavoritedUsers: []model.User{*userBar},
		Comments: []model.Comment{
			{
				Body:      "foo comment",
				UserID:    2,
				Author:    *userBar,
				ArticleID: 1,
			},
		},
		Tags: []model.Tag{
			{
				Name: "foo",
			},
			{
				Name: "bar",
			},
		},
	}
)
