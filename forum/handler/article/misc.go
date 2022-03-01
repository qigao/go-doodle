package article

import (
	"forum/entity"
	"forum/model"
	"forum/service"

	"github.com/volatiletech/null/v8"
)

type Handler struct {
	Service service.ServiceArticle
}

func NewArticleHandler(as service.ServiceArticle) *Handler {
	return &Handler{
		Service: as,
	}
}

func populateSingleArticle(s *model.SimpleArticle) *entity.Article {
	var a *entity.Article
	a.Title = s.Title
	a.Slug = s.Slug
	a.Description = null.StringFrom(s.Description)
	a.Body = null.StringFrom(s.Body)
	return a
}
