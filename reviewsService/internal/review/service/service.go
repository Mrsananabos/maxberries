package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reviewsService/internal/review/model"
	"reviewsService/internal/review/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}

func (s Service) GetByProductId(ctx context.Context, id string) ([]model.Review, error) {
	reviews, err := s.repo.GetByProductId(ctx, id)
	return reviews, err
}

func (s Service) Create(ctx context.Context, review model.Review) error {
	return s.repo.Create(ctx, review)
}

func (s Service) Update(ctx context.Context, id primitive.ObjectID, review model.ContentReview) error {
	if err := review.Validate(); err != nil {
		return err
	}

	return s.repo.Update(ctx, id, review)
}

func (s Service) DeleteByProductId(ctx context.Context, id string) error {
	return s.repo.DeleteByProductId(ctx, id)
}

func (s Service) DeleteById(ctx context.Context, id primitive.ObjectID) error {
	return s.repo.DeleteById(ctx, id)
}
