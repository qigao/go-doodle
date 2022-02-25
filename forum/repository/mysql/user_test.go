package mysql

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepo_CreateUser(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when create user", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewUserRepo(db)
		err = repo.CreateUser(userFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when create user ", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewUserRepo(db)
		err = repo.CreateUser(userFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transactions begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewUserRepo(db)
		err = repo.CreateUser(userFoo)
		assert.Errorf(t, err, "failed to start transaction")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUserRepo_UpdateUser(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when update user", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewUserRepo(db)
		err = repo.UpdateUser(userFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transactions begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewUserRepo(db)
		err = repo.UpdateUser(userFoo)
		assert.Errorf(t, err, "failed to start transaction")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when update user", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewUserRepo(db)
		err = repo.UpdateUser(userFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUserRepo_AddFollower(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when add follower", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("insert into")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewUserRepo(db)
		err = repo.AddFollower(userFoo, userBar)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transactions begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewUserRepo(db)
		err = repo.AddFollower(userFoo, userBar)
		assert.Errorf(t, err, "failed to start transaction")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when add follower", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("insert into")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewUserRepo(db)
		err = repo.AddFollower(userFoo, userBar)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUerRepo_RemoveFollower(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when remove follower", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("delete")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewUserRepo(db)
		err = repo.RemoveFollower(userFoo, userBar)
		assert.Errorf(t, err, "some error")
	})
	t.Run("transactions begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewUserRepo(db)
		err = repo.RemoveFollower(userFoo, userBar)
		assert.Errorf(t, err, "failed to start transaction")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when remove follower", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("delete")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewUserRepo(db)
		err = repo.RemoveFollower(userFoo, userBar)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUserRepo_IsFollower(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when follower is true", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "bio", "image", "created_at", "updated_at", "deleted_at", "following_id"}).
			AddRow(2, "foo", "foo@foo.com", "foo-password", "foo desc", "http://foo.com/foo.jpg", null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), 2)
		userBar.ID = 2
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
			WillReturnRows(rows)
		repo := NewUserRepo(db)
		result, err := repo.IsFollower(userFoo, userBar)
		assert.NoError(t, err)
		assert.True(t, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when follower is false", func(t *testing.T) {
		userBar.ID = 2
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
			WillReturnError(fmt.Errorf("some error"))
		repo := NewUserRepo(db)
		result, err := repo.IsFollower(userFoo, userBar)
		assert.NoError(t, err)
		assert.False(t, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestUerRepo_GetFollowers(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when get followers", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "bio", "image", "created_at", "updated_at", "deleted_at", "following_id"}).
			AddRow(2, "foo", "foo@foo.com", "foo-password", "foo desc", "http://foo.com/foo.jpg", null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), 2).
			AddRow(2, "foo", "foo@foo.com", "foo-password", "foo desc", "http://foo.com/foo.jpg", null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), 2)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
			WillReturnRows(rows)
		repo := NewUserRepo(db)
		result, err := repo.GetFollowers(userFoo)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(result))
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when get followers with error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
			WillReturnError(fmt.Errorf("some error"))
		repo := NewUserRepo(db)
		result, err := repo.GetFollowers(userFoo)
		assert.Errorf(t, err, "some error")
		assert.Nil(t, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestUerRepo_GetFollowingUsers(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when get following users", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "bio", "image", "created_at", "updated_at", "deleted_at", "following_id"}).
			AddRow(2, "foo", "foo@foo.com", "foo-password", "foo desc", "http://foo.com/foo.jpg", null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), 2).
			AddRow(2, "foo", "foo@foo.com", "foo-password", "foo desc", "http://foo.com/foo.jpg", null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), 2)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
			WillReturnRows(rows)
		repo := NewUserRepo(db)
		result, err := repo.GetFollowingUsers(userFoo)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(result))
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when get following users with error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
			WillReturnError(fmt.Errorf("some error"))
		repo := NewUserRepo(db)
		result, err := repo.GetFollowingUsers(userFoo)
		assert.Errorf(t, err, "some error")
		assert.Nil(t, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestUserRepo_FindByID(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when find by id", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "bio", "image", "created_at", "updated_at", "deleted_at", "following_id"}).
			AddRow(2, "foo", "foo@foo.com", "foo-password", "foo desc", "http://foo.com/foo.jpg", null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), 2)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
			WillReturnRows(rows)
		repo := NewUserRepo(db)
		result, err := repo.FindByID(2)
		assert.NoError(t, err)
		assert.Equal(t, uint64(2), result.ID)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when find by id with error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
			WillReturnError(fmt.Errorf("some error"))
		repo := NewUserRepo(db)
		result, err := repo.FindByID(2)
		assert.Errorf(t, err, "some error")
		assert.Nil(t, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUserRepo_FindByEmail(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when find by email", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "bio", "image", "created_at", "updated_at", "deleted_at", "following_id"}).
			AddRow(2, "foo", "foo@foo.com", "foo-password", "foo desc", "http://foo.com/foo.jpg", null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), 2)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
			WillReturnRows(rows)
		repo := NewUserRepo(db)
		result, err := repo.FindByEmail("foo@foo.com")
		assert.NoError(t, err)
		assert.Equal(t, uint64(2), result.ID)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when find by email with error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
			WillReturnError(fmt.Errorf("some error"))
		repo := NewUserRepo(db)
		result, err := repo.FindByEmail("foo@foo.com")
		assert.Errorf(t, err, "some error")
		assert.Nil(t, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestUserRepo_FindByUserName(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when find by username", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "bio", "image", "created_at", "updated_at", "deleted_at", "following_id"}).
			AddRow(2, "foo", "foo@foo.com", "foo-password", "foo desc", "http://foo.com/foo.jpg", null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), 2)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
			WillReturnRows(rows)
		repo := NewUserRepo(db)
		result, err := repo.FindUserByUserName("foo")
		assert.NoError(t, err)
		assert.Equal(t, uint64(2), result.ID)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when find by username with error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
			WillReturnError(fmt.Errorf("some error"))
		repo := NewUserRepo(db)
		result, err := repo.FindUserByUserName("foo")
		assert.Errorf(t, err, "some error")
		assert.Nil(t, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
