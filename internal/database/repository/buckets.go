package repository

import (
	"github.com/ADStefano/NimbusAPI/internal/api/dto"
	"github.com/ADStefano/NimbusAPI/internal/api/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BucketRepository struct {
	db   *gorm.DB
	zlog *zap.Logger
}

// NewBucketRepository cria uma nova instância de BucketRepository
func NewBucketRepository(database *gorm.DB, zlog *zap.Logger) *BucketRepository {
	return &BucketRepository{
		db:   database,
		zlog: zlog,
	}
}

func (repo *BucketRepository) CreateBucket(req dto.BucketRequest) (*models.Buckets, error) {

	var bucket models.Buckets

	bucket.BucketName = req.Name
	bucket.CreatedBy = req.CreatedBy
	bucket.UpdatedBy = req.CreatedBy

	err := repo.db.Create(&bucket).Error
	if err != nil {
		repo.zlog.Error("Erro ao salvar bucket no banco de dados", zap.Error(err))
		return nil, err
	}

	return &bucket, nil
}
