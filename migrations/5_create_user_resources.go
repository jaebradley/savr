package main

import (
	"fmt"

	m "github.com/go-pg/migrations"
)

func init() {
	m.MustRegisterTx(func(db m.DB) error {
		fmt.Println("Creating user resources table")
		_, err := db.Exec(
			`CREATE TABLE user_resources (
				"id" BIGSERIAL PRIMARY KEY,
				"user_id" BIGINT NOT NULL,
				"resource_id" BIGINT NOT NULL,
				"created_at" DATETIME WITH TIMEZONE NOT NULL,
				UNIQUE ("user_id", "resource_id")
			)`,
		)
		return err
	}, func(db m.DB) error {
		fmt.Println("Dropping user resources table")
		_, err := db.Exec("DROP TABLE user_resources")
		return err
	})
}
