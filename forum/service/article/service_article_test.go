package article

import (
	"fmt"
	"schema/entity"
	"testing"

	mockRepo "forum/mock/repository"

	"github.com/stretchr/testify/mock"
	"github.com/volatiletech/null/v8"
	"gotest.tools/assert"
)

func TestArticle_CreateArticle(t *testing.T) {
	t.Run("When CreateArticle, CreateArticle failed with error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)

		// When
		articleMock.On("CreateArticle", mock.Anything).Return(fmt.Errorf("CreateArticle error"))
		// Then
		err := ServiceArticleMock.CreateArticle(articleFoo)
		assert.Error(t, err, "CreateArticle error")
	})
}

func TestArticle_DeleteArticle(t *testing.T) {
	t.Run("When Find Article get error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(nil, fmt.Errorf("FindArticleBySlug error"))
		// Then
		err := ServiceArticleMock.DeleteArticle("slug")
		assert.Error(t, err, "FindArticleBySlug error")
	})
	t.Run("when delete article get error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("DeleteArticle", mock.Anything).Return(fmt.Errorf("DeleteArticle error"))
		// Then
		err := ServiceArticleMock.DeleteArticle("slug")
		assert.Error(t, err, "DeleteArticle error")
	})
	t.Run("when delete article return ok", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("DeleteArticle", mock.Anything).Return(nil)
		// Then
		err := ServiceArticleMock.DeleteArticle("slug")
		assert.NilError(t, err)
	})
}

func TestArticle_FindArticle(t *testing.T) {
	t.Run("When Find Article get error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(nil, fmt.Errorf("FindArticleBySlug error"))
		// Then
		_, _, _, err := ServiceArticleMock.FindArticle("slug")
		assert.Error(t, err, "FindArticleBySlug error")
	})
	t.Run("When find author get error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("FindAuthorByArticle", mock.Anything).Return(nil, fmt.Errorf("FindAuthorByArticle error"))
		// Then
		_, _, _, err := ServiceArticleMock.FindArticle("slug")
		assert.Error(t, err, "FindAuthorByArticle error")
	})
	t.Run("When find tag get error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("FindAuthorByArticle", mock.Anything).Return(userFoo, nil)
		articleMock.On("FindTagsByArticle", mock.Anything).Return(nil, fmt.Errorf("FindTagsByArticle error"))
		// Then
		_, _, _, err := ServiceArticleMock.FindArticle("slug")
		assert.Error(t, err, "FindTagsByArticle error")
	})
	t.Run("When find article return ok", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("FindAuthorByArticle", mock.Anything).Return(userFoo, nil)
		articleMock.On("FindTagsByArticle", mock.Anything).Return(nil, nil)

		// Then
		_, _, _, err := ServiceArticleMock.FindArticle("slug")
		assert.NilError(t, err)
	})
}

func TestArticle_FindArticleByAuthor(t *testing.T) {
	t.Run("When Find user by username return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(nil, fmt.Errorf("FindUserByUsername error"))

		// Then
		_, n, err := ServiceArticleMock.FindArticleByAuthor("username", 0, 1)
		assert.Error(t, err, "FindUserByUsername error")
		assert.Equal(t, n, int64(0))
	})
	t.Run("When Find article by author return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)
		articleMock.On("ListArticlesByAuthor", mock.Anything, mock.Anything, mock.Anything).Return(nil, int64(0), fmt.Errorf("FindArticleByAuthor error"))

		// Then
		_, n, err := ServiceArticleMock.FindArticleByAuthor("username", 0, 1)
		assert.Error(t, err, "FindArticleByAuthor error")
		assert.Equal(t, n, int64(0))
	})
	t.Run("when find article return ok", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)
		articleMock.On("ListArticlesByAuthor", mock.Anything, mock.Anything, mock.Anything).Return([]*entity.Article{articleFoo}, int64(1), nil)

		// Then
		_, n, err := ServiceArticleMock.FindArticleByAuthor("username", 0, 1)
		assert.NilError(t, err)
		assert.Equal(t, n, int64(1))
	})
}

