package handler

import (
	"feed-management/internal/service"
	"feed-management/pkg/response"
	"net/http"
)

type FeedHandler struct {
	feedService *service.FeedService
}

func NewFeedHandler(feedService *service.FeedService) *FeedHandler {
	return &FeedHandler{
		feedService: feedService,
	}
}
func (h *FeedHandler) Foryou(w http.ResponseWriter, r *http.Request) {
	result, err := h.feedService.Foryou()
	if err != nil {
		panic(err)
	}
	response.Success(w, 200, result, r.URL.Path)
}
