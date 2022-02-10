package sqlmock

import (
	"fmt"
	"forum/repository/mysql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestUserRepo_CreateUserByMock(t *testing.T) {
	t.Run("transaction start failed when create user", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			fmt.Println("error creating mock database")
			return
		}
		defer db.Close()
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := mysql.UserRepo{Db: db}
		err = repo.CreateUser(userFoo)
		assert.Errorf(t, err, "some error")
	})
	t.Run("unexpected transactions begin", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			fmt.Println("error creating mock database")
			return
		}
		if _, err := db.Begin(); err == nil {
			t.Error("an error was expected when calling begin, but got none")
		}
		mock.ExpectBegin()
		db.Begin()
		repo := mysql.UserRepo{Db: db}
		err = repo.CreateUser(userFoo)
		assert.Errorf(t, err, "failed to start transaction")
	})
}

func TestUserRepo_UpdateUserByMock(t *testing.T) {
	t.Run("transaction start failed when update user", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			fmt.Println("error creating mock database")
			return
		}
		defer db.Close()
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := mysql.UserRepo{Db: db}
		err = repo.UpdateUser(userFoo)
		assert.Errorf(t, err, "some error")
	})
	t.Run("unexpected transactions begin", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			fmt.Println("error creating mock database")
			return
		}
		if _, err := db.Begin(); err == nil {
			t.Error("an error was expected when calling begin, but got none")
		}
		mock.ExpectBegin()
		db.Begin()
		repo := mysql.UserRepo{Db: db}
		err = repo.UpdateUser(userFoo)
		assert.Errorf(t, err, "failed to start transaction")
	})
}

func TestUserRepo_AddFollowerByMock(t *testing.T) {
	t.Run("transaction start failed when add follower", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			fmt.Println("error creating mock database")
			return
		}
		defer db.Close()
		mock.ExpectBegin()
		mock.MatchExpectationsInOrder(true)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "password"}).
			AddRow(userFoo.ID, userFoo.Username, userFoo.Email, userFoo.Password))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectCommit()
		repo := mysql.UserRepo{Db: db}
		err = repo.AddFollower(userFoo, userBar)
		assert.Errorf(t, err, "some error")
	})
	t.Run("unexpected transactions begin", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			fmt.Println("error creating mock database")
			return
		}
		if _, err := db.Begin(); err == nil {
			t.Error("an error was expected when calling begin, but got none")
		}
		mock.ExpectBegin()
		db.Begin()
		repo := mysql.UserRepo{Db: db}
		err = repo.AddFollower(userFoo, userBar)
		assert.Errorf(t, err, "failed to start transaction")
	})
}

func TestUerRepo_RemoveFollowerByMock(t *testing.T) {
	t.Run("transaction start failed when remove follower", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			fmt.Println("error creating mock database")
			return
		}
		defer db.Close()
		mock.ExpectBegin()
		mock.MatchExpectationsInOrder(true)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "password"}).
			AddRow(userFoo.ID, userFoo.Username, userFoo.Email, userFoo.Password))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectCommit()
		repo := mysql.UserRepo{Db: db}
		err = repo.RemoveFollower(userFoo, userBar)
		assert.Errorf(t, err, "some error")
	})
	t.Run("unexpected transactions begin", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			fmt.Println("error creating mock database")
			return
		}
		if _, err := db.Begin(); err == nil {
			t.Error("an error was expected when calling begin, but got none")
		}
		mock.ExpectBegin()
		db.Begin()
		repo := mysql.UserRepo{Db: db}
		err = repo.RemoveFollower(userFoo, userBar)
		assert.Errorf(t, err, "failed to start transaction")
	})
}
