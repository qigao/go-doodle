package mysql

import (
	"fmt"
	"regexp"
	"schema/entity"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
)

func TestArticle_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when create article failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewArticleRepo(db)
		err = repo.CreateArticle(articleFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when create article success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewArticleRepo(db)
		err = repo.CreateArticle(articleFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		err = repo.CreateArticle(articleFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when update article failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewArticleRepo(db)
		err = repo.UpdateArticle(articleFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when update article success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewArticleRepo(db)
		err = repo.UpdateArticle(articleFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		err = repo.UpdateArticle(articleFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArtcile_DeleteArticle(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when delete article failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewArticleRepo(db)
		err = repo.DeleteArticle(articleFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when delete article success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewArticleRepo(db)
		err = repo.DeleteArticle(articleFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		err = repo.DeleteArticle(articleFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_AddComment(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when add comment failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewArticleRepo(db)
		err = repo.AddComment(articleFoo, commentFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when add comment success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewArticleRepo(db)
		err = repo.AddComment(articleFoo, commentFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		err = repo.AddComment(articleFoo, commentFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_DeleteComment(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when remove comment failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewArticleRepo(db)
		err = repo.DeleteComment(commentFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when remove comment success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewArticleRepo(db)
		err = repo.DeleteComment(commentFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		err = repo.DeleteComment(commentFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_DeleteCommentByArticle(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when delete comment by comment id success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("DELETE")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewArticleRepo(db)
		err := repo.DeleteCommentByCommentID(commentFoo.ID)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when delete comment by comment id failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("DELETE")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewArticleRepo(db)
		err := repo.DeleteCommentByCommentID(commentFoo.ID)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when delete comment by comment id transaction failed", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		err := repo.DeleteCommentByCommentID(commentFoo.ID)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_CreateTag(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when create tag failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewArticleRepo(db)
		err = repo.CreateTag(tagFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when create tag success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewArticleRepo(db)
		err = repo.CreateTag(tagFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		err = repo.CreateTag(tagFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_AddTag(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when add tag failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("insert into")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewArticleRepo(db)
		err = repo.AddTagToArticle(articleFoo, tagFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when add tag success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("insert into")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewArticleRepo(db)
		err = repo.AddTagToArticle(articleFoo, tagFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		err = repo.AddTagToArticle(articleFoo, tagFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("find tag by article", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "tag1"))
		repo := NewArticleRepo(db)
		tags, err := repo.FindTagsByArticle(articleFoo)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(tags))
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_AddTags(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when add tags failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("insert into")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewArticleRepo(db)
		err = repo.AddTagsToArticle(articleFoo, []*entity.Tag{tagFoo, tagBar})
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when add tags success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("insert into")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec(regexp.QuoteMeta("insert into")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewArticleRepo(db)
		err = repo.AddTagsToArticle(articleFoo, []*entity.Tag{tagFoo, tagBar})
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		err = repo.AddTagsToArticle(articleFoo, []*entity.Tag{tagFoo, tagBar})
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_RemoveTag(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when remove tag failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("delete from")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewArticleRepo(db)
		err = repo.RemoveTagFromArticle(articleFoo, tagFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when remove tag success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("delete from")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewArticleRepo(db)
		err = repo.RemoveTagFromArticle(articleFoo, tagFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		err = repo.RemoveTagFromArticle(articleFoo, tagFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_RemoveTags(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when remove tags failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("delete from")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewArticleRepo(db)
		err = repo.RemoveTagsFromArticle(articleFoo, []*entity.Tag{tagFoo, tagBar})
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when remove tags success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("delete from")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewArticleRepo(db)
		err = repo.RemoveTagsFromArticle(articleFoo, []*entity.Tag{tagFoo, tagBar})
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		err = repo.RemoveTagsFromArticle(articleFoo, []*entity.Tag{tagFoo, tagBar})
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_ListTags(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when list tags return OK", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo").AddRow(2, "bar"))
		repo := NewArticleRepo(db)
		tags, err := repo.ListTags()
		assert.NoError(t, err)
		assert.Equal(t, 2, len(tags))
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when list tags return error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		tags, err := repo.ListTags()
		assert.Errorf(t, err, "some error")
		assert.Equal(t, 0, len(tags))
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_AddFavorite(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when add favorite failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("insert into")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewArticleRepo(db)
		err = repo.AddFavoriteArticle(articleFoo, userFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when add favorite success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("insert into")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewArticleRepo(db)
		err = repo.AddFavoriteArticle(articleFoo, userFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		err = repo.AddFavoriteArticle(articleFoo, userFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_RemoveFavorite(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when remove favorite failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("delete from")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewArticleRepo(db)
		err = repo.RemoveFavorite(articleFoo, userFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when remove favorite success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("delete from")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewArticleRepo(db)
		err = repo.RemoveFavorite(articleFoo, userFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		err = repo.RemoveFavorite(articleFoo, userFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_FindArticleBySlug(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when find article by slug success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "slug", "body", "description", "created_at", "updated_at", "deleted_at", "author_id"}).
			AddRow(articleFoo.ID, articleFoo.Title, articleFoo.Slug, articleFoo.Body, articleFoo.Description, articleFoo.CreatedAt, articleFoo.UpdatedAt, articleFoo.DeletedAt, 1)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(rows)
		repo := NewArticleRepo(db)
		article, err := repo.FindArticleBySlug(articleFoo.Slug)
		assert.NoError(t, err)
		assert.Equal(t, articleFoo.ID, article.ID)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when find article by slug failed", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		_, err = repo.FindArticleBySlug(articleFoo.Slug)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_FindArticleByAuthorIDAndSlug(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when find article by author id and slug success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "slug", "body", "description", "created_at", "updated_at", "deleted_at", "author_id"}).
			AddRow(articleFoo.ID, articleFoo.Title, articleFoo.Slug, articleFoo.Body, articleFoo.Description, articleFoo.CreatedAt, articleFoo.UpdatedAt, articleFoo.DeletedAt, 1)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(rows)
		repo := NewArticleRepo(db)
		article, err := repo.FindArticleByAuthorIDAndSlug(1, articleFoo.Slug)
		assert.NoError(t, err)
		assert.Equal(t, articleFoo.ID, article.ID)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when find article by author id and slug failed", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		_, err := repo.FindArticleByAuthorIDAndSlug(1, articleFoo.Slug)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_ListArticles(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when list articles success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "slug", "body", "description", "created_at", "updated_at", "deleted_at", "author_id"}).
			AddRow(articleFoo.ID, articleFoo.Title, articleFoo.Slug, articleFoo.Body, articleFoo.Description, articleFoo.CreatedAt, articleFoo.UpdatedAt, articleFoo.DeletedAt, 1).
			AddRow(articleBar.ID, articleBar.Title, articleBar.Slug, articleBar.Body, articleBar.Description, articleBar.CreatedAt, articleBar.UpdatedAt, articleBar.DeletedAt, 1)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(rows)
		repo := NewArticleRepo(db)
		_, n, err := repo.FindArticles(0, 1)
		assert.NoError(t, err)
		assert.Equal(t, int64(2), n)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when list articles failed", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		_, n, err := repo.FindArticles(0, 1)
		assert.Errorf(t, err, "some error")
		assert.Equal(t, int64(0), n)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_ListArticlesByTag(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when list articles by tag success", func(t *testing.T) {
		tagRows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "tag"}).
			AddRow(tagFoo.ID, tagFoo.CreatedAt, tagFoo.UpdatedAt, tagFoo.DeletedAt, tagFoo.Tag).
			AddRow(tagBar.ID, tagBar.CreatedAt, tagBar.UpdatedAt, tagBar.DeletedAt, tagBar.Tag)
		rows := sqlmock.NewRows([]string{"id", "title", "slug", "body", "description", "created_at", "updated_at", "deleted_at", "author_id"}).
			AddRow(articleFoo.ID, articleFoo.Title, articleFoo.Slug, articleFoo.Body, articleFoo.Description, articleFoo.CreatedAt, articleFoo.UpdatedAt, articleFoo.DeletedAt, 1).
			AddRow(articleBar.ID, articleBar.Title, articleBar.Slug, articleBar.Body, articleBar.Description, articleBar.CreatedAt, articleBar.UpdatedAt, articleBar.DeletedAt, 1)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(tagRows)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(rows)

		repo := NewArticleRepo(db)
		_, n, err := repo.ListArticlesByTag("tag", 0, 1)
		assert.NoError(t, err)
		assert.Equal(t, int64(2), n)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when list articles by tag find article failed", func(t *testing.T) {
		tagRows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "tag"}).
			AddRow(tagFoo.ID, tagFoo.CreatedAt, tagFoo.UpdatedAt, tagFoo.DeletedAt, tagFoo.Tag).
			AddRow(tagBar.ID, tagBar.CreatedAt, tagBar.UpdatedAt, tagBar.DeletedAt, tagBar.Tag)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(tagRows)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		_, n, err := repo.ListArticlesByTag("tag", 0, 1)
		assert.Errorf(t, err, "some error")
		assert.Equal(t, int64(0), n)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when list articles by tag find tag failed", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		_, n, err := repo.ListArticlesByTag("tag", 0, 1)
		assert.Errorf(t, err, "some error")
		assert.Equal(t, int64(0), n)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_ListArticlesByAuthor(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when list articles by author success", func(t *testing.T) {
		articleRows := sqlmock.NewRows([]string{"id", "title", "slug", "body", "description", "created_at", "updated_at", "deleted_at", "author_id"}).
			AddRow(articleFoo.ID, articleFoo.Title, articleFoo.Slug, articleFoo.Body, articleFoo.Description, articleFoo.CreatedAt, articleFoo.UpdatedAt, articleFoo.DeletedAt, 1).
			AddRow(articleBar.ID, articleBar.Title, articleBar.Slug, articleBar.Body, articleBar.Description, articleBar.CreatedAt, articleBar.UpdatedAt, articleBar.DeletedAt, 1)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(articleRows)
		repo := NewArticleRepo(db)
		_, n, err := repo.ListArticlesByAuthor(userFoo, 0, 1)
		assert.NoError(t, err)
		assert.Equal(t, int64(2), n)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when list articles by author find author  failed", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		_, n, err := repo.ListArticlesByAuthor(userFoo, 0, 1)
		assert.Errorf(t, err, "some error")
		assert.Equal(t, int64(0), n)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_FindAuthorBySlug(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when find author by slug success", func(t *testing.T) {
		users := sqlmock.NewRows([]string{"id", "username", "email", "password", "bio", "image", "created_at", "updated_at", "deleted_at", "following_id"}).
			AddRow(2, "foo", "foo@foo.com", "foo-password", "foo desc", "http://foo.com/foo.jpg", null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), null.TimeFrom(time.Now()), 2)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
			WillReturnRows(users)

		repo := NewArticleRepo(db)
		_, err := repo.FindAuthorByArticle(articleFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_FindCommentsByArticle(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when find comments  success", func(t *testing.T) {
		commentRows := sqlmock.NewRows([]string{"id", "body", "created_at", "updated_at", "deleted_at", "author_id", "article_id"}).
			AddRow(commentFoo.ID, commentFoo.Body, commentFoo.CreatedAt, commentFoo.UpdatedAt, commentFoo.DeletedAt, commentFoo.UserID, commentFoo.ArticleID)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(commentRows)
		repo := NewArticleRepo(db)
		_, err := repo.FindCommentsByArticle(articleFoo, 0, 1)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_FindCommentByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when find comment by id success", func(t *testing.T) {
		commentRows := sqlmock.NewRows([]string{"id", "body", "created_at", "updated_at", "deleted_at", "author_id", "article_id"}).
			AddRow(commentFoo.ID, commentFoo.Body, commentFoo.CreatedAt, commentFoo.UpdatedAt, commentFoo.DeletedAt, commentFoo.UserID, commentFoo.ArticleID)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(commentRows)
		repo := NewArticleRepo(db)
		_, err := repo.FindCommentByID(commentFoo.ID)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when find comment by id failed", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		_, err := repo.FindCommentByID(commentFoo.ID)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_DeleteCommentByCommentID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when delete comment by comment id success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("DELETE")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := NewArticleRepo(db)
		err := repo.DeleteCommentByCommentID(commentFoo.ID)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when delete comment by comment id failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("DELETE")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := NewArticleRepo(db)
		err := repo.DeleteCommentByCommentID(commentFoo.ID)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when delete comment by comment id transaction failed", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		err := repo.DeleteCommentByCommentID(commentFoo.ID)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_FindFavoriteArticlesByUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("when find favorite articles by user success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "slug", "body", "description", "created_at", "updated_at", "deleted_at", "author_id"}).
			AddRow(articleFoo.ID, articleFoo.Title, articleFoo.Slug, articleFoo.Body, articleFoo.Description, articleFoo.CreatedAt, articleFoo.UpdatedAt, articleFoo.DeletedAt, 1).
			AddRow(articleBar.ID, articleBar.Title, articleBar.Slug, articleBar.Body, articleBar.Description, articleBar.CreatedAt, articleBar.UpdatedAt, articleBar.DeletedAt, 1)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(rows)
		repo := NewArticleRepo(db)
		_, n, err := repo.FindFavoriteArticlesByUser(userFoo, 0, 1)
		assert.NoError(t, err)
		assert.Equal(t, int64(2), n)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("when find favorite articles by user failed", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnError(fmt.Errorf("some error"))
		repo := NewArticleRepo(db)
		_, _, err := repo.FindFavoriteArticlesByUser(userFoo, 0, 1)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
