package config

import "go.uber.org/zap"

type Config struct {
	AppName      string
	AppSecret    string
	AppHost      string
	AppPort      string
	AppEnv       string
	AppVersion   string
	MainDatabase Database
	LogPath      string
	LogFile      string
	Logger       *zap.Logger
}

type Database struct {
	Driver            string
	Host              string
	Port              int
	User              string
	Password          string
	DatabaseName      string
	MaxIdleConnection int
	MaxOpenConnection int
}

func NewConfig() *Config {
	return &Config{
		AppName:    Env("APP_NAME", defaultAppName),
		AppSecret:  Env("APP_SECRET", ""),
		AppHost:    Env("APP_HOST", "localhost"),
		AppPort:    Env("APP_PORT", defaultAppPort),
		AppEnv:     Env("APP_ENV", "local"),
		AppVersion: Env("APP_VERSION", "1.0.0"),
		LogPath:    Env("LOG_PATH", "logs"),
		LogFile:    Env("LOG_FILE", "franchise-ordering-system.log"),
		MainDatabase: Database{
			Driver:            Env("DB_CONNECTION", "mysql"),
			Host:              Env("DB_HOST", ""),
			Port:              EnvAsInt("DB_PORT", defaultDBPort),
			User:              Env("DB_USERNAME", ""),
			Password:          Env("DB_PASSWORD", ""),
			DatabaseName:      Env("DB_DATABASE", ""),
			MaxIdleConnection: EnvAsInt("DB_MAX_IDLE_CONNECTION", defaultDBMaxIdleConnection),
			MaxOpenConnection: EnvAsInt("DB_MAX_OPEN_CONNECTION", defaultDBMaxOpenConnection),
		},
	}
}
