package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"

	"avitoTest/backend/internal/application"
	"avitoTest/backend/internal/handlers"
	"avitoTest/backend/internal/repository"
	"avitoTest/backend/internal/usecase"
	"avitoTest/backend/pkg/http/ginrouter"

	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %w", err)
	}

	REGISTRY_DB_DSB := os.Getenv("REGISTRY_DB_DSB")

	db, err := sql.Open("postgres", REGISTRY_DB_DSB)
	defer db.Close()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	err = goose.Up(db, "migrations")
	if err != nil {
		fmt.Println(os.Getwd())
		log.Fatalf("migrations error: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error during connection verification: %v", err)
	}
	storge := repository.New(db)

	addr := os.Getenv("host")
	tenderUseCase := usecase.New(storge)
	tenderController := handlers.NewTenderController(tenderUseCase)

	router := ginrouter.New()
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	app := application.New(router, logger)
	app.RegisterTenderHandlers(tenderController)
	app.Run(addr)
}
