package repository

import "forum/entity"

type Article interface {
	FindArticleBySlug(string) (*entity.Article, error)
	FindArticleByAuthorIDAndSlug(userID uint64, slug string) (*entity.Article, error)
	CreateArticle(*entity.Article) error
	UpdateArticle(*entity.Article) error
	DeleteArticle(*entity.Article) error

	ListArticles(offset, limit int) ([]*entity.Article, int64, error)
	ListArticlesByTag(tag string, offset, limit int) ([]*entity.Article, int64, error)
	ListArticlesByAuthor(username string, offset, limit int) ([]*entity.Article, int64, error)
	ListArticlesByWhoFavorited(username string, offset, limit int) ([]*entity.Article, int64, error)
	ListFeed(userID uint, offset, limit int) ([]*entity.Article, int64, error)
	FindAuthorBySlug(slug string) (*entity.User, error)
	AddComment(*entity.Article, *entity.Comment) error
	FindCommentByID(commentID uint64) (*entity.Comment, error)
	FindCommentsBySlug(string, int, int) ([]*entity.Comment, error)
	DeleteComment(*entity.Comment) error
	DeleteCommentBySlugAndCommentID(slug string, commentID uint64) error

	AddFavorite(*entity.Article, uint) error
	RemoveFavorite(*entity.Article, uint) error

	CreateTag(*entity.Tag) error
	AddTagToArticle(*entity.Article, *entity.Tag) error
	AddTagsToArticle(*entity.Article, []*entity.Tag) error
	RemoveTag(*entity.Article, *entity.Tag) error
	FindTagsBySlug(slug string) ([]*entity.Tag, error)
	ListTags() ([]*entity.Tag, error)
}
