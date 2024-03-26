package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/vuluu2k/simple_bank/api"
	db "github.com/vuluu2k/simple_bank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:vuluu2k@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
