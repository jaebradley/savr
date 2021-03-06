package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx"
	"github.com/jaebradley/savr/authentication"
	"github.com/jaebradley/savr/database"
	"github.com/jaebradley/savr/graphql"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
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
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse connection string %v\n", err)
		return
	}

	database.ConnectionConfiguration = dbConnConfig
	poolConfig := pgx.ConnPoolConfig{
		ConnConfig: dbConnConfig,
	}
	database.ConnectionPool, err = pgx.NewConnPool(poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool %v\n", err)
		return
	}
}

func main() {
	defer database.ConnectionPool.Close()

	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("WEB_APPLICATION_DOMAIN")},
		AllowCredentials: true,
		Debug:            true,
	})

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
		},
		Extractor: authentication.FromCookie("access_token"),
	})

	router.Handle("/authentication/google", http.HandlerFunc(authentication.GoogleAuthenticationHandler)).Methods(http.MethodPost)
	router.Handle("/graphql", jwtMiddleware.Handler(graphql.Handler)).Methods(http.MethodPost)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	http.ListenAndServe(fmt.Sprint(":", os.Getenv("PORT")), c.Handler(loggedRouter))
}
