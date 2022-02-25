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
	ListArticlesByAuthor(user *entity.User, offset, limit int) ([]*entity.Article, int64, error)
	FindFavoriteArticlesByUser(user *entity.User, offset, limit int) ([]*entity.Article, int64, error)
	ListFeed(userID uint, offset, limit int) ([]*entity.Article, int64, error)
	FindAuthorByArticle(article *entity.Article) (*entity.User, error)

	AddComment(*entity.Article, *entity.Comment) error
	FindCommentByID(commentID uint64) (*entity.Comment, error)
	FindCommentsByArticle(*entity.Article, int, int) ([]*entity.Comment, error)
	DeleteComment(*entity.Comment) error
	DeleteCommentByCommentID(commentID uint64) error
	DeleteCommentByArticle(article *entity.Article, comment *entity.Comment) error

	AddFavoriteArticle(*entity.Article, *entity.User) error
	RemoveFavorite(*entity.Article, *entity.User) error

	CreateTag(*entity.Tag) error
	AddTagToArticle(*entity.Article, *entity.Tag) error
	AddTagsToArticle(*entity.Article, []*entity.Tag) error
	RemoveTagFromArticle(*entity.Article, *entity.Tag) error
	FindTagsByArticle(artile *entity.Article) ([]*entity.Tag, error)
	ListTags() ([]*entity.Tag, error)
}
