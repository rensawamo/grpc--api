package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" // postgresに接続するために必須
	"github.com/rensawamo/grpc-api/util"
)

var testQueries *Queries
var testDB *sql.DB

// func TestMain(m *testing.M) {
// 	conn, err := sql.Open(dbDriver, dbSource)
// 	if err != nil {
// 		log.Fatal("cannot connect to db:", err)
// 	}
// 	testQueries = New(conn)
// 	os.Exit(m.Run()) //テストの実行
// }

// トランザクション用
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run()) //テストの実行
}
