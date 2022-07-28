package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/diyliv/manga-service/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connTimeout     = 30 * time.Second
	maxConnIdleTime = 3 * time.Minute
	minPoolSize     = 20
	maxPoolSize     = 300
)

func InitConnect(config configs.Config) *mongo.Client {
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s",
		// username,
		// password,
		config.Mongo.Host,
		config.Mongo.Port)).
		SetConnectTimeout(connTimeout).
		SetMaxConnIdleTime(maxConnIdleTime).
		SetMinPoolSize(minPoolSize).
		SetMaxPoolSize(maxPoolSize)

	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatalf("Error while connecting to mongo DB: %v\n", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Error while pinging mongo DB: %v\n", err)
	}

	return client
}
