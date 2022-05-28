package mysql

import (
	"context"
	"database/sql"
	"schema/entity"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// UserRepo is a repository for user
type UserRepo struct {
	Db *sql.DB
}

// NewUserRepo returns a new instance of a user repository.
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}

func (u *UserRepo) FindUserByID(uid uint) (*entity.User, error) {
	user, err := entity.Users(qm.Where("id = ?", uid)).One(context.Background(), u.Db)
	if err != nil {
		log.Error().Err(err).Msg("error in finding user by id")
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) FindByEmail(s string) (*entity.User, error) {
	user, err := entity.Users(qm.Where("email = ?", s)).One(context.Background(), u.Db)
	if err != nil {
		log.Error().Err(err).Msg("error in finding user by email")
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) FindUserByUserName(s string) (*entity.User, error) {
	user, err := entity.Users(qm.Where("username = ?", s)).One(context.Background(), u.Db)
	if err != nil {
		log.Error().Err(err).Msg("error in finding user by username")
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) CreateUser(user *entity.User) error {
	tx, err := u.Db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	err = user.Insert(context.Background(), u.Db, boil.Infer())
	if err != nil {
		log.Error().Err(err).Msg("failed to create user")
		return err
	}
	tx.Commit()
	return nil
}

func (u *UserRepo) UpdateUser(user *entity.User) error {
	tx, err := u.Db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	_, err = user.Update(context.Background(), u.Db, boil.Infer())
	if err != nil {
		log.Error().Err(err).Msg("failed to update user")
		return err
	}
	tx.Commit()
	return nil
}

func (u *UserRepo) AddFollower(user *entity.User, follower *entity.User) error {
	tx, err := u.Db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	err = user.AddFollowerUsers(context.Background(), u.Db, false, follower)
	if err != nil {
		log.Error().Err(err).Msg("failed to add follower")
		return err
	}
	tx.Commit()
	return nil
}

func (u *UserRepo) RemoveFollower(user *entity.User, follower *entity.User) error {
	tx, err := u.Db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to start transaction")
		return err
	}
	defer tx.Rollback()
	err = user.RemoveFollowerUsers(context.Background(), u.Db, follower)
	if err != nil {
		log.Error().Err(err).Msg("failed to remove follower")
		return err
	}
	tx.Commit()
	return nil
}

func (u *UserRepo) IsFollower(user, follower *entity.User) (bool, error) {
	_, err := user.FollowerUsers(qm.Where("follower_id=?", follower.ID)).One(context.Background(), u.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to check follower")
		return false, nil
	}
	return true, err
}

func (u *UserRepo) GetFollowers(user *entity.User) ([]*entity.User, error) {
	followers, err := user.FollowerUsers().All(context.Background(), u.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to get followers")
		return nil, err
	}
	return followers, nil
}

func (u *UserRepo) GetFollowingUsers(user *entity.User) ([]*entity.User, error) {
	following, err := user.FollowingUsers().All(context.Background(), u.Db)
	if err != nil {
		log.Error().Err(err).Msg("failed to get following")
		return nil, err
	}
	return following, nil
}
