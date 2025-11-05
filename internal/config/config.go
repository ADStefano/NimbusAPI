package config

import (
	"context"
	"fmt"
	"reflect"

	"github.com/ADStefano/AmazonHandler/s3handler"
	"github.com/ADStefano/NimbusAPI/internal/api/controllers"
	"github.com/ADStefano/NimbusAPI/internal/database/repository"
	"github.com/ADStefano/NimbusAPI/internal/logger"
	"github.com/ADStefano/NimbusAPI/internal/service"
	"github.com/aws/aws-sdk-go-v2/config"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	AppConfig    *AppConfig
	Logger       *zap.Logger
	BucketConfig *BucketConfig
}

type AppConfig struct {
	AppVersion    string `mapstructure:"APP_VERSION"`
	AppName       string `mapstructure:"APP_NAME"`
	AppEnv        string `mapstructure:"APP_ENV"`
	Port          string `mapstructure:"APP_PORT"`
	LogLevel      string `mapstructure:"LOG_LEVEL"`
	DBConn        string `mapstructure:"DB_CONN"`
	AWSLevel      string `mapstructure:"AWS_LEVEL"`
	AppMainBucket string `mapstructure:"APP_MAIN_BUCKET"`
	GinMode       string `mapstructure:"GIN_MODE"`
}

type BucketConfig struct {
	Controller *controllers.BucketController
	Service    *service.BucketService
	Repository *repository.BucketRepository
}

// LoadConfig Carrega as configurações da aplicação e inicializa os componentes necessários
func LoadConfig() *App {

	app := App{}
	app.AppConfig = LoadAppConfig()
	app.Logger = logger.CreateLogger(app.AppConfig.LogLevel)
	defer app.Logger.Sync()

	cfgS3, err := awsConfig.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(app.AppConfig.AWSLevel))
	if err != nil {
		app.Logger.Error("erro ao carregar as configurações.", zap.Error(err))
	}

	client := s3.NewFromConfig(cfgS3)
	handler := s3handler.NewS3Client(client)

	db, err := repository.Conect(app.AppConfig.DBConn)
	if err != nil {
		app.Logger.Error("erro ao conectar com o banco de dados.", zap.Error(err))
	}

	app.BucketConfig = LoadBucketConfig(db, handler, app.Logger)

	return &app
}

// LoadAppConfig Carrega as configurações da aplicação a partir das variáveis de ambiente
func LoadAppConfig() *AppConfig {

	cfg := AppConfig{}

	viper.SetDefault("AWS_LEVEL", "default")
	viper.SetDefault("GIN_MODE", "release")

	// Não lê as chaves sem o bind env ou se não ler um aquivo, por isso o for loop
	viper.AutomaticEnv()

	t := reflect.TypeOf(cfg)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("mapstructure")
		if tag != "" {
			_ = viper.BindEnv(tag)
		}
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Errorf("erro ao decodificar as envs para o struct: %w", err))
	}

	return &cfg

}

// LoadBucketConfig inicializa as configurações do BucketConfig
func LoadBucketConfig(db *gorm.DB, client *s3handler.Client, zlog *zap.Logger) *BucketConfig {

	repository := repository.NewBucketRepository(db, zlog)
	service := service.NewBucketService(repository, client, zlog)
	controller := controllers.NewBucketController(service)

	return &BucketConfig{
		Controller: controller,
		Service:    service,
		Repository: repository,
	}
}
