package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
)

// Describing the Config structure
type Config struct {
	ServerAddress    string `env:"SERVER_ADDRESS"`
	PostgresConn     string `env:"POSTGRES_CONN"`
	PostgresJdbcUrl  string `env:"POSTGRES_JDBC_URL"`
	PostgresUserName string `env:"POSTGRES_USERNAME"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     string `env:"POSTGRES_PORT"`
	PostgresDatabase string `env:"POSTGRES_DATABASE"`
}

// MustLoad function loads environment variables. islocal means whether the variables are stored inside the project
func MustLoad(isLocal bool) *Config {
	var cfg Config

	if isLocal {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: %w", err)
		}
	}
	// Load environment variables into the structure
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatalf("Error while reading environment variables::", err)
		return nil
	}

	return &cfg
}
