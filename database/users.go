package database

import (
	"fmt"
	"net/mail"
	"net/url"
	"os"
	"time"
)

const (
	googleProviderType = "GOOGLE"
)

// User represents an application user
type User struct {
	ID           uint64
	EmailAddress string
}

// UserResource represents a resource associated with the specified User
type UserResource struct {
	ID        uint64
	UserID    uint64
	Location  string
	CreatedAt uint64
}

// CreateUser adds an application user to the database
func CreateUser(emailAddress *mail.Address) User {
	conn, err := ConnectionPool.Acquire()

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
func GetUserByEmailAddress(emailAddress *mail.Address) User {
	conn, err := ConnectionPool.Acquire()

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
func UserWithEmailAddressExists(emailAddress *mail.Address) bool {
	return GetUserByEmailAddress(emailAddress) != User{}
}

// GetUserByID gets a User by it's id
func GetUserByID(id uint64) User {
	conn, err := ConnectionPool.Acquire()

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
		WHERE id = $1`,
		id,
	).Scan(
		&user.ID,
		&user.EmailAddress,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return user
}

// CreateUserResource creates a URL associated with a user
func CreateUserResource(user *User, url *url.URL) (*UserResource, error) {
	conn, err := ConnectionPool.Acquire()

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database : %v\n", err)
	}

	defer conn.Close()

	var userResource UserResource

	err = conn.QueryRow(`
		INSERT INTO user_resources (user_id, location, created_at)
		VALUES ($1, $2, $3)
		RETURNING
			"id",
			"user_id",
			"location",
			"created_at"`,
		user.ID,
		url.String(),
		time.Now().UTC(),
	).Scan(
		&userResource.ID,
		&userResource.UserID,
		&userResource.Location,
		&userResource.CreatedAt,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed : %v\n", err)
	}

	return &userResource, nil
}

// GetUserResources returns a list of resources for the specified user
func GetUserResources(user *User) (*[]UserResource, error) {
	conn, err := ConnectionPool.Acquire()

	defer conn.Close()

	userResources := []UserResource{}

	rows, err := conn.Query(`
		SELECT
			id,
			user_id,
			location,
			created_at
		FROM user_resources
		WHERE user_id = $1`,
		user.ID,
	)

	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var userResource UserResource
		err = rows.Scan(
			&userResource.ID,
			&userResource.UserID,
			&userResource.Location,
			&userResource.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		userResources = append(userResources, userResource)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return &userResources, nil
}
