package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	db "tutorial.sqlc.dev/app/db/sqlc"
)

var Queries *db.Queries
var DB *pgxpool.Pool

func main() {
	var err error
	ctx := context.Background()

	// TODO: Extract postgress conection string to reuse in main_test
	DB, err = pgxpool.New(ctx, "postgresql://root:admin@localhost:5432/simple_bank")
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	Queries = db.New(DB)
}
