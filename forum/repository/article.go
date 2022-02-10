package repository

import "forum/entities"

type Article interface {
	FindBySlug(string) (*models.Article, error)
	FindArticleByUserIDAndSlug(userID uint, slug string) (*models.Article, error)
	CreateArticle(*models.Article) error
	UpdateArticle(*models.Article, []string) error
	DeleteArticle(*models.Article) error
	ListArticles(offset, limit int) ([]*models.Article, int64, error)
	ListByTag(tag string, offset, limit int) ([]models.Article, int64, error)
	ListByAuthor(username string, offset, limit int) ([]models.Article, int64, error)
	ListByWhoFavorited(username string, offset, limit int) ([]models.Article, int64, error)
	ListFeed(userID uint, offset, limit int) ([]models.Article, int64, error)

	AddComment(*models.Article, *models.Comment) error
	FindCommentsBySlug(string) ([]models.Comment, error)
	FindCommentByID(uint) (*models.Comment, error)
	DeleteComment(*models.Comment) error

	AddFavorite(*models.Article, uint) error
	RemoveFavorite(*models.Article, uint) error
	ListTags() ([]models.Tag, error)
}
