package service

import (
	"feed-management/internal/entity"
	"feed-management/internal/repository"
	"feed-management/pkg/dto"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

type PostService struct {
	logger         zerolog.Logger
	validator      *validator.Validate
	postRepository *repository.PostRepository
}

func NewPostService(logger zerolog.Logger, validator *validator.Validate, postRepository *repository.PostRepository) *PostService {
	return &PostService{
		logger:         logger,
		validator:      validator,
		postRepository: postRepository,
	}
}
func (s *PostService) Create(consumer *dto.EventPostConsumer) error {
	if err := s.validator.Struct(consumer); err != nil {
		s.logger.Warn().Err(err).Msg("failed to validate request")
		return err
	}
	post := &entity.Post{
		PostId:       consumer.PostId,
		UserId:       consumer.UserId,
		Description:  consumer.Description,
		TotalLike:    consumer.TotalLike,
		TotalComment: consumer.TotalComment,
		CreatedAt:    consumer.CreatedAt,
		UpdatedAt:    consumer.UpdatedAt,
	}
	if err := s.postRepository.Create(post); err != nil {
		s.logger.Error().Err(err).Msg("failed create to database")
		return err
	}
	s.logger.Info().Str("post_id", strconv.Itoa(int(consumer.PostId))).Msg("post create success")
	return nil
}
