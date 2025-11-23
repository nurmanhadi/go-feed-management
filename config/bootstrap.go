package config

import (
	"feed-management/delivery/messaging/consumer"
	"feed-management/delivery/rest/handler"
	"feed-management/delivery/rest/middleware"
	"feed-management/delivery/rest/routes"
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
	feedServ := service.NewFeedService(deps.Logger, deps.Validator, postRepo)

	// handler
	feedHand := handler.NewFeedHandler(feedServ)

	// middleware
	deps.Router.Use(middleware.ErrorHandler)

	// routes
	r := routes.Router{
		Router:      deps.Router,
		FeedHandler: feedHand,
	}
	r.New()

	// subcriber
	postSubs := consumer.NewPostConsumer(deps.Logger, deps.Ch, postServ)
	postSubs.PostCreated()
	postSubs.PostUpdated()
	postSubs.LikeTotal()
	postSubs.CommentTotal()
}
