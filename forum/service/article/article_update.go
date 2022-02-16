package article

import (
	"forum/entity"
	"forum/model"
	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
)

type UpdateArticleRequest struct {
	Article model.UpdateArticle `json:"article"`
}

func (r *UpdateArticleRequest) Populate(a *entity.Article) {
	r.Article.Title = a.Title
	if a.Description.Valid {
		r.Article.Description = a.Description.String
	}
	if a.Body.Valid {
		r.Article.Body = a.Body.String
	}
}

func (r *UpdateArticleRequest) Bind(c echo.Context, a *entity.Article) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.Title = r.Article.Title
	a.Slug = slug.Make(a.Title)

	a.Description = null.StringFrom(r.Article.Description)
	a.Body = null.StringFrom(r.Article.Body)
	return nil
}
