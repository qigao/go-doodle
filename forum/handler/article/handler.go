package article

import (
	"forum/entity"
	"forum/model"
	"forum/repository"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"
)

type Handler struct {
	article repository.Article
}

func NewArticleHandler(as repository.Article) *Handler {
	return &Handler{
		article: as,
	}
}

func bindJson(c echo.Context, doc interface{}) error {
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

func populateSingleArticle(s *model.SimpleArticle) (*entity.Article, []string) {
	var a *entity.Article
	a.Title = s.Title
	a.Slug = s.Slug
	a.Description = null.StringFrom(s.Description)
	a.Body = null.StringFrom(s.Body)
	return a, s.TagList
}
