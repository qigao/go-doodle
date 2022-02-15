package repository

import "forum/entity"

type Article interface {
	FindBySlug(string) (*entity.Article, error)
	FindArticleByUserIDAndSlug(userID uint64, slug string) (*entity.Article, error)
	Create(*entity.Article) error
	Update(*entity.Article) error
	Delete(*entity.Article) error

	List(offset, limit int) ([]*entity.Article, int64, error)
	ListByTag(tag string, offset, limit int) ([]*entity.Article, int64, error)
	ListByAuthor(username string, offset, limit int) ([]*entity.Article, int64, error)
	ListByWhoFavorited(username string, offset, limit int) ([]*entity.Article, int64, error)
	ListFeed(userID uint, offset, limit int) ([]*entity.Article, int64, error)

	AddComment(*entity.Article, *entity.Comment) error
	FindCommentByID(commentID uint64) (*entity.Comment, error)
	FindCommentsBySlug(string) ([]*entity.Comment, error)
	FindCommentsByArticleID(uint) ([]*entity.Comment, error)
	DeleteComment(*entity.Comment) error

	AddFavorite(*entity.Article, uint) error
	RemoveFavorite(*entity.Article, uint) error

	CreateTag(*entity.Tag) error
	AddTag(*entity.Article, *entity.Tag) error
	RemoveTag(*entity.Article, *entity.Tag) error
	FindTagsByArticleID(article *entity.Article) ([]*entity.Tag, error)
	ListTags() ([]*entity.Tag, error)
}
