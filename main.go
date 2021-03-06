package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/johncortes8906/stamps/database"
	"github.com/johncortes8906/stamps/user"
	"github.com/spf13/viper"
)

const basePath = "/api"

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	database.SetupDatabase()
	user.SetupRoutes(basePath)
	err := http.ListenAndServe(fmt.Sprintf(":%v", viper.Get("SERVER.PORT")), nil)
	if err != nil {
		log.Fatal(err)
	}
}