func TestArticle_FindArticles(t *testing.T) {
	t.Run("When Find articles return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(nil, fmt.Errorf("FindUserByUsername error"))

		// Then
		_, n, err := ServiceArticleMock.FindArticles("test-tag", "test-user", 0, 1)
		assert.Error(t, err, "FindUserByUsername error")
		assert.Equal(t, n, int64(0))
	})
	t.Run("When Find articles by tag return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)
		articleMock.On("ListArticlesByTag", mock.Anything, mock.Anything, mock.Anything).Return(nil, int64(0), fmt.Errorf("FindArticleByTag error"))
		// Then
		_, n, err := ServiceArticleMock.FindArticles("test-tag", "test-user", 0, 1)
		assert.Error(t, err, "FindArticleByTag error")
		assert.Equal(t, n, int64(0))
	})
	t.Run("When Find articles by tag return OK", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)
		articleMock.On("ListArticlesByTag", mock.Anything, mock.Anything, mock.Anything).Return([]*entity.Article{articleBar}, int64(1), nil)
		// Then
		_, n, err := ServiceArticleMock.FindArticles("test-tag", "test-user", 0, 1)
		assert.NilError(t, err)
		assert.Equal(t, n, int64(1))
	})
	t.Run("When find articles by author return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)

		articleMock.On("ListArticlesByAuthor", mock.Anything, mock.Anything, mock.Anything).Return(nil, int64(0), fmt.Errorf("FindArticleByAuthor error"))
		// Then
		_, n, err := ServiceArticleMock.FindArticles("", "test-user", 0, 1)
		assert.Error(t, err, "FindArticleByAuthor error")
		assert.Equal(t, n, int64(0))
	})
	t.Run("When find articles by author return OK", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)

		articleMock.On("ListArticlesByAuthor", mock.Anything, mock.Anything, mock.Anything).Return([]*entity.Article{articleBar}, int64(1), nil)
		// Then
		_, n, err := ServiceArticleMock.FindArticles("", "test-user", 0, 1)
		assert.NilError(t, err)
		assert.Equal(t, n, int64(1))
	})
	t.Run("When find articles without user return ok", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)
		articleMock.On("FindArticles", mock.Anything, mock.Anything).Return([]*entity.Article{articleBar}, int64(1), nil)
		// Then
		_, n, err := ServiceArticleMock.FindArticles("", "", 0, 1)
		assert.NilError(t, err)
		assert.Equal(t, n, int64(1))
	})
	t.Run("When find articles without user get error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)
		articleMock.On("FindArticles", mock.Anything, mock.Anything).Return([]*entity.Article{articleBar}, int64(1), fmt.Errorf("FindArticle error"))
		// Then
		_, n, err := ServiceArticleMock.FindArticles("", "", 0, 1)
		assert.Error(t, err, "FindArticle error")
		assert.Equal(t, n, int64(0))
	})
}

func TestArticle_FindCommentsBySlug(t *testing.T) {
	t.Run("When find article by slug return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(nil, fmt.Errorf("FindArticleBySlug error"))

		// Then
		_, err := ServiceArticleMock.FindCommentsBySlug("test-slug", 0, 1)
		assert.Error(t, err, "FindArticleBySlug error")
	})
	t.Run("When find comments by article return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleBar, nil)
		articleMock.On("FindCommentsByArticle", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, fmt.Errorf("FindCommentsBySlug error"))
		// Then
		comments, err := ServiceArticleMock.FindCommentsBySlug("test-slug", 0, 1)
		assert.Error(t, err, "FindCommentsBySlug error")
		assert.Equal(t, len(comments), 0)
	})
	t.Run("When Find comments by slug return ok", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleBar, nil)
		articleMock.On("FindCommentsByArticle", mock.Anything, mock.Anything, mock.Anything).
			Return([]*entity.Comment{commentFoo}, nil)
		// Then
		comments, err := ServiceArticleMock.FindCommentsBySlug("test-slug", 0, 1)
		assert.NilError(t, err)
		assert.Equal(t, len(comments), 1)
	})
}

func TestArticle_FindAuthorBySlug(t *testing.T) {
	t.Run("when find article by slug return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(nil, fmt.Errorf("FindArticleBySlug error"))

		// Then
		_, err := ServiceArticleMock.FindAuthorBySlug("test-slug")
		assert.Error(t, err, "FindArticleBySlug error")
	})
	t.Run("when find author by article return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleBar, nil)
		articleMock.On("FindAuthorByArticle", mock.Anything).Return(nil, fmt.Errorf("FindAuthorBySlug error"))

		// Then
		_, err := ServiceArticleMock.FindAuthorBySlug("test-slug")
		assert.Error(t, err, "FindAuthorBySlug error")
	})
	t.Run("when find author by article return ok", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleBar, nil)
		articleMock.On("FindAuthorByArticle", mock.Anything).Return(userFoo, nil)
		// Then
		author, err := ServiceArticleMock.FindAuthorBySlug("test-slug")
		assert.NilError(t, err)
		assert.Equal(t, author.ID, userFoo.ID)
	})
}

