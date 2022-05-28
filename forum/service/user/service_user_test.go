package user

import (
	"fmt"
	. "forum/mock/repository"
	"forum/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUser_CheckUser(t *testing.T) {
	t.Run("when findbyemail return error", func(t *testing.T) {
		// Given
		userMock := NewIRepoUser(t)
		mockRequestUser := NewUserService(userMock)

		// When
		userMock.On("FindByEmail", mock.Anything).Return(nil, fmt.Errorf("findbyemail error"))
		// Then
		err := mockRequestUser.CheckUser(&model.LoginUser{Email: "foo@foo.com", Password: "foo"})
		assert.EqualError(t, err, "findbyemail error")
	})
	t.Run("when find by email return ok", func(t *testing.T) {
		// Given
		userMock := NewIRepoUser(t)
		mockRequestUser := NewUserService(userMock)

		// When
		userMock.On("FindByEmail", mock.Anything).Return(userFoo, nil)
		// Then
		err := mockRequestUser.CheckUser(&model.LoginUser{Email: "foo@foo.com", Password: "123456"})
		assert.Errorf(t, err, "crypto/bcrypt: hashedSecret too short to be a bcrypted password")
	})
}

func TestUser_CreateUser(t *testing.T) {
	t.Run("when create user return error", func(t *testing.T) {
		// Given
		userMock := NewIRepoUser(t)
		mockRequestUser := NewUserService(userMock)

		// When
		userMock.On("CreateUser", mock.Anything).Return(fmt.Errorf("create user error"))
		// Then
		err := mockRequestUser.CreateUser(&model.RegisterUser{Username: "foo", Email: "foo@foo.com", Password: "123456"})
		assert.EqualError(t, err, "create user error")
	})
	t.Run("when user password is empty", func(t *testing.T) {
		// Given
		userMock := NewIRepoUser(t)
		mockRequestUser := NewUserService(userMock)

		// When
		//userMock.On("CreateUser", mock.Anything).Return(nil)
		// Then
		err := mockRequestUser.CreateUser(&model.RegisterUser{Username: "foo", Email: "foo@foo.com", Password: ""})
		assert.Errorf(t, err, "password should not be empty")
	})
}

func TestUser_FollowUser(t *testing.T) {
	t.Run("when follow user return error", func(t *testing.T) {
		// Given
		userMock := NewIRepoUser(t)
		mockRequestUser := NewUserService(userMock)

		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)
		userMock.On("FindUserByID", mock.Anything).Return(userFoo, nil)
		userMock.On("AddFollower", mock.Anything, mock.Anything).Return(fmt.Errorf("add follower error"))
		// Then
		err := mockRequestUser.FollowUserByUserName(1, "foo")
		assert.EqualError(t, err, "add follower error")
	})
	t.Run("when FindUserByUserName return error", func(t *testing.T) {
		// Given
		userMock := NewIRepoUser(t)
		mockRequestUser := NewUserService(userMock)

		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(nil, fmt.Errorf("find user by username error"))
		// Then
		err := mockRequestUser.FollowUserByUserName(1, "foo")
		assert.EqualError(t, err, "find user by username error")
	})
	t.Run("when FindUserByID return error", func(t *testing.T) {
		// Given
		userMock := NewIRepoUser(t)
		mockRequestUser := NewUserService(userMock)

		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)
		userMock.On("FindUserByID", mock.Anything).Return(nil, fmt.Errorf("find user by id error"))
		// Then
		err := mockRequestUser.FollowUserByUserName(1, "foo")
		assert.EqualError(t, err, "find user by id error")
	})

}

