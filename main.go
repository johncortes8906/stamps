package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/johncortes8906/stamps/database"
	"github.com/johncortes8906/stamps/user"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	user.SetupRoutes(basePath)
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
