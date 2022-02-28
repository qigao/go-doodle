package article

import (
	"forum/entity"
	"forum/model"
	"forum/service/article"

	"github.com/volatiletech/null/v8"
)

type Handler struct {
	service article.ServiceArticle
}

func NewArticleHandler(as article.ServiceArticle) *Handler {
	return &Handler{
		service: as,
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
