package controllers

import (
	"github.com/ADStefano/NimbusAPI/internal/api/dto"
	"github.com/ADStefano/NimbusAPI/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BucketController struct {
	service *service.BucketService
}

// NewBucketController cria uma nova instância de BucketController
func NewBucketController(svc *service.BucketService) *BucketController {
	return &BucketController{
		service: svc,
	}
}

func (ctrl *BucketController) CreateBucket(c *gin.Context) {

	var bucket dto.BucketRequest

	zlog := c.MustGet("logger").(*zap.Logger)

	if err := c.ShouldBindJSON(&bucket); err != nil {
		zlog.Error("Erro ao fazer parse dos dados para struct", zap.Error(err))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	zlog.Info("Criando bucket", zap.String("bucketName", bucket.Name))

	createdBucket, err := ctrl.service.CreateBucket(bucket)
	if err != nil {
		zlog.Error("Erro ao criar bucket", zap.Error(err))
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	zlog.Info("Bucket criado com sucesso", zap.String("bucketName", createdBucket.Name))
	c.JSON(201, createdBucket)
}
