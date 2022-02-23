package mysql

import (
	"fmt"
	"forum/entity"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestArticle_Create(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when create article failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := ArticleRepo{Db: db}
		err = repo.CreateArticle(articleFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when create article success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := ArticleRepo{Db: db}
		err = repo.CreateArticle(articleFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := ArticleRepo{Db: db}
		err = repo.CreateArticle(articleFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestArticle_Update(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when update article failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := ArticleRepo{Db: db}
		err = repo.UpdateArticle(articleFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when update article success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := ArticleRepo{Db: db}
		err = repo.UpdateArticle(articleFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := ArticleRepo{Db: db}
		err = repo.UpdateArticle(articleFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestArtcile_DeleteArticle(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when delete article failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := ArticleRepo{Db: db}
		err = repo.DeleteArticle(articleFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when delete article success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := ArticleRepo{Db: db}
		err = repo.DeleteArticle(articleFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := ArticleRepo{Db: db}
		err = repo.DeleteArticle(articleFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_AddComment(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when add comment failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := ArticleRepo{Db: db}
		err = repo.AddComment(articleFoo, commentFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when add comment success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := ArticleRepo{Db: db}
		err = repo.AddComment(articleFoo, commentFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := ArticleRepo{Db: db}
		err = repo.AddComment(articleFoo, commentFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_CreateTag(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when create tag failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := ArticleRepo{Db: db}
		err = repo.CreateTag(tagFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when create tag success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := ArticleRepo{Db: db}
		err = repo.CreateTag(tagFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := ArticleRepo{Db: db}
		err = repo.CreateTag(tagFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_AddTag(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when add tag failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("insert into")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := ArticleRepo{Db: db}
		err = repo.AddTagToArticle(articleFoo, tagFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when add tag success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("insert into")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := ArticleRepo{Db: db}
		err = repo.AddTagToArticle(articleFoo, tagFoo)
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := ArticleRepo{Db: db}
		err = repo.AddTagToArticle(articleFoo, tagFoo)
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestArticle_AddTags(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	t.Run("transaction rollback when add tags failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("insert into")).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()
		repo := ArticleRepo{Db: db}
		err = repo.AddTagsToArticle(articleFoo, []*entity.Tag{tagFoo, tagBar})
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction commit when add tags success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("insert")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		repo := ArticleRepo{Db: db}
		err = repo.AddTagsToArticle(articleFoo, []*entity.Tag{tagFoo, tagBar})
		assert.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("transaction begin with error", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))
		repo := ArticleRepo{Db: db}
		err = repo.AddTagsToArticle(articleFoo, []*entity.Tag{tagFoo, tagBar})
		assert.Errorf(t, err, "some error")
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
