package handler

import (
	"gforum/repository"
	"github.com/rs/zerolog"
)

// Handler definition
type Handler struct {
	logger *zerolog.Logger
	us     *repository.UserRepository
	as     *repository.ArticleRepository
}

// New returns a new handler with logger and database
func New(l *zerolog.Logger, us *repository.UserRepository, as *repository.ArticleRepository) *Handler {
	return &Handler{logger: l, us: us, as: as}
}
