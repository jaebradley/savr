package database

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

type UserResourceTag struct {
	ID             uint64
	UserResourceID uint64
	TagID          uint64
	CreatedAt      uint64
}

func CreateResourceTag(connection *pgx.Conn, userResourceID uint64, tagID uint64) (*UserResourceTag, error) {
	var userResourceTag UserResourceTag

	err := connection.QueryRow(`
		INSERT INTO user_resource_tags (user_resource_id, tag_id, created_at)
		VALUES ($1, $2, $3, current_timestamp)
		RETURNING
			"id",
			"user_resource_id",
			"tag_id",
			"created_at"`,
		userResourceID,
		tagID,
	).Scan(
		&userResourceTag.ID,
		&userResourceID.UserResourceID,
		&userResourceTag.TagID,
		&userResourceTag.CreatedAt,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	return &userResourceTag, nil
}
