package postgresdb

import (
	"avitoTest/backend/internal/config"
	"fmt"
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func ConnectToStorage(config *config.Config, isLocal bool) *Storage {
	var connStr string
	if isLocal {
		connStr = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
			config.PorstgresUserName, config.PorstgresPassword, config.PorstgresDatabase)
	} else {
		connStr = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=require",
			config.PorstgresUserName, config.PorstgresPassword, config.PorstgresHost, config.PorstgresPort, config.PorstgresDatabase)
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error during connection verification: %v", err)
	}

	log.Println("Connection to the database was successful")
	return &Storage{db: db}
}

// Close connection to BD
func (s *Storage) Close() {
	s.db.Close()
}

// CreateTables method for adding the required databases
func (s *Storage) CreateTables() {
	err := s.NewTenderStorage()
	if err != nil {
		log.Fatal("failed to create tender storage")
	}
	err = s.NewVersionStorage()
	if err != nil {
		log.Fatal("failed to create version storage")
	}
	err = s.CreateBidsDB()
	if err != nil {
		log.Fatal("failed to create version storage")
	}

	err = s.CreateBidStory()
	if err != nil {
		log.Fatal("failed to create bid story storage")
	}

}
