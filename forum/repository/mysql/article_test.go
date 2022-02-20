package mysql

import (
	"database/sql"
	"forum/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
)

func TestArticleRepo_Articles(t *testing.T) {
	t.Run("Create a user", func(t *testing.T) {
		userFoo.ID = 170
		err := userRepo.CreateUser(userFoo)
		assert.NoError(t, err)
	})
	t.Run("create article", func(t *testing.T) {
		article := &entity.Article{
			Title:       "title",
			Slug:        "simple-slug",
			Description: null.NewString("description", true),
			Body:        null.NewString("body", true),
			AuthorID:    null.NewUint64(userFoo.ID, true),
		}
		article.ID = 10
		err := articleRepo.CreateArticle(article)
		assert.NoError(t, err)
	})
	t.Run("find article by slug then update", func(t *testing.T) {
		article, err := articleRepo.FindArticleBySlug("simple-slug")
		assert.NoError(t, err)
		article.Description = null.NewString("new description", true)
		err = articleRepo.UpdateArticle(article)
		assert.NoError(t, err)
		assert.Equal(t, "title", article.Title)
		assert.Equal(t, "body", article.Body.String)
		assert.Equal(t, "new description", article.Description.String)
	})
	t.Run("find article by not exists slug", func(t *testing.T) {
		article, err := articleRepo.FindArticleBySlug("simple-slug-test")
		assert.Error(t, sql.ErrNoRows, err)
		assert.Nil(t, article)
	})
	t.Run("find article by userID and slug", func(t *testing.T) {
		article, err := articleRepo.FindArticleByAuthorIDAndSlug(userFoo.ID, "simple-slug")
		assert.NoError(t, err)
		assert.Equal(t, "title", article.Title)
	})
	t.Run("find article by userID and not existing slug", func(t *testing.T) {
		article, err := articleRepo.FindArticleByAuthorIDAndSlug(userFoo.ID, "simple-slug-test")
		assert.Error(t, sql.ErrNoRows, err)
		assert.Nil(t, article)
	})
	t.Run("find article by not existing userID", func(t *testing.T) {
		article, err := articleRepo.FindArticleByAuthorIDAndSlug(10, "simple-slug")
		assert.Error(t, sql.ErrNoRows, err)
		assert.Nil(t, article)
	})
	t.Run("create article return error", func(t *testing.T) {
		article := &entity.Article{
			Title:       "title",
			Slug:        "simple-slug",
			Description: null.NewString("description", true),
			Body:        null.NewString("body", true),
			AuthorID:    null.NewUint64(userFoo.ID, true),
		}
		article.ID = 10
		err := articleRepo.CreateArticle(article)
		assert.Error(t, err)
	})
	t.Run("list articles ", func(t *testing.T) {
		articles, num, err := articleRepo.ListArticles(0, 10)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(articles))
		assert.Equal(t, int64(1), num)
	})
	t.Run("List articles by author", func(t *testing.T) {
		articles, num, err := articleRepo.ListArticlesByAuthor(userFoo.Username, 0, 10)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(articles))
		assert.Equal(t, int64(1), num)
	})
}

