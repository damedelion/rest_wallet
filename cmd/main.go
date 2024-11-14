package main

import (
	"fmt"

	"github.com/damedelion/rest_wallet/config"
	"github.com/damedelion/rest_wallet/internal/server"
	"github.com/damedelion/rest_wallet/internal/sql"
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
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec(sql.CreateTableQuery)
	if err != nil {
		fmt.Printf("failed to create db schema, err: %v\n", err)
		return
	}

	server := server.NewServer(&config.Server, db, mux.NewRouter())
	server.Run()
}
