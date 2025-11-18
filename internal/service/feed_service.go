package service

import (
	"feed-management/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

type FeedService struct {
	logger         zerolog.Logger
	validator      *validator.Validate
	postRepository *repository.PostRepository
}

func NewFeedService(logger zerolog.Logger, validator *validator.Validate, postRepository *repository.PostRepository) *FeedService {
	return &FeedService{
		logger:         logger,
		validator:      validator,
		postRepository: postRepository,
	}
}
