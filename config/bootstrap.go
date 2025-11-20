package config

import (
	"feed-management/delivery/messaging/consumer"
	"feed-management/internal/repository"
	"feed-management/internal/service"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Bootstrap struct {
	DB        *mongo.Database
	Cache     *memcache.Client
	Logger    zerolog.Logger
	Validator *validator.Validate
	Router    *chi.Mux
	Ch        *amqp.Channel
}

func Initialize(deps *Bootstrap) {
	// publisher

	// cache

	// repository
	postRepo := repository.NewPostRepository(deps.DB)

	// service
	postServ := service.NewPostService(deps.Logger, deps.Validator, postRepo)

	// handler

	// middleware

	// routes

	// subcriber
	postSubs := consumer.NewPostConsumer(deps.Logger, deps.Ch, postServ)
	postSubs.PostCreated()
	postSubs.PostUpdated()
	postSubs.LikeTotal()
	postSubs.CommentTotal()
}
