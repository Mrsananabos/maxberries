package service

import (
	"catalogService/internal/category/model"
	"catalogService/internal/category/repository"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
)

type Service struct {
	repo  repository.Repository
	redis *redis.Client
}

func NewService(r repository.Repository, redis *redis.Client) Service {
	return Service{
		repo:  r,
		redis: redis,
	}
}

func (s Service) GetAll(ctx context.Context) ([]model.Category, error) {
	categories, err := s.repo.GetAll()
	if err == nil {
		categoriesJSON, err := json.Marshal(categories)
		if err != nil {
			log.Printf("failed marshalling categories to JSON: %v", err)
		}

		err = s.redis.Set(ctx, "categories", categoriesJSON, 0).Err()
		if err != nil {
			log.Printf("failed save catefories in Redis: %v", err)
		}
	}

	return categories, err
}

func (s Service) GetById(id int64) (model.Category, error) {
	return s.repo.GetById(id)
}

func (s Service) Create(category model.Category) error {
	if err := category.Validate(); err != nil {
		return err
	}

	return s.repo.Create(category)
}

func (s Service) Update(category model.Category) error {
	if err := category.Validate(); err != nil {
		return err
	}

	return s.repo.Update(category)
}

func (s Service) Delete(id int64) error {
	return s.repo.Delete(id)
}
