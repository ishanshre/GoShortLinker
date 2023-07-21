package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	GetUrlCollections() *mongo.Collection
}

type DB struct {
	Client *mongo.Client
}

func NewConnection(dsn string) (Database, error) {
	client, err := newDatabase(dsn)
	if err != nil {
		return nil, err
	}
	return &DB{
		Client: client,
	}, nil
}

func (db *DB) GetUrlCollections() *mongo.Collection {
	return db.Client.Database("myDB").Collection("urls")
}

func newDatabase(dsn string) (*mongo.Client, error) {
	// Cancel this funxtion request if it takes more than 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// config the mongo client connection configuration with data source name
	clientOptions := options.Client().ApplyURI(dsn)

	// connect to the mongo db with the client options
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// test if the connection is alive/active
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return client, nil
}
