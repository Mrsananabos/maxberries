package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"reviewsService/configs"
)

func Connect(cnf configs.MongoDB) (*mongo.Client, error) {
	url := fmt.Sprintf("mongodb://%s:%d", cnf.Host, cnf.Port)
	fmt.Println(url)
	clientOptions := options.Client().ApplyURI(url)

	//"mongodb://localhost:27017"
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Все гуд с монго")

	return client, nil
}

func initMongoDB() (*mongo.Client, *mongo.Collection, error) {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Println("MONGO_URI не задана, кэширование отключено.")
		return nil, nil, nil
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Printf("Ошибка подключения к MongoDB: %v", err)
		return nil, nil, err
	}

	// Проверяем соединение (ping)
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Printf("Ошибка ping MongoDB: %v", err)
		return nil, nil, err
	}

	collection := client.Database("receiptsdb").Collection("receipts")
	log.Println("Успешное подключение к MongoDB!")
	return client, collection, nil
}