func TestArticleRepo_Tags(t *testing.T) {
	t.Run("Create tag1", func(t *testing.T) {
		tag := &entity.Tag{
			ID:  1,
			Tag: null.StringFrom("tag1"),
		}
		err := articleRepo.CreateTag(tag)
		assert.NoError(t, err)
	})
	t.Run("Create tag2", func(t *testing.T) {
		tag := &entity.Tag{
			ID:  2,
			Tag: null.StringFrom("tag2"),
		}
		err := articleRepo.CreateTag(tag)
		assert.NoError(t, err)
	})
	t.Run("Create tag3", func(t *testing.T) {
		tag := &entity.Tag{
			ID:  3,
			Tag: null.StringFrom("tag3"),
		}
		err := articleRepo.CreateTag(tag)
		assert.NoError(t, err)
	})

	t.Run("update article's tag", func(t *testing.T) {
		article, err := articleRepo.FindArticleBySlug("simple-slug")
		require.NoError(t, err)
		tag := &entity.Tag{
			ID:  1,
			Tag: null.StringFrom("tag1"),
		}
		err = articleRepo.AddTag(article, tag)
		assert.NoError(t, err)
	})
	t.Run("list  article by tag", func(t *testing.T) {
		articles, num, err := articleRepo.ListArticlesByTag("tag1", 0, 10)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), num)
		assert.Equal(t, 1, len(articles))
	})
	t.Run("list  article when not tagged", func(t *testing.T) {
		articles, num, err := articleRepo.ListArticlesByTag("tag2", 0, 10)
		assert.Nil(t, err)
		assert.Equal(t, int64(0), num)
		assert.Equal(t, 0, len(articles))
	})
	t.Run("list tags return 3 tags", func(t *testing.T) {
		tags, err := articleRepo.ListTags()
		require.Nil(t, err)
		var tagArray []string
		for _, tag := range tags {
			tagArray = append(tagArray, tag.Tag.String)
		}
		assert.ElementsMatch(t, []string{"tag1", "tag2", "tag3"}, tagArray)
	})
	t.Run("list tags of the article return error", func(t *testing.T) {
		articles, num, err := articleRepo.ListArticlesByTag("test-tag-fake1", 0, 10)
		assert.Error(t, sql.ErrNoRows, err)
		assert.Equal(t, int64(0), num)
		assert.Equal(t, 0, len(articles))
	})
	t.Run("list tags", func(t *testing.T) {
		tags, err := articleRepo.ListTags()
		assert.Nil(t, err)
		var tagArray []string
		for _, tag := range tags {
			tagArray = append(tagArray, tag.Tag.String)
		}
		assert.ElementsMatch(t, []string{"tag1", "tag2", "tag3"}, tagArray)
	})


}
func TestArticleCreateArticleWithTags(t *testing.T) {
	t.Run("create article Foo", func(t *testing.T) {
		err := articleRepo.CreateArticle(articleFoo)
		require.NoError(t, err)
		tags := []*entity.Tag{
			{Tag: null.StringFrom("tag1"), ID: 1},
			{Tag: null.StringFrom("tag2"), ID: 2},
			{Tag: null.StringFrom("tag3"), ID: 3},
		}
		err = articleRepo.AddTags(articleFoo, tags)
		assert.NoError(t, err)
	})
	t.Run("find tags by slug", func(t *testing.T) {
		tags, err := articleRepo.FindTagsBySlug("foo-slug")
		assert.NoError(t, err)
		var tagArray []string
		for _, tag := range tags {
			tagArray = append(tagArray, tag.Tag.String)
		}
		assert.ElementsMatch(t, []string{"tag1", "tag2", "tag3"}, tagArray)
	})
}
func TestArticleRepo_Comments(t *testing.T) {
	t.Run("Add comment1", func(t *testing.T) {
		article, err := articleRepo.FindArticleBySlug("simple-slug")
		require.NoError(t, err)
		comment := &entity.Comment{
			Body: null.StringFrom("test1"),
		}
		err = articleRepo.AddComment(article, comment)
		assert.NoError(t, err)
	})
	t.Run("Add comment2", func(t *testing.T) {
		article, err := articleRepo.FindArticleBySlug("simple-slug")
		require.NoError(t, err)
		comment := &entity.Comment{
			Body: null.StringFrom("test2"),
		}
		err = articleRepo.AddComment(article, comment)
		assert.NoError(t, err)
	})
	t.Run("find comments by slug", func(t *testing.T) {
		comments, err := articleRepo.FindCommentsBySlug("simple-slug", 0, 10)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(comments))
		var commArr []string
		for _, comment := range comments {
			commArr = append(commArr, comment.Body.String)
		}
		assert.ElementsMatch(t, []string{"test1", "test2"}, commArr)
	})
	t.Run("find comment by comment id", func(t *testing.T) {
		comment, err := articleRepo.FindCommentByID(1)
		require.NoError(t, err)
		err = articleRepo.DeleteComment(comment)
		assert.NoError(t, err)
	})
	t.Run("find coment by non exists comment id", func(t *testing.T) {
		comment, err := articleRepo.FindCommentByID(1)
		assert.Error(t, sql.ErrNoRows, err)
		assert.Nil(t, comment)
	})
}

func TestArticleRepo_Favorites(t *testing.T) {
	t.Run("Create a user", func(t *testing.T) {
		userFoo.ID = 171
		userFoo.Username = "foo2"
		err := userRepo.CreateUser(userFoo)
		assert.NoError(t, err)
	})
	t.Run("Add favorite", func(t *testing.T) {
		article, err := articleRepo.FindArticleBySlug("simple-slug")
		require.NoError(t, err)
		err = articleRepo.AddFavorite(article, 171)
		assert.NoError(t, err)
	})
	t.Run("find articles by favorited", func(t *testing.T) {
		articles, num, err := articleRepo.ListArticlesByWhoFavorited("foo2", 0, 10)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), num)
		assert.Equal(t, "simple-slug", articles[0].Slug)
	})

}
