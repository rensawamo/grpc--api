package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // postgresに接続するために必須
	"github.com/rensawamo/grpc-api/api"
	db "github.com/rensawamo/grpc-api/db/sqlc"
	 "github.com/rensawamo/grpc-api/util"

)

const ()

// postgres 接続
func main() {

	config, err :=  util.LoadConfig(".") //カレント指定
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connet to db:", err) //ログのメッセージを残し プログラム終了までもっていく
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot connect to server:", err)
	}
}
