package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(url, database string) (*mongo.Database, error) {
	clOpts := options.Client().ApplyURI(url)

	cl, err := mongo.Connect(context.TODO(), clOpts)
	if err != nil {
		return nil, err
	}

	if err := cl.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	db := cl.Database(database)

	if err := db.Client().Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return db, nil
}
