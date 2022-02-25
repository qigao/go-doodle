package article

import "forum/model"

type singleArticle struct {
	Article *model.ArticleResponse `json:"article"`
}

type SimpleArticleList struct {
	Articles      []*model.SimpleArticle `json:"articles"`
	ArticlesCount int64                  `json:"articlesCount"`
}

type singleArticleResponse struct {
	Article *model.ArticleResponse `json:"article"`
}

type commentListResponse struct {
	Comments []model.CommentResponse `json:"comments"`
}

type tagListResponse struct {
	Tags []string `json:"tags"`
}

type CommentRequest struct {
	Comment struct {
		Body string `json:"body" validate:"required"`
	} `json:"comment"`
}
