package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/vsynclabs/billsoft/pkg/database"
)

func init() {
	// Get the SERVER_MODE from the environment variable
	serverMode := os.Getenv("SERVER_MODE")

	// Check if SERVER_MODE is "dev"  //  SERVER_MODE=dev go run ./cmd
	if serverMode == "dev" {
		// Load the .env file
		if err := godotenv.Load(".env"); err != nil {
			log.Fatalf("Unable to load .env file: %v", err)
		}
		log.Println(".env file loaded successfully.")

	} else if serverMode != "prod" {
		// If SERVER_MODE is neither "dev" nor "prod", show an error and exit
		log.Fatalln("Invalid SERVER_MODE. Please set SERVER_MODE to 'dev' or 'prod'.")
	}
}

func Start() {

	// Database Connection
	conn := NewDatabaseConnection()
	conn.CheckStatus()
	defer conn.Close()
	// Initilize Database
	query := database.NewQuery(conn.db)
	if err := query.InitilizeDatabase(); err != nil {
		log.Fatalf("Unable to initilize database: %v", err)
	}
	// New Server
	server := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: NewRouter(conn).mux,
	}
	log.Printf("Server is running on port %v", os.Getenv("PORT"))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Unable to start the server: %v", err)
	}
}
