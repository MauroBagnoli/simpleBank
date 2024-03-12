package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"tutorial.sqlc.dev/app/api"
	db "tutorial.sqlc.dev/app/db/sqlc"
	"tutorial.sqlc.dev/app/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	ctx := context.Background()

	conn, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
