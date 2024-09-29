package main

import (
	tenderApplication "avitoTest/backend/application/tender"
	"avitoTest/backend/config"
	"avitoTest/backend/domain/shared"
	"avitoTest/backend/infrastructure/postgresdb"
	"avitoTest/backend/interfaces/http/handlers/ping"
	"avitoTest/backend/interfaces/http/handlers/tenders"
	NewRouter "avitoTest/backend/interfaces/router"
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
	logic := shared.NewSharedDonain(storageData)

	tenderApplication := tenderApplication.NewTenderApplication{storageData, logic}

	router := NewRouter.NewRouter()

	ping.PingController(router)
	tenders.AddTenderController(router, tenderApplication)

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
