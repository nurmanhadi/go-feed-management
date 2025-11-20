package consumer

import (
	"feed-management/internal/service"
	"feed-management/pkg"
	"feed-management/pkg/dto"

	"github.com/goccy/go-json"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

type PostConsumer struct {
	logger      zerolog.Logger
	ch          *amqp.Channel
	postService *service.PostService
}

func NewPostConsumer(logger zerolog.Logger, ch *amqp.Channel, postService *service.PostService) *PostConsumer {
	return &PostConsumer{
		logger:      logger,
		ch:          ch,
		postService: postService,
	}
}
func (c *PostConsumer) PostCreated() {
	queue, err := c.ch.QueueDeclare(pkg.QUEUE_POST_CREATED, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	msgs, err := c.ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	go func() {
		for x := range msgs {
			consumer := new(dto.EventConsumer[dto.EventPostConsumer])
			if err := json.Unmarshal(x.Body, consumer); err != nil {
				c.logger.Error().Err(err).Msg("failed unmarshal to payload")
				return
			}
			err := c.postService.PostCreate(&consumer.Data)
			if err != nil {
				return
			}
		}
	}()
}
func (c *PostConsumer) PostUpdated() {
	queue, err := c.ch.QueueDeclare(pkg.QUEUE_POST_UPDATED, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	msgs, err := c.ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	go func() {
		for x := range msgs {
			consumer := new(dto.EventConsumer[dto.EventPostUpdatedConsumer])
			if err := json.Unmarshal(x.Body, consumer); err != nil {
				c.logger.Error().Err(err).Msg("failed unmarshal to payload")
				return
			}
			err := c.postService.PostUpdate(&consumer.Data)
			if err != nil {
				return
			}
		}
	}()
}
func (c *PostConsumer) LikeTotal() {
	queue, err := c.ch.QueueDeclare(pkg.QUEUE_LIKE_TOTAL, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	msgs, err := c.ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	go func() {
		for x := range msgs {
			consumer := new(dto.EventConsumer[dto.EventLikeTotalConsumer])
			if err := json.Unmarshal(x.Body, consumer); err != nil {
				c.logger.Error().Err(err).Msg("failed unmarshal to payload")
				return
			}
			err := c.postService.PostLike(&consumer.Data)
			if err != nil {
				return
			}
		}
	}()
}
func (c *PostConsumer) CommentTotal() {
	queue, err := c.ch.QueueDeclare(pkg.QUEUE_COMMENT_TOTAL, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	msgs, err := c.ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	go func() {
		for x := range msgs {
			consumer := new(dto.EventConsumer[dto.EventCommentTotalConsumer])
			if err := json.Unmarshal(x.Body, consumer); err != nil {
				c.logger.Error().Err(err).Msg("failed unmarshal to payload")
				return
			}
			err := c.postService.CommentTotal(&consumer.Data)
			if err != nil {
				return
			}
		}
	}()
}
