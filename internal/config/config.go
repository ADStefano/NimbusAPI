package config

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

type AppConfig struct {
	AppVersion string `mapstructure:"APP_VERSION"`
	AppName    string `mapstructure:"APP_NAME"`
	AppEnv     string `mapstructure:"APP_ENV"`
	Port       string `mapstructure:"APP_PORT"`
	LogLevel   string `mapstructure:"LOG_LEVEL"`
	DB_CONN    string `mapstructure:"DB_CONN"`
}

func LoadConfig() *AppConfig {

	cfg := AppConfig{}

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
