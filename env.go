package env

import (
	"os"
)

// GetEnvOrDefault get enviroment variable based on the key or use the default value
func GetEnvOrDefault(key, def string) string {
	env := os.Getenv(key)
	if env == "" {
		return def
	}
	return env
}

// GetEnvOrPanic get enviroment variable based on the key or panic
func GetEnvOrPanic(key string) string {
	env := os.Getenv(key)
	if env == "" {
		panic("failed to find env for the following key" + key)
	}
	return env
}
