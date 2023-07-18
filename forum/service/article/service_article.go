package article

import (
	"schema/entity"
	"sort"

	"forum/repository"

	"github.com/rs/zerolog/log"
)

type Service struct {
	Repo     repository.IRepoArticle
	UserRepo repository.IRepoUser
}

func NewServiceArticle(r repository.IRepoArticle, u repository.IRepoUser) *Service {
	return &Service{
		Repo:     r,
		UserRepo: u,
	}
}

func (r *Service) CreateArticle(a *entity.Article) error {
	return r.Repo.CreateArticle(a)
}

func (r *Service) UpdateArticle(slug string, newArticle *entity.Article) error {
	as, err := r.Repo.FindArticleBySlug(slug)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleBySlug error")
		return err
	}
	if newArticle.Body.Valid {
		as.Body.String = newArticle.Body.String
	}
	if newArticle.Slug != "" {
		as.Slug = newArticle.Slug
	}
	if newArticle.Title != "" {
		as.Title = newArticle.Title
	}
	if newArticle.Description.Valid {
		as.Description = newArticle.Description
	}
	err = r.Repo.UpdateArticle(as)
	if err != nil {
		log.Error().Err(err).Msg("UpdateArticle error")
		return err
	}
	return nil
}

func (r *Service) DeleteArticle(slug string) error {
	a, err := r.Repo.FindArticleBySlug(slug)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleBySlug error")
		return err
	}
	err = r.Repo.DeleteArticle(a)
	if err != nil {
		log.Error().Err(err).Msg("DeleteArticle error")
		return err
	}
	return nil
}

func (r *Service) FindArticle(slug string) (*entity.Article, *entity.User, []*entity.Tag, error) {
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

func (r *Service) FindArticleByAuthor(userName string, offset, limit int) ([]*entity.Article, int64, error) {
	u, err := r.UserRepo.FindUserByUserName(userName)
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

func (r *Service) FindArticles(tag, author string, offset, limit int) ([]*entity.Article, int64, error) {
	user, err := r.UserRepo.FindUserByUserName(author)
	if err != nil {
		log.Error().Err(err).Msg("FindByUserName error")
		return nil, 0, err
	}
	if tag != "" {
		a, n, err := r.Repo.ListArticlesByTag(tag, offset, limit)
		if err != nil {
			log.Error().Err(err).Msg("FindArticlesByTag error")
			return nil, 0, err
		}
		return a, n, nil
	} else if author != "" {
		a, n, err := r.Repo.ListArticlesByAuthor(user, offset, limit)
		if err != nil {
			log.Error().Err(err).Msg("FindArticleByAuthor error")
			return nil, 0, err
		}
		return a, n, nil
	} else {
		a, n, err := r.Repo.FindArticles(offset, limit)
		if err != nil {
			log.Error().Err(err).Msg("FindArticleByID error")
			return nil, 0, err
		}
		return a, n, nil
	}
}

func (r *Service) FindCommentsBySlug(slug string, offset, limit int) ([]*entity.Comment, error) {
	a, err := r.Repo.FindArticleBySlug(slug)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleBySlug error")
		return nil, err
	}
	c, err := r.Repo.FindCommentsByArticle(a, offset, limit)
	if err != nil {
		log.Error().Err(err).Msg("FindCommentsBySlug error")
		return nil, err
	}
	return c, nil
}

func (r *Service) FindAuthorBySlug(slug string) (*entity.User, error) {
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

func (r *Service) AddCommentToArticle(slug string, cm *entity.Comment) error {
	a, err := r.Repo.FindArticleBySlug(slug)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleBySlug error")
		return err
	}
	err = r.Repo.AddComment(a, cm)
	if err != nil {
		log.Error().Err(err).Msg("AddComment error")
		return err
	}
	return nil
}

func (r *Service) DeleteCommentFromArticle(slug string, commentId uint64) error {
	a, err := r.Repo.FindArticleBySlug(slug)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleBySlug error")
		return err
	}
	c, err := r.Repo.FindCommentByID(commentId)
	if err != nil {
		log.Error().Err(err).Msg("FindCommentByID error")
		return err
	}
	err = r.Repo.DeleteCommentByArticle(a, c)
	if err != nil {
		log.Error().Err(err).Msg("DeleteCommentByArticle error")
		return err
	}
	return nil
}

func (r *Service) AddFavoriteArticleBySlug(slug string, uid uint) error {
	a, u, err := r.FindArticleAndUserBySlugAndUserID(slug, uid)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleAndUserBySlugAndUserID error")
		return err
	}
	err = r.Repo.AddFavoriteArticle(a, u)
	if err != nil {
		log.Error().Err(err).Msg("AddFavoriteArticle error")
		return err
	}
	return nil
}

func (r *Service) RemoveFavoriteArticleBySlug(slug string, uid uint) error {
	a, u, err := r.FindArticleAndUserBySlugAndUserID(slug, uid)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleAndUserBySlugAndUserID error")
		return err
	}
	err = r.Repo.RemoveFavorite(a, u)
	if err != nil {
		log.Error().Err(err).Msg("RemoveFavorite error")
		return err
	}
	return nil
}

func (r *Service) FindArticleAndUserBySlugAndUserID(slug string, uid uint) (*entity.Article, *entity.User, error) {
	a, err := r.Repo.FindArticleBySlug(slug)
	if err != nil {
		log.Error().Err(err).Msg("FindArticleBySlug error")
		return nil, nil, err
	}
	u, err := r.UserRepo.FindUserByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindUserByID error")
		return nil, nil, err
	}
	return a, u, nil
}

func (r *Service) AddTagToArticle(slug string, tagStr []string) error {
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

func (r *Service) GetAllTags() ([]*entity.Tag, error) {
	t, err := r.Repo.ListTags()
	if err != nil {
		log.Error().Err(err).Msg("ListTags error")
		return nil, err
	}
	return t, nil
}
