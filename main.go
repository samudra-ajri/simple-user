package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/samudra-ajri/simple-user/api"
	db "github.com/samudra-ajri/simple-user/db/sqlc"
	"github.com/samudra-ajri/simple-user/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(":" + config.Port)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
