package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	driverConn = "postgres"
	host       = "localhost"
	port       = 15432
	user       = "root"
	password   = "root"
	dbname     = "go-restaurant-api"
)

var DB *sql.DB

type Row interface {
	Scan(dest ...any) error
}

func init() {
	log.Println("Connecting to the db...")
	var err error
	dataSourceName := dataSourceGenerate()
	DB, err = sql.Open(driverConn, dataSourceName)
	if err != nil {
		log.Println("Database connection error.")
		log.Panicln(err)
	}
	log.Println("Database connection started.")
}

func dataSourceGenerate() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}
