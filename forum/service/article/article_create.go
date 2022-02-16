package article

import (
	"forum/entity"
	"forum/model"
	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
)

type CreateArticleRequest struct {
	Article model.CreateArticle `json:"article"`
}

func (r *CreateArticleRequest) Bind(c echo.Context, a *entity.Article, tags []*entity.Tag) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.Title = r.Article.Title
	a.Slug = slug.Make(r.Article.Title)
	a.Description = null.StringFrom(r.Article.Description)
	a.Body = null.StringFrom(r.Article.Body)
	if r.Article.Tags != nil {
		for _, t := range r.Article.Tags {
			tags = append(tags, &entity.Tag{Tag: null.StringFrom(t)})
		}
	}
	return nil
}
