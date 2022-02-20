package article

import (
	"forum/entity"
	"forum/model"
	"forum/repository"
	"sort"

	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"
)

type RequestArticle struct {
	Repo repository.Article
}

func (r *RequestArticle) Bind(c echo.Context, doc interface{}) error {
	if err := c.Bind(doc); err != nil {
		log.Error().Err(err).Msg("Bind error")
		return err
	}
	if err := c.Validate(doc); err != nil {
		log.Error().Err(err).Msg("Validate error")
		return err
	}
	return nil
}

func (r *RequestArticle) Populate(s *model.SingleArticle) (*entity.Article, []string) {
	var a *entity.Article
	a.Title = s.Title
	a.Slug = slug.Make(a.Title)
	a.Description = null.StringFrom(s.Description)
	a.Body = null.StringFrom(s.Body)
	return a, s.Tags
}

func (r *RequestArticle) InsertArticleWithTags(a *entity.Article, tagStr []string) error {
	err := r.Repo.CreateArticle(a)
	if err != nil {
		log.Error().Err(err).Msg("InsertArticle error")
		return err
	}
	err = r.AddTagToArticle(a, tagStr)
	return nil
}

func (r *RequestArticle) UpdateArticle(uid uint, slug string) error {
	a, err := r.Repo.FindArticleByAuthorIDAndSlug(uint64(uid), slug)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleBySlug error")
		return err
	}
	err = r.Repo.UpdateArticle(a)
	if err != nil {
		return err
	}
	return nil
}

func (r *RequestArticle) DeleteArticle(uid uint, slug string) error {
	a, err := r.Repo.FindArticleByAuthorIDAndSlug(uint64(uid), slug)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleBySlug error")
		return err
	}
	err = r.Repo.DeleteArticle(a)
	if err != nil {
		return err
	}
	return nil
}

func (r *RequestArticle) FindArticle(slug string) (*entity.Article, *entity.User, []*entity.Tag, error) {
	a, err := r.Repo.FindArticleBySlug(slug)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleBySlug error")
		return nil, nil, nil, err
	}
	u, err := r.Repo.FindAuthorBySlug(slug)
	if err != nil {
		log.Error().Err(err).Msg("FindAuthorBySlug error")
		return nil, nil, nil, err
	}
	t, err := r.Repo.FindTagsBySlug(slug)
	if err != nil {
		log.Error().Err(err).Msg("FindTagsBySlug error")
		return nil, nil, nil, err
	}
	return a, u, t, nil
}

func (r *RequestArticle) FindArticleByAuthor(userName string, offset, limit int) ([]*entity.Article, int64, error) {
	a, n, err := r.Repo.ListArticlesByAuthor(userName, offset, limit)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleByID error")
		return nil, 0, err
	}
	return a, n, nil
}

func (r *RequestArticle) FindArticles(tag, author string, offset, limit int) ([]*entity.Article, int64, error) {

	if tag != "" {
		a, n, err := r.Repo.ListArticlesByTag(tag, offset, limit)
		if err != nil {
			log.Error().Err(err).Msg("FindArticleByID error")
			return nil, 0, err
		}
		return a, n, nil
	} else if author != "" {
		a, n, err := r.Repo.ListArticlesByAuthor(author, offset, limit)
		if err != nil {
			log.Error().Err(err).Msg("FindArticleByID error")
			return nil, 0, err
		}
		return a, n, nil
	} else {
		a, n, err := r.Repo.ListArticles(offset, limit)
		if err != nil {
			log.Error().Err(err).Msg("FindArticleByID error")
			return nil, 0, err
		}
		return a, n, nil
	}
}

func (r *RequestArticle) FindCommentsBySlug(slug string, offset, limit int) ([]*entity.Comment, error) {
	c, err := r.Repo.FindCommentsBySlug(slug, offset, limit)
	if err != nil {
		log.Error().Err(err).Msg("ListCommentsBySlug error")
		return nil, err
	}
	return c, nil
}

func (r *RequestArticle) FindAuthorBySlug(slug string) (*entity.User, error) {
	u, err := r.Repo.FindAuthorBySlug(slug)
	if err != nil {
		log.Error().Err(err).Msg("FindAuthorBySlug error")
		return nil, err
	}
	return u, nil
}

func (r *RequestArticle) AddCommentToArticle(slug string, uid uint64, cm *entity.Comment) error {
	a, err := r.Repo.FindArticleByAuthorIDAndSlug(uid, slug)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleBySlug error")
		return err
	}
	err = r.Repo.AddComment(a, cm)
	if err != nil {
		return err
	}
	return nil
}

func (r *RequestArticle) DeleteCommentBySlugAndCommentID(slug string, commentId uint64) error {
	err := r.Repo.DeleteCommentBySlugAndCommentID(slug, commentId)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleBySlug error")
		return err
	}
	return nil
}

func (r *RequestArticle) AddFavoriteArticleBySlug(slug string, uid uint) error {
	a, err := r.Repo.FindArticleBySlug(slug)
	if err != nil {
		return err
	}
	err = r.Repo.AddFavorite(a, uid)
	if err != nil {
		return err
	}
	return nil
}

func (r *RequestArticle) RemoveFavoriteArticleBySlug(slug string, uid uint) error {
	a, err := r.Repo.FindArticleBySlug(slug)
	if err != nil {
		return err
	}
	err = r.Repo.RemoveFavorite(a, uid)
	if err != nil {
		return err
	}
	return nil
}

func (r *RequestArticle) AddTagToArticle(slug string, tagStr []string) error {
	a, err := r.Repo.FindArticleBySlug(slug)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleBySlug error")
		return err
	}
	t, err := r.Repo.ListTags()
	if err != nil {
		log.Error().Err(err).Msg("ListTags error")
		return err
	}
	sort.Strings(tagStr)
	var tag []*entity.Tag
	for _, v := range t {
		if contains(tagStr, v.Tag.String) {
			tag = append(tag, v)
		}
	}
	err = r.Repo.AddTags(a, tag)
	if err != nil {
		log.Error().Err(err).Msg("AddTag error")
		return err
	}
	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
