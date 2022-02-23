package article

import (
	"forum/entity"
	"forum/model"
)

func SingleArticleResponseMapper(a *entity.Article, author *entity.User, tags []*entity.Tag) *singleArticleResponse {
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

	for _, t := range tags {
		ar.TagList = append(ar.TagList, t.Tag.String)
	}

	ar.Author.Username = author.Username
	if author.Bio.Valid {
		ar.Author.Bio = &author.Bio.String
	}
	if author.Image.Valid {
		ar.Author.Image = &author.Image.String
	}
	return &singleArticleResponse{ar}
}

func SimpleArticleListMapper(articles []*entity.Article, count int64) *SimpleArticleList {
	r := new(SimpleArticleList)
	r.Articles = make([]*model.SimpleArticle, 0)
	for _, a := range articles {
		ar := new(model.SimpleArticle)
		ar.Slug = a.Slug
		ar.Title = a.Title
		if a.Description.Valid {
			ar.Description = a.Description.String
		}
		if a.CreatedAt.Valid {
			ar.CreatedAt = a.CreatedAt.Time
		}
		if a.UpdatedAt.Valid {
			ar.UpdatedAt = a.UpdatedAt.Time
		}
		r.Articles = append(r.Articles, ar)
	}
	r.ArticlesCount = count
	return r
}

func TagListResponseMapper(tags []*entity.Tag) *tagListResponse {
	r := new(tagListResponse)
	for _, t := range tags {
		r.Tags = append(r.Tags, t.Tag.String)
	}
	return r
}

func CommentResponseMapper(cm *entity.Comment) *model.CommentResponse {
	comment := new(model.CommentResponse)
	comment.ID = uint(cm.ID)
	comment.Body = cm.Body.String
	if cm.CreatedAt.Valid {
		comment.CreatedAt = cm.CreatedAt.Time
	}
	if cm.UpdatedAt.Valid {
		comment.UpdatedAt = cm.UpdatedAt.Time
	}
	comment.Author.Username = cm.R.User.Username
	if cm.R.User.Bio.Valid {
		comment.Author.Bio = &cm.R.User.Bio.String
	}
	if cm.R.User.Image.Valid {
		comment.Author.Image = &cm.R.User.Image.String
	}
	return comment
}

func CommentListResponseMapper(comments []*entity.Comment) *commentListResponse {
	r := new(commentListResponse)
	cr := new(model.CommentResponse)
	r.Comments = make([]model.CommentResponse, 0)
	for _, i := range comments {
		cr = CommentResponseMapper(i)
		r.Comments = append(r.Comments, *cr)
	}
	return r
}
