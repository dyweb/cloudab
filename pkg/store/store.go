package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/dyweb/cloudab/pkg/config"
)

var (
	cli *mongo.Client
	DB  *mongo.Database
)

func Initialize(c *config.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(c.MongoURI).SetMinPoolSize(2).SetMaxPoolSize(2))
	if err != nil {
		return err
	}

	// The Client.Ping method can be used to verify that the deployment is successfully connected and the
	// Client was correctly configured.
	if err := cli.Ping(ctx, nil); err != nil {
		return err
	}
	DB = cli.Database(c.DBName)
	return nil
}

func Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return cli.Disconnect(ctx)
}
