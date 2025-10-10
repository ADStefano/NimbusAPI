package main


import (
	"context"
	"time"

	"github.com/ADStefano/AmazonHandler/s3handler"
	"github.com/ADStefano/NimbusAPI/logger"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/zap"
)

func main() {

	zlog := logger.CreateLogger()
	defer zlog.Sync()

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		zlog.Error("Erro ao carregar as configurações. (%e)", zap.Error(err))
	}

	client := s3.NewFromConfig(cfg)

	handler := s3handler.NewS3Client(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	obj, err := handler.ListObjects("adstefano", "", 12, ctx)
	if err != nil {
		zlog.Error("Erro ao listar os objetos. (%e)", zap.Error(err))
	}
	for _, o := range obj {
		zlog.Info("Objeto encontrado", zap.String("key", *o.Key), zap.Int64("tamanho", *o.Size))
	}

}

