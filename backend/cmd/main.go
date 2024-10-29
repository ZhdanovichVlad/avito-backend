package main

import (
	"avitoTest/backend/internal/application"
	"avitoTest/backend/internal/handlers"
	"avitoTest/backend/internal/repository"
	"avitoTest/backend/internal/usecase"
	"avitoTest/backend/pkg/http/ginrouter"
	"log"
	"os"

	"github.com/joho/godotenv"

	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %w", err)
	}
	REGISTRY_DB_DSB := os.Getenv("REGISTRY_DB_DSB")

	db, err := sql.Open("postgres", REGISTRY_DB_DSB)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error during connection verification: %v", err)
	}

	defer db.Close()
	storge := repository.New(db)

	addr := "0.0.0.0:8080"
	tenderUseCase := usecase.New(storge)
	tenderHandlers := handlers.NewTenderPresentation(tenderUseCase)

	router := ginrouter.New()
	app := application.New(router)
	app.RegisterTenderHandlers(tenderHandlers)
	app.Run(addr)
}
