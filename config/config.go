package config

import (
	"os"
	"strconv"
)

type EnvVar struct {
	UIText string
}

type Config struct {
	Env    EnvVar
	Number int
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		Env: EnvVar{
			UIText: getEnv("ENV_VAR_STRUCT", "No env in struct"),
		},
		Number: getEnvAsInt("MAGIC_NUMBER", 0),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
