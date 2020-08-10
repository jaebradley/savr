package database

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

// Tag is a identifier that can be applied to a Resource
type Tag struct {
	ID        uint64
	Name      string
	CreatedAt uint64
}

// CreateTag creates a Tag
func CreateTag(connection *pgx.Conn, name string) (*Tag, error) {
	var tag Tag

	err := connection.QueryRow(`
		INSERT INTO tags (name, created_at)
		VALUES ($1, current_time)
		RETURNING
			"id",
			"name",
			"created_at"
	`).Scan(
		&tag.ID,
		&tag.Name,
		&tag.CreatedAt,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	return &tag, nil
}
