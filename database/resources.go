package database

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

// Resource represents a resource that's saved
type Resource struct {
	ID        uint64
	Location  string
	CreatedAt uint64
}

// CreateResource creates a resource to be saved
func CreateResource(connection *pgx.Conn, location string) (*Resource, error) {
	var resource Resource

	err := connection.QueryRow(`
		INSERT INTO resources (location, created_at)
		VALUES ($1, current_timestamp)
		RETURNING 
			"id",
			"location",
			"created_at"
	`).Scan(
		&resource.ID,
		&resource.Location,
		&resource.CreatedAt,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	return &resource, nil
}
