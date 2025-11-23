package dto

import "time"

type PostResponse struct {
	PostId       int64     `json:"post_id"`
	UserId       int64     `json:"user_id"`
	Description  string    `json:"description"`
	TotalLike    int       `json:"total_like"`
	TotalComment int       `json:"total_comment"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
