package main

import (
	sv "Messaging/api"
	db "Messaging/db/sqlc"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const dbDriver string = "postgres"
const dbSource string = "postgresql://root:secret@localhost:5432/messaging_app?sslmode=disable"
const address string = "localhost:8080"

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Could not connect to database", err)
	}
	query := db.New(conn)
	server := sv.NewServer(query)
	err = server.Start(address)
	if err != nil {
		log.Fatal("Could not start server", err)
	}
}
