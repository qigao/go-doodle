package repository

import (
	"fmt"
	"gforum/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {

	t.Run("Create a record ", func(t *testing.T) {

		err := userRepo.Create(userFoo)
		assert.NoError(t, err)
	})
	t.Run("Get a record By Email", func(t *testing.T) {
		user, err := userRepo.GetByEmail("foo@foo.com")
		require.NoError(t, err)
		assert.Equal(t, "foo@foo.com", user.Email)
	})
	t.Run("Get a record By Username", func(t *testing.T) {
		user, err := userRepo.GetByUsername("foo")
		require.NoError(t, err)
		assert.Equal(t, "foo", user.Username)
	})
	t.Run("Get a record By ID", func(t *testing.T) {
		user, err := userRepo.GetByID(1)
		require.NoError(t, err)
		assert.Equal(t, "foo", user.Username)
	})
	t.Run("Update user info", func(t *testing.T) {
		user, err := userRepo.GetByID(1)
		require.NoError(t, err)
		user.Password = "654321"
		err = userRepo.Update(user)
		if err != nil {
			t.Error(err)
		}
		user, err = userRepo.GetByID(1)
		require.NoError(t, err)
		assert.Equal(t, "654321", user.Password)
	})
}

func TestUserRepository_Follow(t *testing.T) {
	t.Run("Create userB", func(t *testing.T) {
		err := userRepo.Create(userBar)
		assert.NoError(t, err)
	})
	t.Run("Follow a user", func(t *testing.T) {
		err := userRepo.Follow(userFoo, userBar)
		assert.NoError(t, err)
	})
	t.Run("userB is following userA", func(t *testing.T) {
		following, err := userRepo.IsFollowing(userFoo, userBar)
		require.NoError(t, err)
		assert.Equal(t, true, following)
	})
	t.Run("following id of userA", func(t *testing.T) {
		following, err := userRepo.GetFollowingUserIDs(userFoo)
		require.NoError(t, err)
		assert.Equal(t, 1, len(following))
	})
	t.Run("unfollow a user", func(t *testing.T) {
		err := userRepo.UnFollow(userFoo, userBar)
		assert.NoError(t, err)
	})
	t.Run("userB is not following userA", func(t *testing.T) {
		following, err := userRepo.IsFollowing(userFoo, userBar)
		require.NoError(t, err)
		assert.Equal(t, false, following)
	})
}

func TestUserRepository_GetUser_Return_Error(t *testing.T) {
	t.Run("Get a record By Email", func(t *testing.T) {
		_, err := userRepo.GetByEmail("foo1@foo.com")
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
	t.Run("Get a record By Username", func(t *testing.T) {
		_, err := userRepo.GetByUsername("foo1")
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
	t.Run("Get a record By ID", func(t *testing.T) {
		_, err := userRepo.GetByID(11)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
}

func TestUserRepository_IsFollowing_Return_Error(t *testing.T) {

	t.Run("user  is null", func(t *testing.T) {
		result, err := userRepo.IsFollowing(nil, nil)
		assert.False(t, false, result)
		require.NoError(t, err)
	})
	t.Run("userB is following userA", func(t *testing.T) {
		result, err := userRepo.IsFollowing(userFoo, userBar)
		assert.False(t, false, result)
		assert.Nil(t, nil, err)
	})
}

func TestUserRepository_IsFollowing_WithMock(t *testing.T) {
	gormDB, mock := mockMysql()
	repo := repository.NewUserRepository(gormDB)
	mock.MatchExpectationsInOrder(true)

	t.Run("mock return error", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WithArgs(1, 0).WillReturnError(fmt.Errorf("some error"))
		result, err := repo.IsFollowing(userFoo, userFoo)
		require.False(t, result)
		assert.Error(t, err)
	})
}

func TestUserRepository_GetFollowingUserIDs_WithMock(t *testing.T) {
	gormDB, mock := mockMysql()
	repo := repository.NewUserRepository(gormDB)
	mock.MatchExpectationsInOrder(true)
	t.Run("Get following userids", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WithArgs(userFoo.ID).WillReturnError(
			fmt.Errorf("some error"))
		result, err := repo.GetFollowingUserIDs(userFoo)
		require.Equal(t, []uint{}, result)
		require.NotNil(t, err)
	})
}
