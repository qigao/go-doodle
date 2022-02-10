package repository

import (
	"fmt"
	"gforum/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArticleRepository_CreateAndUpdateArticle(t *testing.T) {
	t.Run("create a new article", func(t *testing.T) {
		err := articleRepo.Create(articleFoo)
		require.NoError(t, err)
	})
	t.Run("get article by id", func(t *testing.T) {
		article, err := articleRepo.GetByID(articleFoo.ID)
		require.NoError(t, err)
		require.Equal(t, articleFoo.ID, article.ID)
		require.Equal(t, articleFoo.Title, article.Title)
		require.Equal(t, articleFoo.Body, article.Body)
	})
	t.Run("get article by id return error", func(t *testing.T) {
		gormDB, mock := mockMysql()
		repo := repository.NewArticleRepository(gormDB)
		mock.ExpectQuery("SELECT").WithArgs(articleFoo.ID).WillReturnError(fmt.Errorf("error"))
		result, err := repo.GetByID(articleFoo.ID)
		require.Error(t, err)
		require.Nil(t, result)
	})
	t.Run("update article", func(t *testing.T) {
		articleFoo.Title = "foo"
		articleFoo.Body = "bar"
		err := articleRepo.Update(articleFoo)
		require.NoError(t, err)
	})
	t.Run("get article by id", func(t *testing.T) {
		article, err := articleRepo.GetByID(articleFoo.ID)
		require.NoError(t, err)
		require.Equal(t, articleFoo.ID, article.ID)
		require.Equal(t, articleFoo.Title, article.Title)
		require.Equal(t, articleFoo.Body, article.Body)
	})

}

func TestArticleRepository_GetArticles(t *testing.T) {
	t.Run("get articles", func(t *testing.T) {
		articles, err := articleRepo.GetArticles("foo", "foo", userBar, 1, 10)
		require.NoError(t, err)
		require.Equal(t, 1, len(articles))
		require.Equal(t, articleFoo.ID, articles[0].ID)
		require.Equal(t, articleFoo.Title, articles[0].Title)
		require.Equal(t, articleFoo.Body, articles[0].Body)
	})
}
