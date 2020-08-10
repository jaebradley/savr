package main

import (
	"fmt"

	m "github.com/go-pg/migrations"
)

func init() {
	m.MustRegisterTx(func(db m.DB) error {
		fmt.Println("Creating tags table")
		_, err := db.Exec(
			`CREATE TABLE tags (
				"id" BIGSERIAL PRIMARY KEY,
				"name" TEXT NOT NULL,
				"created_at" DATETIME WITH TIMEZONE NOT NULL,
				UNIQUE ("name"),
				CONSTRAINT valid_name CHECK (is_valid_hashtag(name))
			)`,
		)
		return err
	}, func(db m.DB) error {
		fmt.Println("Dropping tags table")
		_, err := db.Exec("DROP TABLE tags")
		return err
	})
}
