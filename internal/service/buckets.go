package service

import (
	"github.com/ADStefano/AmazonHandler/s3handler"
	"github.com/ADStefano/NimbusAPI/internal/database/repository"
)

type BucketService struct {
	repository *repository.BucketRepository
	handler    *s3handler.Client
}

// NewBucketService cria uma nova instância de BucketService
func NewBucketService(repo *repository.BucketRepository, handler *s3handler.Client) *BucketService {
	return &BucketService{
		repository: repo,
		handler: handler,
	}
}
