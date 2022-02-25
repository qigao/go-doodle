package user

import (
	"fmt"
	"forum/entity"
	"forum/model"
	"forum/repository"
	"forum/service"

	"github.com/rs/zerolog/log"
)

type UserRequest struct {
	Repo repository.User
}

func (r *UserRequest) ValidateUser(user *model.LoginUser) error {
	userInfo, err := r.Repo.FindByEmail(user.Email)
	if err != nil {
		return fmt.Errorf("user not found")
	}
	if !service.CheckPassword(userInfo.Password, user.Password) {
		return fmt.Errorf("password is not correct")
	}
	return nil
}

func (r *UserRequest) CreateUser(user *model.RegisterUser) error {
	passWord, err := service.HashPassword(user.Password)
	if err != nil {
		return err
	}
	var u *entity.User
	u.Username = user.Username
	u.Email = user.Email
	u.Password = passWord
	return r.Repo.CreateUser(u)
}

func (r *UserRequest) FllowUserByUserName(uid uint, userName string) error {
	currentUser, targetUser, err := r.findCurrentUserAndTargetUser(uid, userName)
	if err != nil {
		log.Error().Err(err).Msg("findCurrentUserAndTargetUser error")
		return err
	}
	return r.Repo.AddFollower(currentUser, targetUser)
}
func (p *UserRequest) GetUserByID(uid uint) (*entity.User, error) {
	u, err := p.Repo.FindByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindByID error")
		return nil, err
	}
	return u, nil
}

func (p *UserRequest) GetUserByEmail(email string) (*entity.User, error) {
	u, err := p.Repo.FindByEmail(email)
	if err != nil {
		log.Error().Err(err).Msg("FindByEmail error")
		return nil, err
	}
	return u, nil
}

func (p *UserRequest) GetUserByUsername(username string) (*entity.User, error) {
	u, err := p.Repo.FindUserByUserName(username)
	if err != nil {
		log.Error().Err(err).Msg("FindByUserName error")
		return nil, err
	}
	return u, nil
}

func (r *UserRequest) UnFllowUserByUserName(uid uint, userName string) error {
	currentUser, targetUser, err := r.findCurrentUserAndTargetUser(uid, userName)
	if err != nil {
		log.Error().Err(err).Msg("findCurrentUserAndTargetUser error")
		return err
	}
	return r.Repo.RemoveFollower(currentUser, targetUser)
}

func (r *UserRequest) GetFollowersByUserID(uid uint) ([]*entity.User, error) {
	currentUser, err := r.Repo.FindByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindByID error")
		return nil, err
	}
	return r.Repo.GetFollowers(currentUser)
}

func (r *UserRequest) GetFollowingUser(uid uint) ([]*entity.User, error) {
	currentUser, err := r.Repo.FindByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindByID error")
		return nil, err
	}
	return r.Repo.GetFollowingUsers(currentUser)
}

func (r *UserRequest) findCurrentUserAndTargetUser(uid uint, userName string) (*entity.User, *entity.User, error) {
	currentUser, err := r.Repo.FindByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindByID error")
		return nil, nil, err
	}
	targetUser, err := r.Repo.FindUserByUserName(userName)
	if err != nil {
		log.Error().Err(err).Msg("FindByUserName error")
		return nil, nil, err
	}
	return currentUser, targetUser, nil
}
