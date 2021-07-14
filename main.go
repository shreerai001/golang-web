package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"sigma/config"
)

var dbServer = config.DatabaseServer{}

func Run() {

	var err error = godotenv.Load()

	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("Getting environment variables configuration")
	}

	dbServer.InitializeDatabase(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	// seed.Load(server.DB)

	config.RunHttpServer(os.Getenv("APP_PORT"))

}

func main() {
	Run()
}
