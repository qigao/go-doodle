package mysql

import (
	"database/sql"
	models "forum/entities"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
	"testing"
)

func TestArticleRepo_Articles(t *testing.T) {
	t.Run("Create a user", func(t *testing.T) {
		userFoo.ID = 17
		err := userRepo.CreateUser(userFoo)
		assert.NoError(t, err)
	})
	t.Run("create article", func(t *testing.T) {
		article := &models.Article{
			Title:       "title",
			Slug:        "simple-slug",
			Description: null.NewString("description", true),
			Body:        null.NewString("body", true),
			AuthorID:    null.NewUint64(uint64(userFoo.ID), true),
		}
		err := articleRepo.CreateArticle(article)
		assert.NoError(t, err)
	})
	t.Run("find article by slug", func(t *testing.T) {
		article, err := articleRepo.FindBySlug("simple-slug")
		assert.NoError(t, err)
		assert.Equal(t, article.Title, "title")
	})
	t.Run("find article by not exists slug", func(t *testing.T) {
		article, err := articleRepo.FindBySlug("simple-slug-test")
		assert.Error(t, sql.ErrNoRows, err)
		assert.Nil(t, article)
	})
	t.Run("find article by userID and slug", func(t *testing.T) {
		article, err := articleRepo.FindArticleByUserIDAndSlug(uint(userFoo.ID), "simple-slug")
		assert.NoError(t, err)
		assert.Equal(t, article.Title, "title")
	})
	t.Run("find article by userID and not exists slug", func(t *testing.T) {
		article, err := articleRepo.FindArticleByUserIDAndSlug(uint(userFoo.ID), "simple-slug-test")
		assert.Error(t, sql.ErrNoRows, err)
		assert.Nil(t, article)
	})
	t.Run("find article by userID and not exists userID", func(t *testing.T) {
		article, err := articleRepo.FindArticleByUserIDAndSlug(uint(userFoo.ID+1), "simple-slug")
		assert.Error(t, sql.ErrNoRows, err)
		assert.Nil(t, article)
	})
}
