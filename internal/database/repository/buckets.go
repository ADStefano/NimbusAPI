package repository

import "gorm.io/gorm"

type BucketRepository struct {
	db *gorm.DB
}

// NewBucketRepository cria uma nova instância de BucketRepository
func NewBucketRepository(database *gorm.DB) *BucketRepository {
	return &BucketRepository{
		db: database,
	}
}