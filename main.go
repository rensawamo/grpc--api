package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // postgresに接続するために必須
	"github.com/rensawamo/grpc-api/api"
	db "github.com/rensawamo/grpc-api/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)


// postgres 接続
func main() {
	conn, err := sql.Open(dbDriver,dbSource)
	if err != nil {
		log.Fatal("cannot connet to db:", err) //ログのメッセージを残し プログラム終了までもっていく
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot connect to server:", err)
	}
}