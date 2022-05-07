package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

const defaultConfigFile = "config.env"

func Env(key, defaultValue string) string {
	if strVal, ok := os.LookupEnv(key); ok {
		return strVal
	}
	return defaultValue
}

func EnvAsInt(key string, defaultVal int) int {
	strVal := Env(key, "")
	if val, err := strconv.Atoi(strVal); err == nil {
		return val
	}
	return defaultVal
}

func LoadEnv(file string) {
	once.Do(func() {
		if file == "" {
			file = defaultConfigFile
		}
		if err := godotenv.Load(file); err != nil {
			log.Fatalf("Error loading %s file", file)
		}
	})
}
