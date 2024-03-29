// Code generated by ifacemaker; DO NOT EDIT.

package repository

import (
	"schema/entity"
)

// IRepoArticle ...
type IRepoArticle interface {
	FindArticleBySlug(s string) (*entity.Article, error)
	FindArticleByAuthorIDAndSlug(userID uint64, slug string) (*entity.Article, error)
	CreateArticle(article *entity.Article) error
	// UpdateArticle  update article
	UpdateArticle(article *entity.Article) error
	DeleteArticle(article *entity.Article) error
	// FindArticles all the articles with pagination
	FindArticles(offset, limit int) ([]*entity.Article, int64, error)
	ListArticlesByTag(tagStr string, offset, limit int) ([]*entity.Article, int64, error)
	ListArticlesByAuthor(user *entity.User, offset, limit int) ([]*entity.Article, int64, error)
	FindAuthorByArticle(article *entity.Article) (*entity.User, error)
	ListFeed(userID uint, offset, limit int) ([]*entity.Article, int64, error)
	AddComment(article *entity.Article, comment *entity.Comment) error
	FindCommentsByArticle(article *entity.Article, offset int, limit int) ([]*entity.Comment, error)
	FindCommentByID(commentID uint64) (*entity.Comment, error)
	DeleteComment(comment *entity.Comment) error
	DeleteCommentByCommentID(commentID uint64) error
	DeleteCommentByArticle(article *entity.Article, comment *entity.Comment) error
	AddFavoriteArticle(article *entity.Article, user *entity.User) error
	RemoveFavorite(article *entity.Article, user *entity.User) error
	FindFavoriteArticlesByUser(user *entity.User, offset, limit int) ([]*entity.Article, int64, error)
	CreateTag(tag *entity.Tag) error
	AddTagToArticle(article *entity.Article, tag *entity.Tag) error
	AddTagsToArticle(article *entity.Article, tag []*entity.Tag) error
	RemoveTagFromArticle(article *entity.Article, tag *entity.Tag) error
	RemoveTagsFromArticle(article *entity.Article, tags []*entity.Tag) error
	FindTagsByArticle(article *entity.Article) ([]*entity.Tag, error)
	ListTags() ([]*entity.Tag, error)
}
