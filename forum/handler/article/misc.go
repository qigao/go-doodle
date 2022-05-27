package article

import (
	"forum/model"
	"forum/service"
	"schema/entity"

	"github.com/volatiletech/null/v8"
)

type Handler struct {
	Service service.ArticleService
}

func NewArticleHandler(as service.ArticleService) *Handler {
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
