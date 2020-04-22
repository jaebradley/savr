package main

import (
	"fmt"

	m "github.com/go-pg/migrations"
)

func init() {
	m.MustRegisterTx(func(db m.DB) error {
		fmt.Println("Creating user resource tags table")
		_, err := db.Exec(
			`CREATE TABLE user_resource_tags (
				"id" BIGSERIAL PRIMARY KEY,
				"user_resource_id" BIGINT REFERENCES user_resources(id) ON DELETE CASCADE,
				"name" TEXT NOT NULL,
				UNIQUE ("user_resource_id", "name"),
				CONSTRAINT valid_name CHECK (is_valid_hashtag(name))
			)`,
		)
		return err
	}, func(db m.DB) error {
		fmt.Println("Dropping user resource tags table")
		_, err := db.Exec("DROP TABLE user_resource_tags")
		return err
	})
}
