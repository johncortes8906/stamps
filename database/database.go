package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

//DbConn variable used to setup the database
var DbConn *sql.DB

//SetupDatabase sets the DB configuration
func SetupDatabase() {
	var err error
	fmt.Println(viper.Get("DB.HOST"))
	config := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", viper.Get("DB.USER"), viper.Get("DB.PASSWORD"), viper.Get("DB.HOST"), viper.Get("DB.PORT"), viper.Get("DB.NAME"))
	DbConn, err = sql.Open("mysql", config)
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(4)
	DbConn.SetMaxIdleConns(4)
	DbConn.SetConnMaxLifetime(60 * time.Second)

}
