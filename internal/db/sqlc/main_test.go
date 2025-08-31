package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries

var testPool *pgxpool.Pool

const dbURL = "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable"

func TestMain(m *testing.M) {
	var err error
	testPool, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testPool)

	code := m.Run()

	testPool.Close()

	os.Exit(code)
}
