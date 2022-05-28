package article

import (
	"schema/entity"

	"forum/model"
	"forum/service"

	"github.com/volatiletech/null/v8"
)

type Handler struct {
	Service service.IServiceArticle
}

func NewArticleHandler(as service.IServiceArticle) *Handler {
	return &Handler{
		Service: as,
	}
}

func populateSimpleArticle(s *model.SimpleArticle) *entity.Article {
	var a entity.Article
	a.Title = s.Title
	a.Slug = s.Slug
	a.Description = null.StringFrom(s.Description)
	a.Body = null.StringFrom(s.Body)
	return &a
}
