package mysql

import (
	"context"
	"database/sql"
	models "forum/entities"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ArticleRepo struct {
	Db *sql.DB
}

func (a *ArticleRepo) FindBySlug(s string) (*models.Article, error) {
	article, err := models.Articles(models.ArticleWhere.Slug.EQ(s)).One(context.Background(), a.Db)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a *ArticleRepo) FindArticleByUserIDAndSlug(userID uint, slug string) (*models.Article, error) {
	criteriaUserid := models.ArticleWhere.AuthorID.EQ(null.NewUint64(uint64(userID), true))
	criteriaSlug := models.ArticleWhere.Slug.EQ(slug)

	article, err := models.Articles(
		criteriaSlug,
		criteriaUserid).One(context.Background(), a.Db)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a *ArticleRepo) CreateArticle(article *models.Article) error {
	err := article.Insert(context.Background(), a.Db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (a *ArticleRepo) UpdateArticle(article *models.Article, strings []string) error {
	tx, err := a.Db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	rows, err := article.Update(context.Background(), a.Db, boil.Infer())
	if err != nil {
		log.Error().Err(err).Msg("failed to update article")
		return err
	}
	log.Info().Msgf("updated %d rows", rows)
	tx.Commit()
	return nil
}

func (a *ArticleRepo) DeleteArticle(article *models.Article) error {
	tx, err := a.Db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	rows, err := article.Delete(context.Background(), a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to delete article")
		return err
	}
	log.Info().Msgf("deleted %d rows", rows)
	tx.Commit()
	return nil
}

func (a *ArticleRepo) ListArticles(offset, limit int) ([]*models.Article, int64, error) {
	articles, err := models.Articles().All(context.Background(), a.Db)
	if err != nil {
		return nil, 0, err
	}
	return articles, int64(len(articles)), nil
}

func (a *ArticleRepo) ListByTag(tag string, offset, limit int) ([]models.Article, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleRepo) ListByAuthor(username string, offset, limit int) ([]models.Article, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleRepo) ListByWhoFavorited(username string, offset, limit int) ([]models.Article, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleRepo) ListFeed(userID uint, offset, limit int) ([]models.Article, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleRepo) AddComment(article *models.Article, comment *models.Comment) error {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleRepo) FindCommentsBySlug(s string) ([]models.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleRepo) FindCommentByID(u uint) (*models.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleRepo) DeleteComment(comment *models.Comment) error {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleRepo) AddFavorite(article *models.Article, u uint) error {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleRepo) RemoveFavorite(article *models.Article, u uint) error {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleRepo) ListTags() ([]models.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func NewArticleRepo(db *sql.DB) *ArticleRepo {
	return &ArticleRepo{Db: db}
}
