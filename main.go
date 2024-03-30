package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/vuluu2k/simple_bank/api"
	db "github.com/vuluu2k/simple_bank/db/sqlc"
	"github.com/vuluu2k/simple_bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	fmt.Println(config)
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(config.SeverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
