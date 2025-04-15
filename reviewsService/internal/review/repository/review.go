package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"reviewsService/internal/review/model"
)

const REVIEWS_COLLECETION_NAME = "reviews"

type Repository struct {
	DB *mongo.Client
}

func NewRepository(db *mongo.Client) Repository {
	return Repository{DB: db}
}

func (r Repository) GetAll() ([]model.Review, error) {
	// Выбор коллекции
	collection := r.DB.Database("your_database_name").Collection(REVIEWS_COLLECETION_NAME)

	// Получение всех записей из коллекции
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	// Перебор записей
	var reviews []model.Review
	for cursor.Next(context.TODO()) {
		var review model.Review
		if err := cursor.Decode(&review); err != nil {
			log.Fatal(err)
		}
		reviews = append(reviews, review)
	}

	if err = cursor.Err(); err != nil {
		return []model.Review{}, err
	}

	// Вывод всех записей
	for _, review := range reviews {
		fmt.Printf("ID: %s, ProductID: %s, UserID: %s, Rating: %d, Text: %s\n",
			review.ID.Hex(), review.ProductID, review.UserID, review.Rating, review.Text)
	}

	return reviews, nil
}

func (r Repository) GetByProductId(id string) ([]model.Review, error) {
	collection := r.DB.Database("your_database_name").Collection(REVIEWS_COLLECETION_NAME)

	filter := bson.D{{"product_id", id}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return []model.Review{}, err
	}
	defer cursor.Close(context.TODO())

	// Перебор записей
	var reviews []model.Review
	for cursor.Next(context.TODO()) {
		var review model.Review
		if err = cursor.Decode(&review); err != nil {
			return []model.Review{}, err
		}
		reviews = append(reviews, review)
	}

	// Проверка на ошибки после перебора
	if err = cursor.Err(); err != nil {
		return []model.Review{}, err
	}

	// Вывод найденных записей
	if len(reviews) == 0 {
		fmt.Println("No reviews found for ProductID:", id)
	} else {
		for _, review := range reviews {
			fmt.Printf("ID: %s, ProductID: %s, UserID: %s, Rating: %d, Text: %s\n",
				review.ID.Hex(), review.ProductID, review.UserID, review.Rating, review.Text)
		}
	}

	return reviews, nil
}

func (r Repository) Create(review model.Review) error {
	collection := r.DB.Database("receiptsdb").Collection("receipts")

	bsonReview, err := bson.Marshal(review)
	if err != nil {
		return fmt.Errorf("error marshal review to bson %w", err)
	}

	_, err = collection.InsertOne(context.TODO(), bsonReview)
	return err
}

func (r Repository) DeleteById(id string) error {
	collection := r.DB.Database("your_database_name").Collection("reviews")

	filter := bson.D{{"_id", id}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("not review with id = %d", id)
	}

	return nil
}

func (r Repository) DeleteByProductId(id string) error {
	collection := r.DB.Database("your_database_name").Collection("reviews")

	filter := bson.D{{"product_id", id}}
	deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("not reviews for product with id = %d", id)
	}

	return nil
}

func (r Repository) Update(review model.Review) (err error) {
	collection := r.DB.Database("your_database_name").Collection("categories")

	bsonReview, err := bson.Marshal(review)
	if err != nil {
		return fmt.Errorf("error marshal review to bson %w", err)
	}

	filter := bson.D{{"_id", review.ID}}
	update := bson.D{{"$set", bsonReview}}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("error update review %w", err)
	}

	if updateResult.MatchedCount != 1 {
		return fmt.Errorf("not fount review for update with id = %d", review.ID)
	}

	return nil
}
