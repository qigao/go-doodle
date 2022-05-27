// DON'T EDIT: This is generated code

package service

import (
	"schema/entity"
)

// ArticleService ...
type ArticleService interface {
	CreateArticle(a *entity.Article) error
	UpdateArticle(slug string, newArticle *entity.Article) error
	DeleteArticle(slug string) error
	FindArticle(slug string) (*entity.Article, *entity.User, []*entity.Tag, error)
	FindArticleByAuthor(userName string, offset, limit int) ([]*entity.Article, int64, error)
	FindArticles(tag, author string, offset, limit int) ([]*entity.Article, int64, error)
	FindCommentsBySlug(slug string, offset, limit int) ([]*entity.Comment, error)
	FindAuthorBySlug(slug string) (*entity.User, error)
	AddCommentToArticle(slug string, cm *entity.Comment) error
	DeleteCommentFromArticle(slug string, commentId uint64) error
	AddFavoriteArticleBySlug(slug string, uid uint) error
	RemoveFavoriteArticleBySlug(slug string, uid uint) error
	FindArticleAndUserBySlugAndUserID(slug string, uid uint) (*entity.Article, *entity.User, error)
	AddTagToArticle(slug string, tagStr []string) error
	GetAllTags() ([]*entity.Tag, error)
}
