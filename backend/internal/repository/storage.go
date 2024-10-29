package repository

import (
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(db *sql.DB) *Storage {
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
	//err = s.NewVersionStorage()
	//if err != nil {
	//	log.Fatal("failed to create version storage")
	//}
	//err = s.CreateBidsDB()
	//if err != nil {
	//	log.Fatal("failed to create version storage")
	//}
	//
	//err = s.CreateBidStory()
	//if err != nil {
	//	log.Fatal("failed to create bid story storage")
	//}

}
