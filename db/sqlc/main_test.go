package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const dbDriver string = "postgres"
const dbSource string = "postgresql://root:secret@localhost:5432/messaging_app?sslmode=disable"

var testQueries *Queries

func TestMain(t *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Could not connect to database", err)
	}
	testQueries = New(conn)
	os.Exit(t.Run())
}
