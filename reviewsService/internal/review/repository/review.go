package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reviewsService/internal/review/model"
)

type Repository struct {
	MCollection *mongo.Collection
}

func NewRepository(mCollection *mongo.Collection) Repository {
	return Repository{MCollection: mCollection}
}

func (r Repository) GetByProductId(ctx context.Context, id int64) ([]model.Review, error) {
	reviews := []model.Review{}
	filter := bson.D{{"product_id", id}}
	cursor, err := r.MCollection.Find(ctx, filter)
	if err != nil {
		return reviews, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var review model.Review
		if err = cursor.Decode(&review); err != nil {
			return reviews, err
		}
		reviews = append(reviews, review)
	}

	if err = cursor.Err(); err != nil {
		return reviews, err
	}

	return reviews, nil
}

func (r Repository) GetByUserId(ctx context.Context, userId string) ([]model.Review, error) {
	reviews := []model.Review{}
	filter := bson.D{{"user_id", userId}}
	cursor, err := r.MCollection.Find(ctx, filter)
	if err != nil {
		return reviews, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var review model.Review
		if err = cursor.Decode(&review); err != nil {
			return reviews, err
		}
		reviews = append(reviews, review)
	}

	if err = cursor.Err(); err != nil {
		return reviews, err
	}

	return reviews, nil
}

func (r Repository) Create(ctx context.Context, review model.Review) (model.Review, error) {
	bsonReview, err := bson.Marshal(review)
	if err != nil {
		return model.Review{}, fmt.Errorf("error marshal review to bson: %w", err)
	}

	insertResult, err := r.MCollection.InsertOne(ctx, bsonReview)
	if err != nil {
		return model.Review{}, fmt.Errorf("error inserting review into database: %w", err)
	}

	review.ID = insertResult.InsertedID.(primitive.ObjectID)

	return review, nil
}

func (r Repository) DeleteById(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	deleteResult, err := r.MCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("not review with id = %s", id)
	}

	return nil
}

func (r Repository) DeleteByProductId(ctx context.Context, id int64) error {
	filter := bson.D{{"product_id", id}}
	deleteResult, err := r.MCollection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("not reviews for product with id = %s", id)
	}

	return nil
}

func (r Repository) Update(ctx context.Context, id primitive.ObjectID, review model.ContentReview) (err error) {
	bsonReview := bson.M{"rating": review.Rating, "text": review.Text}
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bsonReview}}

	updateResult, err := r.MCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error update review %w", err)
	}

	if updateResult.MatchedCount != 1 {
		return fmt.Errorf("not fount review for update with id = %s", id.Hex())
	}

	return nil
}
