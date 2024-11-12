package main

import (
	"fmt"

	"github.com/damedelion/rest_wallet/config"
	"github.com/damedelion/rest_wallet/internal/server"
	"github.com/damedelion/rest_wallet/pkg/db/postgres"
	"github.com/gorilla/mux"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		fmt.Printf("failed to get config, err: %v\n", err)
		return
	}

	db, err := postgres.NewDBOpen(&config.DB)
	if err != nil {
		fmt.Printf("failed to open db connection, err: %v\n", err)
		return
	}
	defer db.Close()

	server := server.NewServer(config, db, mux.NewRouter())
	server.Run()
}
