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
	Repo     repository.Article
	UserRepo repository.User
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
	err = r.AddTagToArticle(a.Slug, tagStr)
	if err != nil {
		log.Error().Err(err).Msg("AddTagToArticle error")
		return err
	}
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
	u, err := r.Repo.FindAuthorByArticle(a)
	if err != nil {
		log.Error().Err(err).Msg("FindAuthorByArticle error")
		return nil, nil, nil, err
	}
	t, err := r.Repo.FindTagsByArticle(a)
	if err != nil {
		log.Error().Err(err).Msg("FindTagsByArticle error")
		return nil, nil, nil, err
	}
	return a, u, t, nil
}

func (r *RequestArticle) FindArticleByAuthor(userName string, offset, limit int) ([]*entity.Article, int64, error) {
	u, err := r.UserRepo.FindByUserName(userName)
	if err != nil {
		log.Error().Err(err).Msg("FindByUserName error")
		return nil, 0, err
	}
	a, n, err := r.Repo.ListArticlesByAuthor(u, offset, limit)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleByID error")
		return nil, 0, err
	}
	return a, n, nil
}

func (r *RequestArticle) FindArticles(tag, author string, offset, limit int) ([]*entity.Article, int64, error) {
	user, err := r.UserRepo.FindByUserName(author)
	if err != nil {
		log.Error().Err(err).Msg("FindByUserName error")
		return nil, 0, err
	}
	if tag != "" {
		a, n, err := r.Repo.ListArticlesByTag(tag, offset, limit)
		if err != nil {
			log.Error().Err(err).Msg("FindArticleByID error")
			return nil, 0, err
		}
		return a, n, nil
	} else if author != "" {
		a, n, err := r.Repo.ListArticlesByAuthor(user, offset, limit)
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
	a, err := r.Repo.FindArticleBySlug(slug)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleBySlug error")
		return nil, err
	}
	c, err := r.Repo.FindCommentsByArticle(a, offset, limit)
	if err != nil {
		log.Error().Err(err).Msg("ListCommentsBySlug error")
		return nil, err
	}
	return c, nil
}

func (r *RequestArticle) FindAuthorBySlug(slug string) (*entity.User, error) {
	a, err := r.Repo.FindArticleBySlug(slug)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleBySlug error")
		return nil, err
	}
	u, err := r.Repo.FindAuthorByArticle(a)
	if err != nil {
		log.Error().Err(err).Msg("FindAuthorByArticle error")
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
	err := r.Repo.DeleteCommentByCommentID(commentId)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleBySlug error")
		return err
	}
	return nil
}

func (r *RequestArticle) AddFavoriteArticleBySlug(slug string, uid uint) error {
	a, u, err := r.FindUserAndArticle(slug, uid)
	if err != nil {
		log.Error().Err(err).Msg("FindUserAndArticle error")
		return err
	}
	err = r.Repo.AddFavorite(a, u)
	if err != nil {
		log.Error().Err(err).Msg("AddFavorite error")
		return err
	}
	return nil
}

func (r *RequestArticle) FindUserAndArticle(slug string, uid uint) (*entity.Article, *entity.User, error) {
	a, err := r.Repo.FindArticleBySlug(slug)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleBySlug error")
		return nil, nil, err
	}
	u, err := r.UserRepo.FindByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindByID error")
		return nil, nil, err
	}
	return a, u, nil
}

func (r *RequestArticle) RemoveFavoriteArticleBySlug(slug string, uid uint) error {
	a, err := r.Repo.FindArticleBySlug(slug)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleBySlug error")
		return err
	}
	u, err := r.UserRepo.FindByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindByID error")
		return err
	}
	err = r.Repo.RemoveFavorite(a, u)
	if err != nil {
		log.Error().Err(err).Msg("RemoveFavorite error")
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
	err = r.Repo.AddTagsToArticle(a, tag)
	if err != nil {
		log.Error().Err(err).Msg("AddTagToArticle error")
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
