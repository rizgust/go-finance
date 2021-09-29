package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/rizgust/go-finance/api"
	db "github.com/rizgust/go-finance/providers/db/sqlc"
	"github.com/rizgust/go-finance/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
