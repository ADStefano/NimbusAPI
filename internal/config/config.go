package config

// TODO TROCAR PARA VIPER
var (
	AppVersion = "development"
	AppName    = "nimbus_api"
	AppPort    = "8080"
	AppEnv     = "development"
	DBHost     = "localhost"
	DBPort     = "5432"
	DBUser     = "root"
	DBPassword = "root"
	DBName     = "nimbus"
	LOG_LEVEL  = "debug"
	DSN = "postgres://root:root@localhost:5432/nimbus"
)
