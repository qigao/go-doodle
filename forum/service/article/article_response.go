package article

import (
	"forum/entity"
	"forum/handler"
	"forum/model"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
)

type singleArticleResponse struct {
	Article *model.ArticleResponse `json:"article"`
}

type articleListResponse struct {
	Articles      []*model.ArticleResponse `json:"articles"`
	ArticlesCount int64                    `json:"articlesCount"`
}

func NewArticleResponse(c echo.Context, a *entity.Article) *singleArticleResponse {
	ar := new(model.ArticleResponse)
	ar.TagList = make([]string, 0)
	ar.Slug = a.Slug
	ar.Title = a.Title
	if a.Description.Valid {
		ar.Description = a.Description.String
	}
	if a.Body.Valid {
		ar.Body = a.Body.String
	}
	if a.CreatedAt.Valid {
		ar.CreatedAt = a.CreatedAt.Time
	}
	if a.UpdatedAt.Valid {
		ar.UpdatedAt = a.UpdatedAt.Time
	}

	for _, t := range a.Tags {
		ar.TagList = append(ar.TagList, t.Tag)
	}
	for _, u := range a.Favorites {
		if u.ID == handler.UserIDFromToken(c) {
			ar.Favorited = true
		}
	}
	ar.FavoritesCount = len(a.Favorites)
	ar.Author.Username = a.Author(entity.UserWhere.ID.EQ(a.AuthorID.Uint64))
	ar.Author.Image = a.Author.Image
	ar.Author.Bio = a.Author.Bio
	ar.Author.Following = a.Author.FollowedBy(handler.UserIDFromToken(c))
	return &singleArticleResponse{ar}
}

func NewArticleListResponse(userID uint, articles []entity.Article, count int64) *articleListResponse {
	r := new(articleListResponse)
	r.Articles = make([]*model.ArticleResponse, 0)
	for _, a := range articles {
		ar := new(model.ArticleResponse)
		ar.TagList = make([]string, 0)
		ar.Slug = a.Slug
		ar.Title = a.Title
		if a.Description.Valid {
			ar.Description = a.Description.String
		}
		if a.Body.Valid {
			ar.Body = a.Body.String
		}
		if a.CreatedAt.Valid {
			ar.CreatedAt = a.CreatedAt.Time
		}
		if a.UpdatedAt.Valid {
			ar.UpdatedAt = a.UpdatedAt.Time
		}
		for _, t := range a.Tags {
			ar.TagList = append(ar.TagList, t.Tag)
		}
		for _, u := range a.Favorites {
			if u.ID == userID {
				ar.Favorited = true
			}
		}
		ar.FavoritesCount = len(a.Favorites)
		ar.Author.Username = a.Author.Username
		ar.Author.Image = a.Author.Image
		ar.Author.Bio = a.Author.Bio
		r.Articles = append(r.Articles, ar)
	}
	r.ArticlesCount = count
	return r
}
