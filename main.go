package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx"
	"github.com/jaebradley/savr/authentication"
	"github.com/jaebradley/savr/database"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	dbConnConfig, err := pgx.ParseConnectionString(os.Getenv("DATABASE_URL"))
	database.DatabaseConnectionConfiguration = dbConnConfig
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse connection string %v\n", err)
	}
}

func main() {
	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("WEB_APPLICATION_DOMAIN")},
		AllowCredentials: true,
		Debug:            true,
	})

	router.HandleFunc("/authentication/google", authentication.GoogleAuthenticationHandler).Methods("POST")

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	http.ListenAndServe(fmt.Sprint(":", os.Getenv("PORT")), c.Handler(loggedRouter))
}
