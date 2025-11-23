package algoritm

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func ForYou() mongo.Pipeline {
	pipeline := mongo.Pipeline{
		{
			{Key: "$addFields", Value: bson.D{
				{Key: "hours_since", Value: bson.D{
					{Key: "$divide", Value: bson.A{
						bson.D{{Key: "$subtract", Value: bson.A{time.Now(), "$created_at"}}},
						1000 * 60 * 60,
					}},
				}},
			}},
		},
		{
			{Key: "$addFields", Value: bson.D{
				{Key: "score", Value: bson.D{
					{Key: "$add", Value: bson.A{
						bson.D{{Key: "$multiply", Value: bson.A{"$total_like", 2}}},
						bson.D{{Key: "$multiply", Value: bson.A{"$total_comment", 3}}},
						bson.D{{Key: "$subtract", Value: bson.A{1000, "$hours_since"}}},
					}},
				}},
			}},
		},
		{
			{Key: "$sort", Value: bson.D{
				{Key: "score", Value: -1},
			}},
		},
		{
			{Key: "$limit", Value: 50},
		},
	}
	return pipeline
}
