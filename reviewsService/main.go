package main

import (
	"fmt"
	"reviewsService/configs"
	"reviewsService/internal/review/model"
	"reviewsService/internal/review/repository"
	"reviewsService/pkg/db/mongo"
)

func main() {
	c := configs.MongoDB{Port: 27017, Host: "localhost"}
	mongoc, err := mongo.Connect(c)
	if err != nil {
		fmt.Println(mongoc)
	}

	r := repository.NewRepository(mongoc)
	mr := model.Review{ProductID: "14", UserID: "2", Rating: 5, Text: "superpuper"}
	err = r.Create(mr)
	if err != nil {
		fmt.Println(mongoc)
	}
}
