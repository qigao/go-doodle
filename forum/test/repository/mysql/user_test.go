package mysql

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepo_User(t *testing.T) {
	t.Run("Create a user", func(t *testing.T) {
		userFoo.ID = 7
		err := userRepo.CreateUser(userFoo)
		assert.NoError(t, err)
	})
	t.Run("find user by id", func(t *testing.T) {
		user, err := userRepo.FindByID(7)
		assert.NoError(t, err)
		assert.Equal(t, userFoo.ID, user.ID)
		assert.Equal(t, userFoo.Username, user.Username)
		assert.Equal(t, userFoo.Email, user.Email)
		assert.Equal(t, userFoo.Password, user.Password)
	})
	t.Run("find user by id return nil", func(t *testing.T) {
		user, err := userRepo.FindByID(77)
		assert.Error(t, sql.ErrNoRows, err)
		assert.Nil(t, user)
	})
	t.Run("find user by email", func(t *testing.T) {
		user, err := userRepo.FindByEmail(userFoo.Email)
		assert.NoError(t, err)
		assert.Equal(t, userFoo.ID, user.ID)
		assert.Equal(t, userFoo.Username, user.Username)
		assert.Equal(t, userFoo.Email, user.Email)
		assert.Equal(t, userFoo.Password, user.Password)
	})
	t.Run("find user by email return nil", func(t *testing.T) {
		user, err := userRepo.FindByEmail("foo1")
		assert.Error(t, sql.ErrNoRows, err)
		assert.Nil(t, user)
	})

	t.Run("find user by username", func(t *testing.T) {
		user, err := userRepo.FindByUserName(userFoo.Username)
		assert.NoError(t, err)
		assert.Equal(t, userFoo.ID, user.ID)
		assert.Equal(t, userFoo.Username, user.Username)
		assert.Equal(t, userFoo.Email, user.Email)
		assert.Equal(t, userFoo.Password, user.Password)
	})
	t.Run("find user by username return nil", func(t *testing.T) {
		user, err := userRepo.FindByUserName("foo1")
		assert.Error(t, sql.ErrNoRows, err)
		assert.Nil(t, user)
	})
	t.Run("Create a user return error", func(t *testing.T) {
		userFoo.ID = 7
		err := userRepo.CreateUser(userFoo)
		assert.Error(t, err)
	})
}
func TestUserRepo_UpdateUser(t *testing.T) {
	t.Run("Create a user", func(t *testing.T) {
		userFoo.ID = 17
		err := userRepo.CreateUser(userFoo)
		assert.NoError(t, err)
	})
	t.Run("Update user", func(t *testing.T) {
		userFoo.ID = 17
		userFoo.Password = "abcd"
		err := userRepo.UpdateUser(userFoo)
		assert.NoError(t, err)
	})
	t.Run("find user by id", func(t *testing.T) {
		user, err := userRepo.FindByID(17)
		assert.NoError(t, err)
		assert.Equal(t, userFoo.ID, user.ID)
		assert.Equal(t, userFoo.Username, user.Username)
		assert.Equal(t, userFoo.Email, user.Email)
		assert.Equal(t, userFoo.Password, user.Password)
	})
	t.Run("update not exits user", func(t *testing.T) {
		userFoo.ID = 77
		err := userRepo.UpdateUser(userFoo)
		assert.Error(t, sql.ErrNoRows, err)
	})
	t.Run("find user by id", func(t *testing.T) {
		user, err := userRepo.FindByID(77)
		assert.Error(t, sql.ErrNoRows, err)
		assert.Nil(t, user)
	})
}
func TestUserRepo_Follower(t *testing.T) {
	t.Run("Create a user", func(t *testing.T) {
		userFoo.ID = 27
		err := userRepo.CreateUser(userFoo)
		assert.NoError(t, err)
	})
	t.Run("Create another user", func(t *testing.T) {
		userFoo.ID = 28
		err := userRepo.CreateUser(userFoo)
		assert.NoError(t, err)
	})
	t.Run("add follower", func(t *testing.T) {
		userFoo.ID = 27
		userBar.ID = 28
		err := userRepo.AddFollower(userFoo, userBar)
		assert.NoError(t, err)
	})
	t.Run("is follower", func(t *testing.T) {
		isFollower, err := userRepo.IsFollower(userFoo, userBar)
		assert.NoError(t, err)
		assert.True(t, isFollower)
	})
	t.Run("get followers", func(t *testing.T) {
		users, err := userRepo.GetFollowers(userFoo)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(users))
		assert.Equal(t, userBar.ID, users[0].ID)
	})
	t.Run("get following user", func(t *testing.T) {
		users, err := userRepo.GetFollowingUsers(userBar)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(users))
		assert.Equal(t, userFoo.ID, users[0].ID)
	})
	t.Run("remove follower", func(t *testing.T) {
		userFoo.ID = 27
		err := userRepo.RemoveFollower(userFoo, userBar)
		assert.NoError(t, err)
	})
	t.Run("is not follower", func(t *testing.T) {
		isFollower, err := userRepo.IsFollower(userFoo, userBar)
		assert.Error(t, sql.ErrNoRows, err)
		assert.False(t, isFollower)
	})
	t.Run("get no followers", func(t *testing.T) {
		users, err := userRepo.GetFollowers(userFoo)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(users))
	})
	t.Run("get no following user", func(t *testing.T) {
		users, err := userRepo.GetFollowingUsers(userFoo)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(users))
	})
	t.Run("user not exits is not follower", func(t *testing.T) {
		isFollower, err := userRepo.IsFollower(userFoo, userBar)
		assert.Error(t, sql.ErrNoRows, err)
		assert.False(t, isFollower)
	})
	t.Run("user itself is not follower", func(t *testing.T) {
		isFollower, err := userRepo.IsFollower(userFoo, userFoo)
		assert.Error(t, sql.ErrNoRows, err)
		assert.False(t, isFollower)
	})
	t.Run("not exists user has no follower", func(t *testing.T) {
		userBar.ID = 29
		isFollower, err := userRepo.IsFollower(userBar, userFoo)
		assert.Error(t, sql.ErrNoRows, err)
		assert.False(t, isFollower)
	})
}