func TestUser_GetUser(t *testing.T) {
	t.Run("when get user by id return error", func(t *testing.T) {
		// Given
		userMock := NewIRepoUser(t)
		mockRequestUser := NewUserService(userMock)

		// When
		userMock.On("FindUserByID", mock.Anything).Return(nil, fmt.Errorf("find user by id error"))
		// Then
		_, err := mockRequestUser.GetUserByID(1)
		assert.EqualError(t, err, "find user by id error")
	})
	t.Run("when get user by email return ok", func(t *testing.T) {
		// Given
		userMock := NewIRepoUser(t)
		mockRequestUser := NewUserService(userMock)

		// When
		userMock.On("FindByEmail", mock.Anything).Return(userFoo, nil)
		// Then
		_, err := mockRequestUser.GetUserByEmail("foo@foo.com")
		assert.NoError(t, err)
	})
	t.Run("when get user by username return ok", func(t *testing.T) {
		// Given
		userMock := NewIRepoUser(t)
		mockRequestUser := NewUserService(userMock)

		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)
		// Then
		_, err := mockRequestUser.GetUserByUserName("foo")
		assert.NoError(t, err)
	})

}

func TestUser_UnFollowUser(t *testing.T) {
	t.Run("when FindUserByUserName return error", func(t *testing.T) {
		// Given
		userMock := NewIRepoUser(t)
		mockRequestUser := NewUserService(userMock)

		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(nil, fmt.Errorf("find user by username error"))
		// Then
		err := mockRequestUser.UnFollowUserByUserName(1, "foo")
		assert.EqualError(t, err, "find user by username error")
	})
	t.Run("when FindUserByID return error", func(t *testing.T) {
		// Given
		userMock := NewIRepoUser(t)
		mockRequestUser := NewUserService(userMock)

		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)
		userMock.On("FindUserByID", mock.Anything).Return(nil, fmt.Errorf("find user by id error"))
		// Then
		err := mockRequestUser.UnFollowUserByUserName(1, "foo")
		assert.EqualError(t, err, "find user by id error")
	})
	t.Run("when UnFollowUser return error", func(t *testing.T) {
		// Given
		userMock := NewIRepoUser(t)
		mockRequestUser := NewUserService(userMock)

		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)
		userMock.On("FindUserByID", mock.Anything).Return(userFoo, nil)
		userMock.On("RemoveFollower", mock.Anything, mock.Anything).Return(fmt.Errorf("unfollow user error"))
		// Then
		err := mockRequestUser.UnFollowUserByUserName(1, "foo")
		assert.EqualError(t, err, "unfollow user error")
	})

}

func TestUser_GetFollowers(t *testing.T) {
	t.Run("when FindUserByID return error", func(t *testing.T) {
		// Given
		userMock := NewIRepoUser(t)
		mockRequestUser := NewUserService(userMock)

		// When
		userMock.On("FindUserByID", mock.Anything).Return(nil, fmt.Errorf("find user by id error"))
		// Then
		_, err := mockRequestUser.GetFollowersByUserID(1)
		assert.EqualError(t, err, "find user by id error")
	})
	t.Run("when GetFollowers return error", func(t *testing.T) {
		// Given
		userMock := NewIRepoUser(t)
		mockRequestUser := NewUserService(userMock)

		// When
		userMock.On("FindUserByID", mock.Anything).Return(userFoo, nil)
		userMock.On("GetFollowers", mock.Anything).Return(nil, fmt.Errorf("get followers error"))
		// Then
		_, err := mockRequestUser.GetFollowersByUserID(1)
		assert.EqualError(t, err, "get followers error")
	})
}
func TestUserGetFollowingUser(t *testing.T) {
	t.Run("when FindUserByID return error", func(t *testing.T) {
		// Given
		userMock := NewIRepoUser(t)
		mockRequestUser := NewUserService(userMock)

		// When
		userMock.On("FindUserByID", mock.Anything).Return(nil, fmt.Errorf("find user by id error"))
		// Then
		_, err := mockRequestUser.GetFollowingUser(1)
		assert.EqualError(t, err, "find user by id error")
	})
	t.Run("when GetFollowingUsers return error", func(t *testing.T) {
		// Given
		userMock := NewIRepoUser(t)
		mockRequestUser := NewUserService(userMock)

		// When
		userMock.On("FindUserByID", mock.Anything).Return(userFoo, nil)
		userMock.On("GetFollowingUsers", mock.Anything).Return(nil, fmt.Errorf("get following users error"))
		// Then
		_, err := mockRequestUser.GetFollowingUser(1)
		assert.EqualError(t, err, "get following users error")
	})
}
