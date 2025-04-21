package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reviewsService/configs"
)

func Connect(cnf configs.MongoDB) (*mongo.Client, error) {
	url := fmt.Sprintf("mongodb://%s:%d", cnf.Host, cnf.Port)
	clientOptions := options.Client().ApplyURI(url)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func ConnectToCollection(cnf configs.MongoDB) (*mongo.Collection, error) {
	client, err := Connect(cnf)
	if err != nil {
		return nil, err
	}

	return client.Database(cnf.Name).Collection(cnf.Collection), nil
}
