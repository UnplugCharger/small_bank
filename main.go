package main

import (
	"database/sql"
	"github.com/UnplugCharger/small_bank/api"
	db "github.com/UnplugCharger/small_bank/db/sqlc"
	"github.com/UnplugCharger/small_bank/utils"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("unable to  not load config files", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start server", err)
	}
}
