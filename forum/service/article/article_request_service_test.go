package article

import (
	"fmt"
	"forum/entity"
	. "forum/mock/repository"
	"testing"

	mock "github.com/stretchr/testify/mock"
	"github.com/volatiletech/null/v8"
	"gotest.tools/assert"
)

func newArticleMock() *Article {
	return &Article{
		Mock: mock.Mock{},
	}
}
func newUserMock() *User {
	return &User{
		Mock: mock.Mock{},
	}
}

func TestArticle_InsertArticleWithTags(t *testing.T) {
	t.Run("When InsertArticleWithTags, CreateArticle failed with error", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)

		// When
		articleMock.On("CreateArticle", mock.Anything).Return(fmt.Errorf("CreateArticle error"))
		// Then
		err := mockRequestArticle.InsertArticleWithTags(articleFoo, []string{"tag1", "tag2"})
		assert.Error(t, err, "CreateArticle error")
	})
	t.Run("When InsertArticleWithTags, AddTagToArticle failed with error", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		tag1 := &entity.Tag{
			Tag: null.StringFrom("tag1"),
		}
		tag2 := &entity.Tag{
			Tag: null.StringFrom("tag2"),
		}
		// When
		articleMock.On("CreateArticle", mock.Anything).Return(nil)
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("ListTags").Return([]*entity.Tag{tag1, tag2}, nil)
		articleMock.On("AddTagsToArticle", mock.Anything, mock.Anything).Return(fmt.Errorf("AddTagToArticle error"))
		// Then
		err := mockRequestArticle.InsertArticleWithTags(articleFoo, []string{"tag2"})
		assert.Error(t, err, "AddTagToArticle error")
	})
	t.Run("When InsertArticleWithTags, without error", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		tag1 := &entity.Tag{
			Tag: null.StringFrom("tag1"),
		}
		tag2 := &entity.Tag{
			Tag: null.StringFrom("tag2"),
		}
		// When
		articleMock.On("CreateArticle", mock.Anything).Return(nil)
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("ListTags").Return([]*entity.Tag{tag1, tag2}, nil)
		articleMock.On("AddTagsToArticle", mock.Anything, mock.Anything).Return(nil)
		// Then
		err := mockRequestArticle.InsertArticleWithTags(articleFoo, []string{"tag2"})
		assert.NilError(t, err)
	})
}

func TestArticle_DeleteArtile(t *testing.T) {
	t.Run("When Find Article get error", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(nil, fmt.Errorf("FindArticleBySlug error"))
		// Then
		err := mockRequestArticle.DeleteArticle("slug")
		assert.Error(t, err, "FindArticleBySlug error")
	})
	t.Run("when delete article get error", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("DeleteArticle", mock.Anything).Return(fmt.Errorf("DeleteArticle error"))
		// Then
		err := mockRequestArticle.DeleteArticle("slug")
		assert.Error(t, err, "DeleteArticle error")
	})
	t.Run("when delete article return ok", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("DeleteArticle", mock.Anything).Return(nil)
		// Then
		err := mockRequestArticle.DeleteArticle("slug")
		assert.NilError(t, err)
	})
}

func TestArticle_FindArticle(t *testing.T) {
	t.Run("When Find Article get error", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(nil, fmt.Errorf("FindArticleBySlug error"))
		// Then
		_, _, _, err := mockRequestArticle.FindArticle("slug")
		assert.Error(t, err, "FindArticleBySlug error")
	})
	t.Run("When find author get error", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("FindAuthorByArticle", mock.Anything).Return(nil, fmt.Errorf("FindAuthorByArticle error"))
		// Then
		_, _, _, err := mockRequestArticle.FindArticle("slug")
		assert.Error(t, err, "FindAuthorByArticle error")
	})
	t.Run("When find tag get error", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("FindAuthorByArticle", mock.Anything).Return(userFoo, nil)
		articleMock.On("FindTagsByArticle", mock.Anything).Return(nil, fmt.Errorf("FindTagsByArticle error"))
		// Then
		_, _, _, err := mockRequestArticle.FindArticle("slug")
		assert.Error(t, err, "FindTagsByArticle error")
	})
	t.Run("When find article return ok", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		// When
		articleMock.On("FindArticleBySlug", mock.Anything).Return(articleFoo, nil)
		articleMock.On("FindAuthorByArticle", mock.Anything).Return(userFoo, nil)
		articleMock.On("FindTagsByArticle", mock.Anything).Return(nil, nil)

		// Then
		_, _, _, err := mockRequestArticle.FindArticle("slug")
		assert.NilError(t, err)
	})
}