func TestArticle_AddCommentToArticle(t *testing.T) {
	t.Run("when Find article by slug return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(nil, fmt.Errorf("FindArticleBySlug error"))
		// Then
		err := ServiceArticleMock.AddCommentToArticle("test-slug", commentFoo)
		assert.Error(t, err, "FindArticleBySlug error")
	})
	t.Run("when add comment return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("AddComment", mock.Anything, mock.Anything).Return(fmt.Errorf("AddComment error"))
		// Then
		err := ServiceArticleMock.AddCommentToArticle("test-slug", commentFoo)
		assert.Error(t, err, "AddComment error")
	})
	t.Run("when add comment return ok", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("AddComment", mock.Anything, mock.Anything).Return(nil)
		// Then
		err := ServiceArticleMock.AddCommentToArticle("test-slug", commentFoo)
		assert.NilError(t, err)
	})
}

func TestArticle_DeleteCommentFromArticle(t *testing.T) {
	t.Run("when Find article by slug return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(nil, fmt.Errorf("FindArticleBySlug error"))
		// Then
		err := ServiceArticleMock.DeleteCommentFromArticle("test-slug", commentFoo.ID)
		assert.Error(t, err, "FindArticleBySlug error")
	})
	t.Run("when Find comment return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("FindCommentByID", mock.Anything).Return(nil, fmt.Errorf("FindCommentByID error"))
		// Then
		err := ServiceArticleMock.DeleteCommentFromArticle("test-slug", commentFoo.ID)
		assert.Error(t, err, "FindCommentByID error")
	})
	t.Run("when delete comment return ok", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("FindCommentByID", mock.Anything).Return(commentFoo, nil)
		articleMock.On("DeleteCommentByArticle", mock.Anything, mock.Anything).Return(fmt.Errorf("DeleteCommentByArticle error"))
		// Then
		err := ServiceArticleMock.DeleteCommentFromArticle("test-slug", commentFoo.ID)
		assert.Error(t, err, "DeleteCommentByArticle error")
	})
	t.Run("when delete comment return ok", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("FindCommentByID", mock.Anything).Return(commentFoo, nil)
		articleMock.On("DeleteCommentByArticle", mock.Anything, mock.Anything).Return(nil)

		// Then
		err := ServiceArticleMock.DeleteCommentFromArticle("test-slug", commentFoo.ID)
		assert.NilError(t, err)
	})
}

func TestArticle_AddFavoriteArticleBySlug(t *testing.T) {
	t.Run("when Find article by slug return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(nil, fmt.Errorf("FindArticleBySlug error"))
		// Then
		err := ServiceArticleMock.AddFavoriteArticleBySlug("test-slug", 1)
		assert.Error(t, err, "FindArticleBySlug error")
	})
	t.Run("when find user by id return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		userMock.On("FindUserByID", mock.Anything).Return(nil, fmt.Errorf("FindUserByID error"))
		// Then
		err := ServiceArticleMock.AddFavoriteArticleBySlug("test-slug", 1)
		assert.Error(t, err, "FindUserByID error")
	})
	t.Run("when find user by id return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		userMock.On("FindUserByID", mock.Anything).Return(userBar, nil)
		articleMock.On("AddFavoriteArticle", mock.Anything, mock.Anything).Return(fmt.Errorf("AddFavoriteArticle error"))
		// Then
		err := ServiceArticleMock.AddFavoriteArticleBySlug("test-slug", 1)
		assert.Error(t, err, "AddFavoriteArticle error")
	})
	t.Run("when add favorite article return ok", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		userMock.On("FindUserByID", mock.Anything).Return(userBar, nil)
		articleMock.On("AddFavoriteArticle", mock.Anything, mock.Anything).Return(nil)
		// Then
		err := ServiceArticleMock.AddFavoriteArticleBySlug("test-slug", 1)
		assert.NilError(t, err)
	})
}

