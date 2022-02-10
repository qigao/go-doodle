package mysql

import (
	"context"
	"database/sql"
	"forum/entities"
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
func (u *UserRepo) FindByID(u2 uint) (*models.User, error) {
	user, err := models.Users(Where("id = ?", u2)).One(context.Background(), u.Db)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) FindByEmail(s string) (*models.User, error) {
	user, err := models.Users(Where("email = ?", s)).One(context.Background(), u.Db)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) FindByUsername(s string) (*models.User, error) {
	user, err := models.Users(Where("username = ?", s)).One(context.Background(), u.Db)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) CreateUser(user *models.User) error {
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

func (u *UserRepo) UpdateUser(user *models.User) error {
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
	return nil
}

func (u *UserRepo) AddFollower(user *models.User, follower *models.User) error {
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

func (u *UserRepo) RemoveFollower(user *models.User, follower *models.User) error {
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

func (u *UserRepo) IsFollower(user, follower *models.User) (bool, error) {
	cnts, err := user.FollowerUsers(Where("id = ?", follower.ID)).Count(context.Background(), u.Db)
	if cnts > 0 {
		return true, nil
	}
	return false, err
}
