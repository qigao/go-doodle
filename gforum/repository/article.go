package repository

import (
	"gforum/model"
	"gorm.io/gorm"
)

// ArticleRepository is data access struct for user
type ArticleRepository struct {
	db *gorm.DB
}

// NewArticleRepository returns a new ArticleRepository
func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{
		db: db,
	}
}

// GetByID finds an article from id
func (s *ArticleRepository) GetByID(id uint) (*model.Article, error) {
	var m model.Article
	err := s.db.Preload("Tags").Preload("Author").Find(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Create creates an article
func (s *ArticleRepository) Create(m *model.Article) error {
	return s.db.Create(&m).Error
}

// Update updates an article
func (s *ArticleRepository) Update(m *model.Article) error {
	return s.db.Model(&m).Save(&m).Error
}

// GetArticles get global articles
func (s *ArticleRepository) GetArticles(tagName, username string, favoritedBy *model.User, limit, offset int) ([]model.Article, error) {
	d := s.db.Preload("Author")

	// author query (has one)
	if username != "" {
		d = d.Joins("join users on articles.user_id = users.id").
			Where("users.username = ?", username)
	}

	// tag query (many to many)
	if tagName != "" {
		d = d.Joins(
			"join article_tags on articles.id = article_tags.article_id "+
				"join tags on tags.id = article_tags.tag_id").
			Where("tags.name = ?", tagName)
	}

	// favorited query
	if favoritedBy != nil {
		rows, err := s.db.Select("article_id").
			Table("favorite_articles").
			Where("user_id = ?", favoritedBy.ID).
			Offset(offset).Limit(limit).Rows()
		if err != nil {
			return []model.Article{}, err
		}
		defer rows.Close()

		var ids []uint
		for rows.Next() {
			var id uint
			rows.Scan(&id)
			ids = append(ids, id)
		}
		d = d.Where("user_id in (?)", ids)
	}

	// offset query, limit query
	d = d.Offset(offset).Limit(limit)

	var as []model.Article
	err := d.Find(&as).Error

	return as, err
}

// GetFeedArticles returns following users' articles
func (s *ArticleRepository) GetFeedArticles(userIDs []uint, limit, offset int) ([]model.Article, error) {
	d := s.db.Preload("Author").
		Where("user_id in (?)", userIDs)

	// offset query, limit query
	d = d.Offset(offset).Limit(limit)

	var as []model.Article
	err := d.Find(&as).Error

	return as, err
}

// Delete deletes an article
func (s *ArticleRepository) Delete(m *model.Article) error {
	return s.db.Delete(m).Error
}

// IsFavorited returns whether the article is favorited by the user
func (s *ArticleRepository) IsFavorited(a *model.Article, u *model.User) (bool, error) {
	if a == nil || u == nil {
		return false, nil
	}

	var count int64
	err := s.db.Table("favorite_articles").
		Where("article_id = ? AND user_id = ?", a.ID, u.ID).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// AddFavorite favorite an article
func (s *ArticleRepository) AddFavorite(a *model.Article, u *model.User) error {
	tx := s.db.Begin()
	//
	//err := tx.Model(a).Association("FavoritedUsers").
	//	Append(u).Error
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//
	//err = tx.Model(a).
	//	Update("favorites_count", gorm.Expr("favorites_count + ?", 1)).Error
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}

	tx.Commit()
	a.FavoritesCount++

	return nil
}

// DeleteFavorite unfavorite an article
func (s *ArticleRepository) DeleteFavorite(a *model.Article, u *model.User) error {
	//tx := s.db.Begin()
	//
	//err := tx.Model(a).Association("FavoritedUsers").
	//	Delete(u).Error
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}

	//err = tx.Model(a).
	//	Update("favorites_count", gorm.Expr("favorites_count - ?", 1)).Error
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//
	//tx.Commit()
	//a.FavoritesCount--

	return nil
}

// GetTags creates a article tag
func (s *ArticleRepository) GetTags() ([]model.Tag, error) {
	var tags []model.Tag
	if err := s.db.Find(&tags).Error; err != nil {
		return tags, err
	}
	return tags, nil
}

// CreateComment creates a comment of the article
func (s *ArticleRepository) CreateComment(m *model.Comment) error {
	return s.db.Create(&m).Error
}

// GetComments gets coments of the article
func (s *ArticleRepository) GetComments(m *model.Article) ([]model.Comment, error) {
	var cs []model.Comment
	err := s.db.Preload("Author").
		Where("article_id = ?", m.ID).
		Find(&cs).Error
	if err != nil {
		return cs, err
	}
	return cs, nil
}

// GetCommentByID finds an comment from id
func (s *ArticleRepository) GetCommentByID(id uint) (*model.Comment, error) {
	var m model.Comment
	err := s.db.Find(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// DeleteComment deletes an comment
func (s *ArticleRepository) DeleteComment(m *model.Comment) error {
	return s.db.Delete(m).Error
}
