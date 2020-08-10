package database

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

type UserResource struct {
	ID         uint64
	UserID     uint64
	ResourceID uint64
	CreatedAt  uint64
}

func CreateUserResource(connection *pgx.Conn, uint64, userID uint64, resourceID uint64) (*UserResource, error) {
	var userResource UserResource

	err := connection.QueryRow(`
		INSERT INTO user_resources (user_id, resource_id, created_at)
		VALUES ($1, $2, current_time)
		RETURNING
			"id",
			"user_id",
			"resource_id",
			"created_at"`,
		userID,
		resourceID,
	).Scan(
		&userResource.ID,
		&userResource.UserID,
		&userResource.ResourceID,
		&userResource.CreatedAt,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	return &userResource, nil
}
