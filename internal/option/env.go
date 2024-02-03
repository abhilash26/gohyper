package option

import (
	"log"
	"os"
	"strconv"
	"time"
)

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Environment variable %s not set, using default: %s", key, defaultValue)
		return defaultValue
	}
	return value
}

func GetIntEnv(key string, defaultValue int) int {
	value, err := strconv.Atoi(getEnv(key, ""))
	if err != nil {
		log.Printf("Invalid value for %s, using default: %d", key, defaultValue)
		return defaultValue
	}
	return value
}

func GetStringEnv(key string, defaultValue string) string {
	return getEnv(key, defaultValue)
}

func GetDurationEnv(key, defaultValue string) time.Duration {
	value, err := time.ParseDuration(getEnv(key, ""))
	if err != nil {
		log.Printf("Invalid value for %s, using default: %s", key, defaultValue)
		duration, _ := time.ParseDuration(defaultValue)
		return duration
	}
	return value
}
