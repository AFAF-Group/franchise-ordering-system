package config

import "time"

const (
	defaultAppName = "Franchise Ordering System"
	defaultAppPort = "8080"

	// default DB
	defaultDBPort              = 3306
	defaultDBMaxIdleConnection = 25
	defaultDBMaxOpenConnection = 25
	defaultDBMaxLifeConnection = 5 * time.Minute
)
