package model

import "time"

type ArticleResponse struct {
	Slug           string    `json:"slug"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Body           string    `json:"body"`
	TagList        []string  `json:"tagList"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Favorited      bool      `json:"favorited"`
	FavoritesCount int       `json:"favoritesCount"`
	Author         struct {
		Username  string  `json:"username"`
		Bio       *string `json:"bio"`
		Image     *string `json:"image"`
		Following bool    `json:"following"`
	} `json:"author"`
}

type CommentResponse struct {
	ID        uint      `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Author    struct {
		Username  string  `json:"username"`
		Bio       *string `json:"bio"`
		Image     *string `json:"image"`
		Following bool    `json:"following"`
	} `json:"author"`
}

type SimpleArticle struct {
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Body        string    `json:"body"`
	TagList     []string  `json:"tagList,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
type SingleArticle struct {
	Article *ArticleResponse `json:"article"`
}

type SimpleArticleList struct {
	Articles      []*SimpleArticle `json:"articles"`
	ArticlesCount int64            `json:"articlesCount"`
}

type SingleArticleResponse struct {
	Article *ArticleResponse `json:"article"`
}

type CommentListResponse struct {
	Comments []CommentResponse `json:"comments"`
}

type TagListResponse struct {
	Tags []string `json:"tags"`
}

type CommentRequest struct {
	Comment struct {
		Body string `json:"body" validate:"required"`
	} `json:"comment"`
}