func TestArticle_FindArticleByAuthor(t *testing.T) {
	t.Run("When Find user by username return error", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(nil, fmt.Errorf("FindUserByUsername error"))

		// Then
		_, n, err := mockRequestArticle.FindArticleByAuthor("username", 0, 1)
		assert.Error(t, err, "FindUserByUsername error")
		assert.Equal(t, n, int64(0))
	})
	t.Run("When Find article by author return error", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)
		articleMock.On("ListArticlesByAuthor", mock.Anything, mock.Anything, mock.Anything).Return(nil, int64(0), fmt.Errorf("FindArticleByAuthor error"))

		// Then
		_, n, err := mockRequestArticle.FindArticleByAuthor("username", 0, 1)
		assert.Error(t, err, "FindArticleByAuthor error")
		assert.Equal(t, n, int64(0))
	})
	t.Run("when find article return ok", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)
		articleMock.On("ListArticlesByAuthor", mock.Anything, mock.Anything, mock.Anything).Return([]*entity.Article{articleFoo}, int64(1), nil)

		// Then
		_, n, err := mockRequestArticle.FindArticleByAuthor("username", 0, 1)
		assert.NilError(t, err)
		assert.Equal(t, n, int64(1))
	})

}

func TestArticle_FindArticles(t *testing.T) {
	t.Run("When Find articles return error", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(nil, fmt.Errorf("FindUserByUsername error"))

		// Then
		_, n, err := mockRequestArticle.FindArticles("test-tag", "test-user", 0, 1)
		assert.Error(t, err, "FindUserByUsername error")
		assert.Equal(t, n, int64(0))
	})
	t.Run("When Find articles by tag return error", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)
		articleMock.On("ListArticlesByTag", mock.Anything, mock.Anything, mock.Anything).Return(nil, int64(0), fmt.Errorf("FindArticleByTag error"))
		// Then
		_, n, err := mockRequestArticle.FindArticles("test-tag", "test-user", 0, 1)
		assert.Error(t, err, "FindArticleByTag error")
		assert.Equal(t, n, int64(0))
	})
	t.Run("When Find articles by tag return OK", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)
		articleMock.On("ListArticlesByTag", mock.Anything, mock.Anything, mock.Anything).Return([]*entity.Article{articleBar}, int64(1), nil)
		// Then
		_, n, err := mockRequestArticle.FindArticles("test-tag", "test-user", 0, 1)
		assert.NilError(t, err)
		assert.Equal(t, n, int64(1))
	})
	t.Run("When find artiles by author return error", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)

		articleMock.On("ListArticlesByAuthor", mock.Anything, mock.Anything, mock.Anything).Return(nil, int64(0), fmt.Errorf("FindArticleByAuthor error"))
		// Then
		_, n, err := mockRequestArticle.FindArticles("", "test-user", 0, 1)
		assert.Error(t, err, "FindArticleByAuthor error")
		assert.Equal(t, n, int64(0))
	})
	t.Run("When find artiles by author return OK", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)

		articleMock.On("ListArticlesByAuthor", mock.Anything, mock.Anything, mock.Anything).Return([]*entity.Article{articleBar}, int64(1), nil)
		// Then
		_, n, err := mockRequestArticle.FindArticles("", "test-user", 0, 1)
		assert.NilError(t, err)
		assert.Equal(t, n, int64(1))
	})
	t.Run("When find artiles without user return ok", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)
		articleMock.On("ListArticles", mock.Anything, mock.Anything).Return([]*entity.Article{articleBar}, int64(1), nil)
		// Then
		_, n, err := mockRequestArticle.FindArticles("", "", 0, 1)
		assert.NilError(t, err)
		assert.Equal(t, n, int64(1))
	})
	t.Run("When find artiles without user get orrer", func(t *testing.T) {
		// Given
		userMock := newUserMock()
		articleMock := newArticleMock()
		mockRequestArticle := NewRequestArticle(articleMock, userMock)
		// When
		userMock.On("FindUserByUserName", mock.Anything).Return(userFoo, nil)
		articleMock.On("ListArticles", mock.Anything, mock.Anything).Return([]*entity.Article{articleBar}, int64(1), fmt.Errorf("FindArticle error"))
		// Then
		_, n, err := mockRequestArticle.FindArticles("", "", 0, 1)
		assert.Error(t, err, "FindArticle error")
		assert.Equal(t, n, int64(0))
	})
}
