package database

import (
	"database/sql"
	"log"
	"time"
)

//DbConn variable used to setup the database
var DbConn *sql.DB

//SetupDatabase sets the DB configuration
func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/stamps")
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(4)
	DbConn.SetMaxIdleConns(4)
	DbConn.SetConnMaxLifetime(60 * time.Second)

}
