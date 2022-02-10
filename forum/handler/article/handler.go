package article

import (
	"forum/repository"
)

type Handler struct {
	article repository.Article
}

func NewArticleHandler(as repository.Article) *Handler {
	return &Handler{
		article: as,
	}
}
