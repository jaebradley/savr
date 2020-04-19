package database

import (
	"fmt"
	"net/mail"
	"os"

	"github.com/jackc/pgx"
)

const (
	googleProviderType = "GOOGLE"
)

// User represents an application user
type User struct {
	ID           uint64
	EmailAddress string
}

// CreateUser adds an application user to the database
func CreateUser(emailAddress mail.Address) User {
	conn, err := pgx.Connect(DatabaseConnectionConfiguration)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	defer conn.Close()

	var user User

	err = conn.QueryRow(`
		INSERT INTO users (email_address)
		VALUES ($1) 
		RETURNING "id"`,
		emailAddress.Address,
	).Scan(
		&user.ID,
		&user.EmailAddress,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return user
}

// GetUserByEmailAddress returns a user for a given input email address
func GetUserByEmailAddress(emailAddress mail.Address) User {
	conn, err := pgx.Connect(DatabaseConnectionConfiguration)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	defer conn.Close()

	var user User

	err = conn.QueryRow(`
		SELECT
			id,
			email_address
		FROM users
		WHERE email_address = $1`,
		emailAddress.Address,
	).Scan(
		&user.ID,
		&user.EmailAddress,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return user
}

// UserWithEmailAddressExists checks if a user already exists with the specified email address
func UserWithEmailAddressExists(emailAddress mail.Address) bool {
	return GetUserByEmailAddress(emailAddress) != User{}
}
