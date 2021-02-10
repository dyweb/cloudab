package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/dyweb/cloudab/pkg/config"
)

var (
	Client *mongo.Client
)

func Initialize(c *config.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(c.MongoURI))
	if err != nil {
		return err
	}
	Client.Database(c.DBName)
	return nil
}
