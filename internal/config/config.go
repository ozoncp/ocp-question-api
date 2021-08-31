package config

import (
	"os"
	"strconv"
	"strings"
)

type DatabaseConfig struct {
	Host         string
	Port         string
	Database     string
	Username     string
	Password     string
	ExternalHost string
	ExternalPort string
}

type KafkaConfig struct {
	Brokers []string
	Topic   string
}

type JaggerConfig struct {
	Host string
	Port string
}

type Config struct {
	Database DatabaseConfig
	Kafka    KafkaConfig
	Jagger   JaggerConfig
}

// NewConfig returns a new Config struct
func NewConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:         getEnv("DB_HOST", ""),
			Port:         getEnv("DB_PORT", ""),
			Database:     getEnv("DB_DATABASE", ""),
			Username:     getEnv("DB_USERNAME", ""),
			Password:     getEnv("DB_PASSWORD", ""),
			ExternalHost: getEnv("DB_EXTERNAL_HOST", ""),
			ExternalPort: getEnv("DB_EXTERNAL_PORT", ""),
		},
		Kafka: KafkaConfig{
			Brokers: getEnvAsSlice("KAFKA_BROKERS", []string{}, ","),
			Topic:   getEnv("KAFKA_TOPIC", ""),
		},
		Jagger: JaggerConfig{
			Host: getEnv("JAGGER_HOST", ""),
			Port: getEnv("JAGGER_PORT", ""),
		},
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

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")

	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
