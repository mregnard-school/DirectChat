package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"server/src/app"
	"server/src/models"
	"log"
	"os"
)

func main() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	models.Open(username, password, dbName, dbHost)

	a := app.Application{}
	a.Initialize()
	//router.NotFoundHandler = app.NotFoundHandler
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")
	address := fmt.Sprintf("%s:%s", addr, port)
	log.Print("Server running at ", address)
	a.Run(address)
}
