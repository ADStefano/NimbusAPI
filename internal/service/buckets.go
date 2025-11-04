package service

import (
	// "context"
	// "time"

	"github.com/ADStefano/AmazonHandler/s3handler"
	"github.com/ADStefano/NimbusAPI/internal/api/dto"
	"github.com/ADStefano/NimbusAPI/internal/database/repository"
	"go.uber.org/zap"
)

type BucketService struct {
	repository *repository.BucketRepository
	handler    *s3handler.Client
	zlog       *zap.Logger
}

// NewBucketService cria uma nova instância de BucketService
func NewBucketService(repo *repository.BucketRepository, handler *s3handler.Client, zlog *zap.Logger) *BucketService {
	return &BucketService{
		repository: repo,
		handler:    handler,
		zlog:       zlog,
	}
}

func (svr *BucketService) CreateBucket(req dto.BucketRequest) (*dto.BucketResponse, error) {

	var resp dto.BucketResponse

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// _, err := svr.handler.CreateBucket(req.Name, ctx)
	// if err != nil {
	// 	svr.zlog.Error("Erro ao criar bucket no S3", zap.Error(err))
	// 	return nil, err
	// }

	createdBucket, err := svr.repository.CreateBucket(req)
	if err != nil {
		svr.zlog.Error("Erro ao criar bucket no repositório", zap.Error(err))
		return nil, err
	}

	resp.ID = createdBucket.ID
	resp.Name = req.Name

	return &resp, nil
}
