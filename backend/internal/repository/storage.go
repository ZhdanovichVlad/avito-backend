package repository

import (
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
