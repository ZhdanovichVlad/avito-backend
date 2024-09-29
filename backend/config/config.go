package config

import (
	"github.com/joho/godotenv"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

// Describing the Config structure
type Config struct {
	ServerAddress     string `env:"SERVER_ADDRESS"`
	PorstgresConn     string `env:"POSTGRES_CONN"`
	PorstgresJdbcUrl  string `env:"POSTGRES_JDBC_URL"`
	PorstgresUserName string `env:"POSTGRES_USERNAME"`
	PorstgresPassword string `env:"POSTGRES_PASSWORD"`
	PorstgresHost     string `env:"POSTGRES_HOST"`
	PorstgresPort     string `env:"POSTGRES_PORT"`
	PorstgresDatabase string `env:"POSTGRES_DATABASE"`
}

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
