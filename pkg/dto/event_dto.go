package dto

import "time"

type EventConsumer[T any] struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
	Data      T         `json:"data"`
}
type EventPostConsumer struct {
	PostId       int64     `json:"post_id"`
	UserId       int64     `json:"user_id"`
	Description  string    `json:"description"`
	TotalLike    int       `json:"total_like"`
	TotalComment int       `json:"total_comment"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type EventPostUpdatedConsumer struct {
	PostId       int64     `json:"post_id"`
	UserId       int64     `json:"user_id"`
	Description  string    `json:"description"`
	TotalLike    int       `json:"total_like"`
	TotalComment int       `json:"total_comment"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type EventPostLikeConsumer struct {
	PostId int64 `json:"post_id"`
	Total  int   `json:"total"`
}
type EventPostUnlikeConsumer struct {
	PostId int64 `json:"post_id"`
	Total  int   `json:"total"`
}
type EventCommentIncrementConsumer struct {
	PostId int64 `json:"post_id"`
	Total  int   `json:"total"`
}
type EventCommentDecrementConsumer struct {
	PostId int64 `json:"post_id"`
	Total  int   `json:"total"`
}
