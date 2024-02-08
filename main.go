package main

import (
	"blogger/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("unable to load environment variables")
	}
}

func main() {
	host := os.Getenv("POSTGRESDB_HOST")
	dbname := os.Getenv("POSTGRESDB_DBNAME")
	user := os.Getenv("POSTGRESDB_USER")
	password := os.Getenv("POSTGRESDB_PASSWORD")
	port := os.Getenv("POSTGRESDB_PORT")

	connectionString := "host=" + host + " user=" + user + " password=" + password +
		" dbname=" + dbname + " port=" + port

	conn, err := database.NewConnection(connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}

}
