package mysql

import (
	"context"
	"database/sql"
	"schema/entity"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ArticleRepo struct {
	Db *sql.DB
}

func NewArticleRepo(db *sql.DB) *ArticleRepo {
	return &ArticleRepo{Db: db}
}

func (a *ArticleRepo) FindArticleBySlug(s string) (*entity.Article, error) {
	article, err := entity.Articles(entity.ArticleWhere.Slug.EQ(s)).One(context.Background(), a.Db)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a *ArticleRepo) FindArticleByAuthorIDAndSlug(userID uint64, slug string) (*entity.Article, error) {
	criteriaUserid := entity.ArticleWhere.AuthorID.EQ(null.NewUint64(userID, true))
	criteriaSlug := entity.ArticleWhere.Slug.EQ(slug)

	article, err := entity.Articles(
		criteriaSlug,
		criteriaUserid).One(context.Background(), a.Db)
	if err != nil {
		log.Error().Err(err).Msg("error while finding article")
		return nil, err
	}
	return article, nil
}

func (a *ArticleRepo) CreateArticle(article *entity.Article) error {
	ctx := context.Background()
	tx, err := a.Db.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	err = article.Insert(ctx, a.Db, boil.Infer())
	if err != nil {
		log.Error().Err(err).Msg("failed to insert article")
		return err
	}
	tx.Commit()
	return nil
}

// UpdateArticle  update article
func (a *ArticleRepo) UpdateArticle(article *entity.Article) error {
	ctx := context.Background()
	tx, err := a.Db.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	_, err = article.Update(ctx, a.Db, boil.Infer())
	if err != nil {
		log.Error().Err(err).Msg("failed to update article")
		return err
	}
	tx.Commit()
	return nil
}

func (a *ArticleRepo) DeleteArticle(article *entity.Article) error {
	ctx := context.Background()
	tx, err := a.Db.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	_, err = article.Delete(ctx, a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to delete article")
		return err
	}
	tx.Commit()
	return nil
}

// FindArticles all the articles with pagination
func (a *ArticleRepo) FindArticles(offset, limit int) ([]*entity.Article, int64, error) {
	articles, err := entity.Articles(qm.Limit(limit), qm.Offset(offset)).All(context.Background(), a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to list articles")
		return nil, 0, err
	}
	return articles, int64(len(articles)), nil
}

func (a *ArticleRepo) ListArticlesByTag(tagStr string, offset, limit int) ([]*entity.Article, int64, error) {
	criteriaTags := entity.TagWhere.Tag.EQ(null.NewString(tagStr, true))
	ctx := context.Background()
	tag, err := entity.Tags(criteriaTags).One(ctx, a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to find tag")
		return nil, 0, err
	}
	articles, err := tag.Articles(qm.Limit(limit), qm.Offset(offset)).All(ctx, a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to list articles by tag")
		return nil, 0, err
	}
	return articles, int64(len(articles)), nil
}

func (a *ArticleRepo) ListArticlesByAuthor(user *entity.User, offset, limit int) ([]*entity.Article, int64, error) {
	articles, err := user.AuthorArticles(qm.Limit(limit), qm.Offset(offset)).All(context.Background(), a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to get articles")
		return nil, 0, err
	}
	return articles, int64(len(articles)), nil
}

func (a *ArticleRepo) FindAuthorByArticle(article *entity.Article) (*entity.User, error) {
	return article.Author().One(context.Background(), a.Db)
}

func (a *ArticleRepo) ListFeed(userID uint, offset, limit int) ([]*entity.Article, int64, error) {
	// TODO implement me
	panic("implement me")
}

func (a *ArticleRepo) AddComment(article *entity.Article, comment *entity.Comment) error {
	tx, err := a.Db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	ctx := context.Background()
	err = article.AddComments(ctx, a.Db, true, comment)
	if err != nil {
		log.Error().Err(err).Msg("failed to add comment")
		return err
	}
	tx.Commit()
	return nil
}

func (a *ArticleRepo) FindCommentsByArticle(article *entity.Article, offset int, limit int) ([]*entity.Comment, error) {
	return article.Comments(qm.Limit(limit), qm.Offset(offset)).All(context.Background(), a.Db)
}

func (a *ArticleRepo) FindCommentByID(commentID uint64) (*entity.Comment, error) {
	ctx := context.Background()
	comment, err := entity.Comments(entity.CommentWhere.ID.EQ(commentID)).One(ctx, a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to find comment")
		return nil, err
	}
	return comment, nil
}

func (a *ArticleRepo) DeleteComment(comment *entity.Comment) error {
	ctx := context.Background()
	tx, err := a.Db.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	_, err = comment.Delete(ctx, a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to delete comment")
		return err
	}
	tx.Commit()
	return nil
}

func (a *ArticleRepo) DeleteCommentByCommentID(commentID uint64) error {
	ctx := context.Background()
	tx, err := a.Db.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	_, err = entity.Comments(
		entity.CommentWhere.ID.EQ(commentID)).DeleteAll(ctx, a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to delete comment")
		return err
	}
	tx.Commit()
	return nil
}

func (a *ArticleRepo) DeleteCommentByArticle(article *entity.Article, comment *entity.Comment) error {
	tx, err := a.Db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	ctx := context.Background()
	err = article.RemoveComments(ctx, a.Db, comment)
	if err != nil {
		log.Error().Err(err).Msg("failed to add comment")
		return err
	}
	tx.Commit()
	return nil
}

func (a *ArticleRepo) AddFavoriteArticle(article *entity.Article, user *entity.User) error {
	ctx := context.Background()
	tx, err := a.Db.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	err = article.AddUsers(ctx, a.Db, false, user)
	if err != nil {
		log.Error().Err(err).Msg("failed to add favorite")
		return err
	}
	tx.Commit()
	return nil
}

func (a *ArticleRepo) RemoveFavorite(article *entity.Article, user *entity.User) error {
	ctx := context.Background()
	tx, err := a.Db.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	err = article.RemoveUsers(ctx, a.Db, user)
	if err != nil {
		log.Error().Err(err).Msg("failed to remove favorite")
		return err
	}
	tx.Commit()
	return nil
}

func (a *ArticleRepo) FindFavoriteArticlesByUser(user *entity.User, offset, limit int) ([]*entity.Article, int64, error) {
	articles, err := user.Articles(qm.Offset(offset), qm.Limit(limit)).All(context.Background(), a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to find articles")
		return nil, 0, err
	}
	return articles, int64(len(articles)), nil
}

func (a *ArticleRepo) CreateTag(tag *entity.Tag) error {
	ctx := context.Background()
	tx, err := a.Db.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	err = tag.Insert(ctx, a.Db, boil.Infer())
	if err != nil {
		log.Error().Err(err).Msg("failed to create tag")
		return err
	}
	tx.Commit()
	return nil
}

func (a *ArticleRepo) AddTagToArticle(article *entity.Article, tag *entity.Tag) error {
	ctx := context.Background()
	tx, err := a.Db.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	err = article.AddTags(ctx, a.Db, false, tag)
	if err != nil {
		log.Error().Err(err).Msg("failed to add tag")
		return err
	}
	tx.Commit()
	return nil
}

func (a *ArticleRepo) AddTagsToArticle(article *entity.Article, tag []*entity.Tag) error {
	ctx := context.Background()
	tx, err := a.Db.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	err = article.AddTags(ctx, a.Db, false, tag...)
	if err != nil {
		log.Error().Err(err).Msg("failed to add tag")
		return err
	}
	tx.Commit()
	return nil
}

func (a *ArticleRepo) RemoveTagFromArticle(article *entity.Article, tag *entity.Tag) error {
	tx, err := a.Db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	err = article.RemoveTags(context.Background(), a.Db, tag)
	if err != nil {
		log.Error().Err(err).Msg("failed to remove tag")
		return err
	}
	tx.Commit()
	return nil
}

func (a *ArticleRepo) RemoveTagsFromArticle(article *entity.Article, tags []*entity.Tag) error {
	tx, err := a.Db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	err = article.RemoveTags(context.Background(), a.Db, tags...)
	if err != nil {
		log.Error().Err(err).Msg("failed to remove tag")
		return err
	}
	tx.Commit()
	return nil
}

func (a *ArticleRepo) FindTagsByArticle(article *entity.Article) ([]*entity.Tag, error) {
	return article.Tags().All(context.Background(), a.Db)
}

func (a *ArticleRepo) ListTags() ([]*entity.Tag, error) {
	tags, err := entity.Tags().All(context.Background(), a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to find tags")
		return nil, err
	}
	return tags, nil
}
