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
func (s *PostService) PostCreate(consumer *dto.EventPostConsumer) error {
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
func (s *PostService) PostUpdate(consumer *dto.EventPostUpdatedConsumer) error {
	if err := s.validator.Struct(consumer); err != nil {
		s.logger.Warn().Err(err).Msg("failed to validate request")
		return err
	}
	post, err := s.postRepository.FindOne(consumer.PostId)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed find one to database")
		return err
	}
	post.Description = consumer.Description
	post.UpdatedAt = consumer.UpdatedAt
	if err := s.postRepository.ReplaceOne(post); err != nil {
		s.logger.Error().Err(err).Msg("failed replace one to database")
		return err
	}
	s.logger.Info().Str("post_id", strconv.Itoa(int(consumer.PostId))).Msg("post updated success")
	return nil
}
func (s *PostService) PostLike(consumer *dto.EventLikeTotalConsumer) error {
	if err := s.validator.Struct(consumer); err != nil {
		s.logger.Warn().Err(err).Msg("failed to validate request")
		return err
	}
	post, err := s.postRepository.FindOne(consumer.PostId)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed find one to database")
		return err
	}
	if consumer.Total {
		post.TotalLike += 1
	} else {
		post.TotalLike -= 1
	}
	if err := s.postRepository.ReplaceOne(post); err != nil {
		s.logger.Error().Err(err).Msg("failed replace one to database")
		return err
	}
	s.logger.Info().Str("post_id", strconv.Itoa(int(consumer.PostId))).Msg("post like success")
	return nil
}
func (s *PostService) CommentTotal(consumer *dto.EventCommentTotalConsumer) error {
	if err := s.validator.Struct(consumer); err != nil {
		s.logger.Warn().Err(err).Msg("failed to validate request")
		return err
	}
	post, err := s.postRepository.FindOne(consumer.PostId)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed find one to database")
		return err
	}
	if consumer.Total {
		post.TotalComment += 1
	} else {
		post.TotalComment -= 1
	}
	if err := s.postRepository.ReplaceOne(post); err != nil {
		s.logger.Error().Err(err).Msg("failed replace one to database")
		return err
	}
	s.logger.Info().Str("post_id", strconv.Itoa(int(consumer.PostId))).Msg("post comment success success")
	return nil
}
