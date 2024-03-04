package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"gomarketplace/api"
	db "gomarketplace/db/sqlc"
	"gomarketplace/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error connection to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
}
