package user

import (
	"forum/entity"
	"forum/model"
	"forum/repository"
	"forum/service"

	"github.com/rs/zerolog/log"
)

type RequestUser struct {
	Repo repository.User
}

func NewRequestUser(r repository.User) *RequestUser {
	return &RequestUser{
		Repo: r,
	}
}

func (r *RequestUser) CheckUser(user *model.LoginUser) error {
	userInfo, err := r.Repo.FindByEmail(user.Email)
	if err != nil {
		log.Error().Err(err).Msg("FindByEmail error")
		return err
	}

	return service.CheckPassword(userInfo.Password, user.Password)
}

func (r *RequestUser) CreateUser(user *model.RegisterUser) error {
	passWord, err := service.HashPassword(user.Password)
	if err != nil {
		log.Error().Err(err).Msg("HashPassword error")
		return err
	}
	var u entity.User
	u.Username = user.Username
	u.Email = user.Email
	u.Password = passWord
	return r.Repo.CreateUser(&u)
}

func (r *RequestUser) FllowUserByUserName(uid uint, userName string) error {
	targetUser, err := r.Repo.FindUserByUserName(userName)
	if err != nil {
		log.Error().Err(err).Msg("FindByUserName error")
		return err
	}
	loggedUser, err := r.Repo.FindUserByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("findCurrentUserAndTargetUser error")
		return err
	}
	return r.Repo.AddFollower(loggedUser, targetUser)
}

func (p *RequestUser) GetUserByID(uid uint) (*entity.User, error) {
	return p.Repo.FindUserByID(uid)
}

func (p *RequestUser) GetUserByEmail(email string) (*entity.User, error) {
	return p.Repo.FindByEmail(email)
}

func (p *RequestUser) GetUserByUserName(username string) (*entity.User, error) {
	return p.Repo.FindUserByUserName(username)
}

func (r *RequestUser) UnFollowUserByUserName(uid uint, userName string) error {
	targetUser, err := r.Repo.FindUserByUserName(userName)
	if err != nil {
		log.Error().Err(err).Msg("FindByUserName error")
		return err
	}
	loggedUser, err := r.Repo.FindUserByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindByUserID error")
		return err
	}
	return r.Repo.RemoveFollower(loggedUser, targetUser)
}

func (r *RequestUser) GetFollowersByUserID(uid uint) ([]*entity.User, error) {
	currentUser, err := r.Repo.FindUserByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindUserByID error")
		return nil, err
	}
	return r.Repo.GetFollowers(currentUser)
}

func (r *RequestUser) GetFollowingUser(uid uint) ([]*entity.User, error) {
	currentUser, err := r.Repo.FindUserByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindUserByID error")
		return nil, err
	}
	return r.Repo.GetFollowingUsers(currentUser)
}
