package repository

import (
	"context"
	"feed-management/internal/entity"
	"feed-management/pkg"
	"feed-management/pkg/algoritm"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PostRepository struct {
	db *mongo.Database
}

func NewPostRepository(db *mongo.Database) *PostRepository {
	return &PostRepository{db: db}
}
func (r *PostRepository) Create(post *entity.Post) error {
	_, err := r.db.Collection(pkg.COLLECTION_POSTS).InsertOne(context.Background(), post)
	if err != nil {
		return err
	}
	return nil
}
func (r *PostRepository) FindOne(postId int64) (*entity.Post, error) {
	filter := bson.D{{Key: "post_id", Value: postId}}
	post := new(entity.Post)
	err := r.db.Collection(pkg.COLLECTION_POSTS).FindOne(context.Background(), filter).Decode(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}
func (r *PostRepository) ReplaceOne(post *entity.Post) error {
	filter := bson.D{{Key: "post_id", Value: post.PostId}}
	_, err := r.db.Collection(pkg.COLLECTION_POSTS).ReplaceOne(context.Background(), filter, post)
	if err != nil {
		return err
	}
	return nil
}
func (r *PostRepository) FindForYou() ([]entity.Post, error) {
	cursor, err := r.db.Collection(pkg.COLLECTION_POSTS).Aggregate(context.Background(), algoritm.ForYou())
	if err != nil {
		return nil, err
	}
	var posts []entity.Post
	err = cursor.All(context.Background(), &posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
