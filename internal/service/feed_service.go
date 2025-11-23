package service

import (
	"feed-management/internal/repository"
	"feed-management/pkg/dto"

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

func (s *FeedService) Foryou() ([]dto.PostResponse, error) {
	posts, err := s.postRepository.FindForYou()
	if err != nil {
		s.logger.Error().Err(err).Msg("failed find for you to database")
		return nil, err
	}
	resp := make([]dto.PostResponse, 0, len(posts))
	if len(posts) != 0 {
		for _, x := range posts {
			resp = append(resp, dto.PostResponse{
				PostId:       x.PostId,
				UserId:       x.UserId,
				Description:  x.Description,
				TotalLike:    x.TotalLike,
				TotalComment: x.TotalComment,
				CreatedAt:    x.CreatedAt,
				UpdatedAt:    x.UpdatedAt,
			})
		}
	}
	return resp, nil
}
