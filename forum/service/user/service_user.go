package user

import (
	"forum/entity"
	"forum/model"
	"forum/repository"
	"forum/service"

	"github.com/rs/zerolog/log"
)

type ServiceUser struct {
	Repo repository.User
}

func NewServiceUser(r repository.User) *ServiceUser {
	return &ServiceUser{
		Repo: r,
	}
}

func (r *ServiceUser) CheckUser(user *model.LoginUser) error {
	userInfo, err := r.Repo.FindByEmail(user.Email)
	if err != nil {
		log.Error().Err(err).Msg("FindByEmail error")
		return err
	}

	return service.CheckPassword(userInfo.Password, user.Password)
}

func (r *ServiceUser) CreateUser(user *model.RegisterUser) error {
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

func (r *ServiceUser) FllowUserByUserName(uid uint, userName string) error {
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

func (p *ServiceUser) GetUserByID(uid uint) (*entity.User, error) {
	return p.Repo.FindUserByID(uid)
}

func (p *ServiceUser) GetUserByEmail(email string) (*entity.User, error) {
	return p.Repo.FindByEmail(email)
}

func (p *ServiceUser) GetUserByUserName(username string) (*entity.User, error) {
	return p.Repo.FindUserByUserName(username)
}

func (r *ServiceUser) UnFollowUserByUserName(uid uint, userName string) error {
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

func (r *ServiceUser) GetFollowersByUserID(uid uint) ([]*entity.User, error) {
	currentUser, err := r.Repo.FindUserByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindUserByID error")
		return nil, err
	}
	return r.Repo.GetFollowers(currentUser)
}

func (r *ServiceUser) GetFollowingUser(uid uint) ([]*entity.User, error) {
	currentUser, err := r.Repo.FindUserByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindUserByID error")
		return nil, err
	}
	return r.Repo.GetFollowingUsers(currentUser)
}
