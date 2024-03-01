package db

import (
	"database/sql"
	"log"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:vuluu2k@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	testQueries = New(conn)
}
