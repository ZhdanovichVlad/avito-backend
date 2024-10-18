package main

import (
	"avitoTest/backend/internal/application"
	"avitoTest/backend/internal/config"
	"avitoTest/backend/internal/handlers/tender"
	"avitoTest/backend/internal/repository"
	tenderUC "avitoTest/backend/internal/usecase/tender"
	"avitoTest/backend/pkg/http/ginrouter"
	"flag"
	"fmt"
)

func main() {
	env := flag.String("env", "", "Specify environment (e.g. 'local')")
	flag.Parse()
	isLocal := false
	fmt.Println(*env)
	// Если передан флаг -env=local, используем .env.local
	if *env == "local" {
		isLocal = true
	}

	conf := config.MustLoad(isLocal)
	storageData := repository.ConnectToStorage(conf, isLocal)
	defer storageData.Close()
	storageData.CreateTables()
	addr := "0.0.0.0:8080"
	tenderUseCase := tenderUC.New(storageData)
	tenderHadnlers := tender.NewTenderHandlers(tenderUseCase)

	router := ginrouter.New()
	app := application.New(router)
	app.RegisterTenderHandlers(tenderHadnlers)
	app.Run(addr)
}
