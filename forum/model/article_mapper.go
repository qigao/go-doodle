package model

import (
	"forum/entity"
	"forum/handler"
	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
)

type singleArticleResponse struct {
	Article *ArticleResponse `json:"article"`
}

type articleListResponse struct {
	Articles      []*ArticleResponse `json:"articles"`
	ArticlesCount int64              `json:"articlesCount"`
}

func NewArticleResponse(c echo.Context, a *entity.Article) *singleArticleResponse {
	ar := new(ArticleResponse)
	ar.TagList = make([]string, 0)
	ar.Slug = a.Slug
	ar.Title = a.Title
	ar.Description = a.Description
	ar.Body = a.Body
	ar.CreatedAt = a.CreatedAt
	ar.UpdatedAt = a.UpdatedAt
	for _, t := range a.Tags {
		ar.TagList = append(ar.TagList, t.Tag)
	}
	for _, u := range a.Favorites {
		if u.ID == handler.UserIDFromToken(c) {
			ar.Favorited = true
		}
	}
	ar.FavoritesCount = len(a.Favorites)
	ar.Author.Username = a.Author.Username
	ar.Author.Image = a.Author.Image
	ar.Author.Bio = a.Author.Bio
	ar.Author.Following = a.Author.FollowedBy(handler.UserIDFromToken(c))
	return &singleArticleResponse{ar}
}

func NewArticleListResponse(userID uint, articles []entity.Article, count int64) *articleListResponse {
	r := new(articleListResponse)
	r.Articles = make([]*ArticleResponse, 0)
	for _, a := range articles {
		ar := new(ArticleResponse)
		ar.TagList = make([]string, 0)
		ar.Slug = a.Slug
		ar.Title = a.Title
		ar.Description = a.Description
		ar.Body = a.Body
		ar.CreatedAt = a.CreatedAt
		ar.UpdatedAt = a.UpdatedAt
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

type singleCommentResponse struct {
	Comment *CommentResponse `json:"comment"`
}

type commentListResponse struct {
	Comments []CommentResponse `json:"comments"`
}

func NewCommentResponse(c echo.Context, cm *entity.Comment) *singleCommentResponse {
	comment := new(CommentResponse)
	comment.ID = cm.ID
	comment.Body = cm.Body
	comment.CreatedAt = cm.CreatedAt
	comment.UpdatedAt = cm.UpdatedAt
	comment.Author.Username = cm.User.Username
	comment.Author.Image = cm.User.Image
	comment.Author.Bio = cm.User.Bio
	comment.Author.Following = cm.User.FollowedBy(handler.UserIDFromToken(c))
	return &singleCommentResponse{comment}
}

func NewCommentListResponse(c echo.Context, comments []entity.Comment) *commentListResponse {
	r := new(commentListResponse)
	cr := CommentResponse{}
	r.Comments = make([]CommentResponse, 0)
	for _, i := range comments {
		cr.ID = i.ID
		cr.Body = i.Body
		cr.CreatedAt = i.CreatedAt
		cr.UpdatedAt = i.UpdatedAt
		cr.Author.Username = i.User.Username
		cr.Author.Image = i.User.Image
		cr.Author.Bio = i.User.Bio
		cr.Author.Following = i.User.FollowedBy(handler.UserIDFromToken(c))

		r.Comments = append(r.Comments, cr)
	}
	return r
}

type tagListResponse struct {
	Tags []string `json:"tags"`
}

func NewTagListResponse(tags []entity.Tag) *tagListResponse {
	r := new(tagListResponse)
	for _, t := range tags {
		r.Tags = append(r.Tags, t.Tag)
	}
	return r
}

type CreateRequest struct {
	Article CreateArticle `json:"article"`
}

func (r *CreateRequest) Bind(c echo.Context, a *entity.Article) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.Title = r.Article.Title
	a.Slug = slug.Make(r.Article.Title)
	a.Description = r.Article.Description
	a.Body = r.Article.Body
	if r.Article.Tags != nil {
		for _, t := range r.Article.Tags {
			a.Tags = append(a.Tags, entity.Tag{Tag: t})
		}
	}
	return nil
}

type UpdateRequest struct {
	Article UpdateArticle `json:"article"`
}

func (r *UpdateRequest) Populate(a *entity.Article) {
	r.Article.Title = a.Title
	r.Article.Description = a.Description
	r.Article.Body = a.Body
}

func (r *UpdateRequest) Bind(c echo.Context, a *entity.Article) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.Title = r.Article.Title
	a.Slug = slug.Make(a.Title)
	a.Description = r.Article.Description
	a.Body = r.Article.Body
	return nil
}

type CommentRequest struct {
	Comment struct {
		Body string `json:"body" validate:"required"`
	} `json:"comment"`
}

func (r *CommentRequest) Bind(c echo.Context, cm *entity.Comment) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	cm.Body = r.Comment.Body
	cm.UserID = handler.UserIDFromToken(c)
	return nil
}
