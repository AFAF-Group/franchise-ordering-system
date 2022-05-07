package config

type Config struct {
	AppName      string
	AppSecret    string
	AppHost      string
	AppPort      string
	AppEnv       string
	AppVersion   string
	MainDatabase Database
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
		AppName:    Env("APP_NAME", defaultAppName, false),
		AppSecret:  Env("APP_SECRET", "", false),
		AppHost:    Env("APP_HOST", "localhost", false),
		AppPort:    Env("APP_PORT", defaultAppPort, false),
		AppEnv:     Env("APP_ENV", "local", false),
		AppVersion: Env("APP_VERSION", "1.0.0", false),
		MainDatabase: Database{
			Driver:            Env("DB_CONNECTION", "mysql", false),
			Host:              Env("DB_HOST", "", false),
			Port:              EnvAsInt("DB_PORT", defaultDBPort, false),
			User:              Env("DB_USERNAME", "", false),
			Password:          Env("DB_PASSWORD", "", false),
			DatabaseName:      Env("DB_DATABASE", "", false),
			MaxIdleConnection: EnvAsInt("DB_MAX_IDLE_CONNECTION", defaultDBMaxIdleConnection, false),
			MaxOpenConnection: EnvAsInt("DB_MAX_OPEN_CONNECTION", defaultDBMaxOpenConnection, false),
		},
	}
}
