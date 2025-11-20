package consumer

import (
	"feed-management/internal/service"
	"feed-management/pkg"
	"feed-management/pkg/dto"
	"fmt"

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
			err := c.postService.Create(&consumer.Data)
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
			fmt.Println(consumer)
		}
	}()
}
func (c *PostConsumer) PostLike() {
	queue, err := c.ch.QueueDeclare(pkg.QUEUE_POST_LIKE, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	msgs, err := c.ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	go func() {
		for x := range msgs {
			consumer := new(dto.EventConsumer[dto.EventPostLikeConsumer])
			if err := json.Unmarshal(x.Body, consumer); err != nil {
				c.logger.Error().Err(err).Msg("failed unmarshal to payload")
				return
			}
			fmt.Println(consumer)
		}
	}()
}
func (c *PostConsumer) PostUnlike() {
	queue, err := c.ch.QueueDeclare(pkg.QUEUE_POST_UNLIKE, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	msgs, err := c.ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	go func() {
		for x := range msgs {
			consumer := new(dto.EventConsumer[dto.EventPostUnlikeConsumer])
			if err := json.Unmarshal(x.Body, consumer); err != nil {
				c.logger.Error().Err(err).Msg("failed unmarshal to payload")
				return
			}
			fmt.Println(consumer)
		}
	}()
}
func (c *PostConsumer) CommentIncrement() {
	queue, err := c.ch.QueueDeclare(pkg.QUEUE_COMMENT_INCREMENT, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	msgs, err := c.ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	go func() {
		for x := range msgs {
			consumer := new(dto.EventConsumer[dto.EventCommentIncrementConsumer])
			if err := json.Unmarshal(x.Body, consumer); err != nil {
				c.logger.Error().Err(err).Msg("failed unmarshal to payload")
				return
			}
			fmt.Println(consumer)
		}
	}()
}
func (c *PostConsumer) CommentDecrement() {
	queue, err := c.ch.QueueDeclare(pkg.QUEUE_COMMENT_DECREMENT, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	msgs, err := c.ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	go func() {
		for x := range msgs {
			consumer := new(dto.EventConsumer[dto.EventCommentDecrementConsumer])
			if err := json.Unmarshal(x.Body, consumer); err != nil {
				c.logger.Error().Err(err).Msg("failed unmarshal to payload")
				return
			}
			fmt.Println(consumer)
		}
	}()
}