func TestArticle_RemoveFavoriteArticleBySlug(t *testing.T) {
	t.Run("when Find article by slug return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(nil, fmt.Errorf("FindArticleBySlug error"))
		// Then
		err := ServiceArticleMock.RemoveFavoriteArticleBySlug("test-slug", 1)
		assert.Error(t, err, "FindArticleBySlug error")
	})
	t.Run("when find user by id return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		userMock.On("FindUserByID", mock.Anything).Return(nil, fmt.Errorf("FindUserByID error"))
		// Then
		err := ServiceArticleMock.AddFavoriteArticleBySlug("test-slug", 1)
		assert.Error(t, err, "FindUserByID error")
	})
	t.Run("when find user by id return error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		userMock.On("FindUserByID", mock.Anything).Return(userBar, nil)
		articleMock.On("RemoveFavorite", mock.Anything, mock.Anything).Return(fmt.Errorf("RemoveFavorite error"))
		// Then
		err := ServiceArticleMock.RemoveFavoriteArticleBySlug("test-slug", 1)
		assert.Error(t, err, "RemoveFavorite error")
	})
	t.Run("when add favorite article return ok", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		userMock.On("FindUserByID", mock.Anything).Return(userBar, nil)
		articleMock.On("RemoveFavorite", mock.Anything, mock.Anything).Return(nil)
		// Then
		err := ServiceArticleMock.RemoveFavoriteArticleBySlug("test-slug", 1)
		assert.NilError(t, err)
	})
}

func TestArticle_AddTagToArticle(t *testing.T) {
	t.Run("When FindArticle failed with error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, fmt.Errorf("FindArticleBySlug error"))
		// Then
		err := ServiceArticleMock.AddTagToArticle("slug-test", []string{"tag2"})
		assert.Error(t, err, "FindArticleBySlug error")
	})
	t.Run("When ListTags failed with error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		tag1 := &entity.Tag{
			Tag: null.StringFrom("tag1"),
		}
		tag2 := &entity.Tag{
			Tag: null.StringFrom("tag2"),
		}
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("ListTags").Return([]*entity.Tag{tag1, tag2}, fmt.Errorf("ListTags error"))
		// Then
		err := ServiceArticleMock.AddTagToArticle("slug-test", []string{"tag2"})
		assert.Error(t, err, "ListTags error")
	})

	t.Run("When AddTagToArticle failed with error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		tag1 := &entity.Tag{
			Tag: null.StringFrom("tag1"),
		}
		tag2 := &entity.Tag{
			Tag: null.StringFrom("tag2"),
		}
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("ListTags").Return([]*entity.Tag{tag1, tag2}, nil)
		articleMock.On("AddTagsToArticle", mock.Anything, mock.Anything).Return(fmt.Errorf("AddTagsToArticle error"))
		// Then
		err := ServiceArticleMock.AddTagToArticle("slug-test", []string{"tag2"})
		assert.Error(t, err, "AddTagsToArticle error")
	})
	t.Run("When AddTagToArticle return ok", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		tag1 := &entity.Tag{
			Tag: null.StringFrom("tag1"),
		}
		tag2 := &entity.Tag{
			Tag: null.StringFrom("tag2"),
		}
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("ListTags").Return([]*entity.Tag{tag1, tag2}, nil)
		articleMock.On("AddTagsToArticle", mock.Anything, mock.Anything).Return(nil)
		// Then
		err := ServiceArticleMock.AddTagToArticle("slug-test", []string{"tag2"})
		assert.NilError(t, err)
	})
}

func TestArticle_UpdateArticle(t *testing.T) {
	t.Run("When FindArticleBySlug failed with error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, fmt.Errorf("FindArticleBySlug error"))
		// Then
		err := ServiceArticleMock.UpdateArticle("slug-test", articleFoo)
		assert.Error(t, err, "FindArticleBySlug error")
	})
	t.Run("When Update article failed with error", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("UpdateArticle", mock.Anything).Return(fmt.Errorf("update article error"))
		// Then
		err := ServiceArticleMock.UpdateArticle("slug-test", articleFoo)
		assert.Error(t, err, "update article error")
	})
	t.Run("When update article return OK", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("UpdateArticle", mock.Anything).Return(nil)
		// Then
		err := ServiceArticleMock.UpdateArticle("slug-test", articleFoo)
		assert.NilError(t, err)
	})
	t.Run("When update article return OK, body is not valid", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleFoo.Body = null.StringFrom("")
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("UpdateArticle", mock.Anything).Return(nil)
		// Then
		err := ServiceArticleMock.UpdateArticle("slug-test", articleFoo)
		assert.NilError(t, err)
	})
	t.Run("When update article return OK, description is not valid", func(t *testing.T) {
		// Given
		userMock := mockRepo.NewIRepoUser(t)
		articleMock := mockRepo.NewIRepoArticle(t)
		ServiceArticleMock := NewServiceArticle(articleMock, userMock)
		// When
		articleFoo.Description = null.StringFrom("")
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("UpdateArticle", mock.Anything).Return(nil)
		// Then
		err := ServiceArticleMock.UpdateArticle("slug-test", articleFoo)
		assert.NilError(t, err)
	})
}
