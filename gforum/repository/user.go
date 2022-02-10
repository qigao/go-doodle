package repository

import (
	"gforum/model"
	"gorm.io/gorm"
)

// UserRepository is data access struct for user
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository returns a new UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// GetByEmail finds a user from email
func (s *UserRepository) GetByEmail(email string) (*model.User, error) {
	var m model.User
	if err := s.db.Where("email = ?", email).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

// GetByID finds a user from id
func (s *UserRepository) GetByID(id uint) (*model.User, error) {
	var m model.User
	if err := s.db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

// GetByUsername finds a user from username
func (s *UserRepository) GetByUsername(username string) (*model.User, error) {
	var m model.User
	if err := s.db.Where("username = ?", username).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

// Create create a user
func (s *UserRepository) Create(m *model.User) error {
	return s.db.Create(m).Error
}

// Update update all of user fields
func (s *UserRepository) Update(m *model.User) error {
	return s.db.Model(m).Save(m).Error
}

// IsFollowing returns whether user A follows user BuildDSN or not
func (s *UserRepository) IsFollowing(a *model.User, b *model.User) (bool, error) {
	if a == nil || b == nil {
		return false, nil
	}

	var count int64
	err := s.db.Table("follows").
		Where("user_id = ? AND follow_id = ?", a.ID, b.ID).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Follow create follow relationship to User BuildDSN from user A
func (s *UserRepository) Follow(a *model.User, b *model.User) error {
	a.Follows = append(a.Follows, *b)
	return s.db.Model(a).Save(a).Error
}

// UnFollow delete follow relationship to User BuildDSN from user A
func (s *UserRepository) UnFollow(a *model.User, b *model.User) error {
	return s.db.Model(a).Association("Follows").Delete(b)
}

// GetFollowingUserIDs returns user ids current user follows
func (s *UserRepository) GetFollowingUserIDs(m *model.User) ([]uint, error) {
	rows, err := s.db.Table("follows").
		Select("user_id").
		Where("user_id = ?", m.ID).
		Rows()
	if err != nil {
		return []uint{}, err
	}
	defer rows.Close()

	var ids []uint
	for rows.Next() {
		var id uint
		rows.Scan(&id)
		ids = append(ids, id)
	}

	return ids, nil
}
