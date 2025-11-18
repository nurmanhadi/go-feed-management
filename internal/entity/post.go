package entity

import "time"

type Post struct {
	PostId       int64     `bson:"post_id"`
	UserId       int64     `bson:"user_id"`
	Description  string    `bson:"description"`
	TotalLike    int       `bson:"total_like"`
	TotalComment int       `bson:"total_comment"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}
