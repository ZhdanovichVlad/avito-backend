package main

import (
	tenderApplication "avitoTest/backend/internal/application/tender"
	"avitoTest/backend/internal/config"
	"avitoTest/backend/internal/domain/shared"
	"avitoTest/backend/internal/infrastructure/postgresdb"
	"avitoTest/backend/internal/presentation/http/handlers/ping"
	"avitoTest/backend/internal/presentation/http/handlers/tenders"
	NewRouter "avitoTest/backend/internal/presentation/router"
	"flag"
	"log"
	"net/http"
	"time"
)

func main() {
	env := flag.String("env", "", "Specify environment (e.g. 'local')")
	flag.Parse()

	isLocal := false

	// Если передан флаг -env=local, используем .env.local
	if *env == "local" {
		isLocal = true
	}

	conf := config.MustLoad(isLocal)
	storageData := postgresdb.ConnectToStorage(conf, isLocal)

	defer storageData.Close()
	storageData.CreateTables()

	//business logic related to user and company verification
	logic := shared.NewSharedDomain(storageData)

	tenderApplication := tenderApplication.Application{storageData, logic}

	router := NewRouter.NewRouter()

	ping.PingController(router)
	tender.AddTenderController(router, tenderApplication)

	addr := "0.0.0.0:8080"

	srv := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("server started on", addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Panic("failed to start server")
	}

	log.Println("stopping server")

}
