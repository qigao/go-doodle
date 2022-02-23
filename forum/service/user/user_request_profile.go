package user

import (
	"forum/entity"
	"forum/repository"

	"github.com/rs/zerolog/log"
)

type UserProfile struct {
	Repo repository.User
}

func (p *UserProfile) GetUserByID(uid uint) (*entity.User, error) {
	u, err := p.Repo.FindByID(uid)
	if err != nil {
		log.Error().Err(err).Msg("FindByID error")
		return nil, err
	}
	return u, nil
}

func (p *UserProfile) GetUserByEmail(email string) (*entity.User, error) {
	u, err := p.Repo.FindByEmail(email)
	if err != nil {
		log.Error().Err(err).Msg("FindByEmail error")
		return nil, err
	}
	return u, nil
}

func (p *UserProfile) GetUserByUsername(username string) (*entity.User, error) {
	u, err := p.Repo.FindByUsername(username)
	if err != nil {
		log.Error().Err(err).Msg("FindByUsername error")
		return nil, err
	}
	return u, nil
}
