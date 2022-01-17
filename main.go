package main

import (
	"database/sql"
	"log"
	"simplebank/api"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("cannot connect to DB", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err := server.Start("localhost:8080")
	if err != nil {
		log.Fatalf("cannot start server", err)
	}
}
