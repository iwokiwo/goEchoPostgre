package main

import (
	"log"
	"os"
	"hotels/db"
	"hotels/routes"
	"github.com/joho/godotenv"
)

func main() {
	db.Init()
	err := godotenv.Load(".env")

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	appPort := os.Getenv("APP_PORT")
	e := routes.Init()

	e.Logger.Fatal(e.Start(":"+appPort))
}
