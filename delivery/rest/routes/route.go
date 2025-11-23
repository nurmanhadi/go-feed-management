package routes

import (
	"feed-management/delivery/rest/handler"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	Router      *chi.Mux
	FeedHandler *handler.FeedHandler
}

func (r *Router) New() {
	r.Router.Route("/api", func(api chi.Router) {
		api.Route("/feed", func(feed chi.Router) {
			feed.Get("/foryou", r.FeedHandler.Foryou)
		})
	})
}
