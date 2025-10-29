package controllers

import (
	"github.com/ADStefano/NimbusAPI/internal/api/models"
	"github.com/ADStefano/NimbusAPI/internal/service"
	"github.com/gin-gonic/gin"
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

	var bucket models.Buckets

	if err := c.ShouldBindJSON(&bucket); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}
