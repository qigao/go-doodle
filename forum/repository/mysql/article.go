package mysql

import (
	"context"
	"database/sql"
	"forum/entity"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ArticleRepo struct {
	Db *sql.DB
}

func (a *ArticleRepo) FindBySlug(s string) (*entity.Article, error) {
	article, err := entity.Articles(entity.ArticleWhere.Slug.EQ(s)).One(context.Background(), a.Db)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a *ArticleRepo) FindArticleByUserIDAndSlug(userID uint64, slug string) (*entity.Article, error) {
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

func (a *ArticleRepo) Create(article *entity.Article) error {
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

//Update  update article
func (a *ArticleRepo) Update(article *entity.Article) error {
	ctx := context.Background()
	tx, err := a.Db.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	row, err := article.Update(ctx, a.Db, boil.Infer())
	if err != nil {
		log.Error().Err(err).Msg("failed to update article")
		return err
	}
	tx.Commit()
	if row == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (a *ArticleRepo) Delete(article *entity.Article) error {
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
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

//List all the articles with pagination
func (a *ArticleRepo) List(offset, limit int) ([]*entity.Article, int64, error) {
	articles, err := entity.Articles(Limit(limit), Offset(offset)).All(context.Background(), a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to list articles")
		return nil, 0, err
	}
	return articles, int64(len(articles)), nil
}

func (a *ArticleRepo) ListByTag(tagStr string, offset, limit int) ([]*entity.Article, int64, error) {
	criteriaTags := entity.TagWhere.Tag.EQ(null.NewString(tagStr, true))
	ctx := context.Background()
	tag, err := entity.Tags(criteriaTags).One(ctx, a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to find tag")
		return nil, 0, err
	}
	articles, err := tag.Articles(Limit(limit), Offset(offset)).All(ctx, a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to list articles by tag")
		return nil, 0, err
	}

	return articles, int64(len(articles)), nil
}

func (a *ArticleRepo) ListByAuthor(username string, offset, limit int) ([]*entity.Article, int64, error) {
	user, err := entity.Users(entity.UserWhere.Username.EQ(username)).One(context.Background(), a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to find user")
		return nil, 0, err
	}
	articles, err := user.AuthorArticles(Limit(limit), Offset(offset)).All(context.Background(), a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to get articles")
		return nil, 0, err
	}
	return articles, int64(len(articles)), nil
}

func (a *ArticleRepo) ListFeed(userID uint, offset, limit int) ([]*entity.Article, int64, error) {
	//TODO implement me
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

func (a *ArticleRepo) FindCommentsBySlug(s string) ([]*entity.Comment, error) {
	article, err := a.FindBySlug(s)
	if err != nil {
		log.Error().Err(err).Msg("failed to find article")
		return nil, err
	}
	comments, err := article.Comments().All(context.Background(), a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to find comments")
		return nil, err
	}
	return comments, nil
}

func (a *ArticleRepo) FindCommentsByArticleID(articleId uint) ([]*entity.Comment, error) {
	ctx := context.Background()
	articles, err := entity.Articles(entity.ArticleWhere.ID.EQ(uint64(articleId))).All(ctx, a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to find comment")
		return nil, err
	}
	for _, article := range articles {
		comments, err := article.Comments().All(ctx, a.Db)
		if err != nil {
			log.Error().Err(err).Msg("failed to find comments")
			return nil, err
		}
		return comments, nil
	}
	return nil, nil
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
	rows, err := comment.Delete(ctx, a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to delete comment")
		return err
	}
	log.Info().Msgf("deleted %d rows", rows)
	tx.Commit()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (a *ArticleRepo) AddFavorite(article *entity.Article, favoritesUserID uint) error {
	ctx := context.Background()
	tx, err := a.Db.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	user, err := entity.Users(entity.UserWhere.ID.EQ(uint64(favoritesUserID))).One(ctx, a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to find user")
		return err
	}
	err = article.AddUsers(ctx, a.Db, false, user)
	if err != nil {
		log.Error().Err(err).Msg("failed to add favorite")
		return err
	}
	tx.Commit()
	return nil
}

func (a *ArticleRepo) RemoveFavorite(article *entity.Article, favoritesUserID uint) error {
	ctx := context.Background()
	tx, err := a.Db.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	user, err := entity.Users(entity.UserWhere.ID.EQ(uint64(favoritesUserID))).One(ctx, a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to find user")
		return err
	}
	err = article.RemoveUsers(ctx, a.Db, user)
	if err != nil {
		log.Error().Err(err).Msg("failed to remove favorite")
		return err
	}
	tx.Commit()
	return nil
}

func (a *ArticleRepo) ListByWhoFavorited(username string, offset, limit int) ([]*entity.Article, int64, error) {
	user, err := entity.Users(Load(entity.UserRels.Articles), entity.UserWhere.Username.EQ(username)).One(context.Background(), a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to find user")
		return nil, 0, err
	}
	articles, err := user.Articles().All(context.Background(), a.Db)
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

func (a *ArticleRepo) AddTag(article *entity.Article, tag *entity.Tag) error {
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

func (a *ArticleRepo) RemoveTag(article *entity.Article, tag *entity.Tag) error {
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

func (a *ArticleRepo) FindTagsByArticleID(article *entity.Article) ([]*entity.Tag, error) {
	ctx := context.Background()
	tags, err := article.Tags().All(ctx, a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to find tags")
		return nil, err
	}
	return tags, nil
}

func (a *ArticleRepo) ListTags() ([]*entity.Tag, error) {
	tags, err := entity.Tags().All(context.Background(), a.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to find tags")
		return nil, err
	}
	return tags, nil
}

func NewArticleRepo(db *sql.DB) *ArticleRepo {
	return &ArticleRepo{Db: db}
}
