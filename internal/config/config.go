package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv               string
	CassandraHost        string
	CassandraPort        int
	CassandraKeyspace    string
	CassandraUsername    string
	CassandraPassword    string
	CassandraConsistency string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{}

	cfg.AppEnv = getEnv("APP_ENV", "development")
	cfg.CassandraPort = getEnvAsInt("CASSANDRA_PORT", 9042)
	cfg.CassandraKeyspace = getEnv("CASSANDRA_KEYSPACE", "url_shortener")
	cfg.CassandraUsername = getEnv("CASSANDRA_USER", "")
	cfg.CassandraPassword = getEnv("CASSANDRA_PASSWORD", "")
	cfg.CassandraConsistency = getEnv("CASSANDRA_CONSISTENCY", "QUORUM")
	log.Printf("Configuration loaded: %+v\n", cfg.AppEnv)
	return cfg
}

//Helpers

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}

	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	valStr := getEnv(key, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultVal
}
