package user

import (
	"forum/model"
	"forum/repository"
	"forum/service"
	"schema/entity"

	"github.com/rs/zerolog/log"
)

type Service struct {
	Repo repository.User
}

func NewUserService(r repository.User) *Service {
	return &Service{
		Repo: r,
	}
}

func (s *Service) CheckUser(user *model.LoginUser) error {
	userInfo, err := s.Repo.FindByEmail(user.Email)
	if err != nil {
		log.Error().Err(err).Msg("FindByEmail error")
		return err
	}

	return service.CheckPassword(userInfo.Password, user.Password)
}

func (s *Service) CreateUser(user *model.RegisterUser) error {
	passWord, err := service.HashPassword(user.Password)
	if err != nil {
		log.Error().Err(err).Msg("HashPassword error")
		return err
	}
	var u entity.User
	u.Username = user.Username
	u.Email = user.Email
	u.Password = passWord
	return s.Repo.CreateUser(&u)
}

func (s *Service) FollowUserByUserName(uid uint, userName string) error {
	targetUser, err := s.Repo.FindUserByUserName(userName)
	if err != nil {
		log.Error().Err(err).Msg("FindByUserName error")
		return err
	}
	loggedUser, err := s.Repo.FindUserByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("findCurrentUserAndTargetUser error")
		return err
	}
	return s.Repo.AddFollower(loggedUser, targetUser)
}

func (s *Service) GetUserByID(uid uint) (*entity.User, error) {
	return s.Repo.FindUserByID(uid)
}

func (s *Service) GetUserByEmail(email string) (*entity.User, error) {
	return s.Repo.FindByEmail(email)
}

func (s *Service) GetUserByUserName(username string) (*entity.User, error) {
	return s.Repo.FindUserByUserName(username)
}

func (s *Service) UnFollowUserByUserName(uid uint, userName string) error {
	targetUser, err := s.Repo.FindUserByUserName(userName)
	if err != nil {
		log.Error().Err(err).Msg("FindByUserName error")
		return err
	}
	loggedUser, err := s.Repo.FindUserByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindByUserID error")
		return err
	}
	return s.Repo.RemoveFollower(loggedUser, targetUser)
}

func (s *Service) GetFollowersByUserID(uid uint) ([]*entity.User, error) {
	currentUser, err := s.Repo.FindUserByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindUserByID error")
		return nil, err
	}
	return s.Repo.GetFollowers(currentUser)
}

func (s *Service) GetFollowingUser(uid uint) ([]*entity.User, error) {
	currentUser, err := s.Repo.FindUserByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindUserByID error")
		return nil, err
	}
	return s.Repo.GetFollowingUsers(currentUser)
}

func (s *Service) UpdateUser(user *entity.User) error {
	return s.Repo.UpdateUser(user)
}
