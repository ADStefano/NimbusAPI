package main

import (
	"context"
	"time"

	"github.com/ADStefano/AmazonHandler/s3handler"
	"github.com/ADStefano/NimbusAPI/internal/api/router"
	"github.com/ADStefano/NimbusAPI/internal/config"
	"github.com/ADStefano/NimbusAPI/internal/database/repository"
	"github.com/ADStefano/NimbusAPI/internal/logger"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	cfg := config.LoadConfig()
	zlog := logger.CreateLogger(cfg.LogLevel)
	defer zlog.Sync()

	cfgS3, err := awsConfig.LoadDefaultConfig(context.Background())
	if err != nil {
		zlog.Error("Erro ao carregar as configurações. (%e)", zap.Error(err))
	}

	client := s3.NewFromConfig(cfgS3)

	handler := s3handler.NewS3Client(client)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	obj, err := handler.ListObjects("adstefano", "", 12, ctx)
	if err != nil {
		zlog.Error("Erro ao listar os objetos. (%e)", zap.Error(err))
	}
	for _, o := range obj {
		zlog.Info("Objeto encontrado", zap.String("key", *o.Key), zap.Int64("tamanho", *o.Size))
	}

	db, err := repository.Conect(cfg.DB_CONN)
	if err != nil {
		zlog.Error("Erro ao conectar com o banco de dados. (%e)", zap.Error(err))
	}

	print(db)
	// db.AutoMigrate(&models.Objects{}, &models.Buckets{}, &models.Executions{})
	router := router.CreateRouter(zlog)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PETECO PETECO PETECO",
		})
	})

	router.Run(cfg.Port)

}
