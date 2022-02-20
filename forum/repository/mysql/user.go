package mysql

import (
	"context"
	"database/sql"
	"forum/entity"

	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

//UserRepo is a repository for user
type UserRepo struct {
	Db *sql.DB
}

// NewUserRepo returns a new instance of a user repository.
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}
func (u *UserRepo) FindByID(uid uint) (*entity.User, error) {
	user, err := entity.Users(Where("id = ?", uid)).One(context.Background(), u.Db)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) FindByEmail(s string) (*entity.User, error) {
	user, err := entity.Users(Where("email = ?", s)).One(context.Background(), u.Db)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) FindByUsername(s string) (*entity.User, error) {
	user, err := entity.Users(Where("username = ?", s)).One(context.Background(), u.Db)
	if err != nil {
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
	rows, err := user.Update(context.Background(), u.Db, boil.Infer())
	if err != nil {
		log.Error().Err(err).Msg("failed to update user")
		return err
	}
	log.Info().Msgf("%d rows updated", rows)
	tx.Commit()
	if rows == 0 {
		return sql.ErrNoRows
	}
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
	cnts, err := user.FollowerUsers(entity.UserWhere.ID.EQ(follower.ID)).Count(context.Background(), u.Db)
	if cnts == 0 {
		log.Error().Err(err).Msg("failed to check if user is follower")
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
