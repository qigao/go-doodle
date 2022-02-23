package user

import (
	"forum/entity"
	"forum/repository"

	"github.com/rs/zerolog/log"
)

type FollowRequest struct {
	Repo repository.User
}

func (r *FollowRequest) FllowUserByUserName(uid uint, userName string) error {
	currentUser, targetUser, err := r.findCurrentUserAndTargetUser(uid, userName)
	if err != nil {
		log.Error().Err(err).Msg("findCurrentUserAndTargetUser error")
		return err
	}
	return r.Repo.AddFollower(currentUser, targetUser)
}

func (r *FollowRequest) UnFllowUserByUserName(uid uint, userName string) error {
	currentUser, targetUser, err := r.findCurrentUserAndTargetUser(uid, userName)
	if err != nil {
		log.Error().Err(err).Msg("findCurrentUserAndTargetUser error")
		return err
	}
	return r.Repo.RemoveFollower(currentUser, targetUser)
}

func (r *FollowRequest) GetFollowersByUserID(uid uint) ([]*entity.User, error) {
	currentUser, err := r.Repo.FindByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindByID error")
		return nil, err
	}
	return r.Repo.GetFollowers(currentUser)
}

func (r *FollowRequest) GetFollowingUser(uid uint) ([]*entity.User, error) {
	currentUser, err := r.Repo.FindByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindByID error")
		return nil, err
	}
	return r.Repo.GetFollowingUsers(currentUser)
}

func (r *FollowRequest) findCurrentUserAndTargetUser(uid uint, userName string) (*entity.User, *entity.User, error) {
	currentUser, err := r.Repo.FindByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindByID error")
		return nil, nil, err
	}
	targetUser, err := r.Repo.FindByUserName(userName)
	if err != nil {
		log.Error().Err(err).Msg("FindByUserName error")
		return nil, nil, err
	}
	return currentUser, targetUser, nil
}
