package main

import (
	"fmt"

	m "github.com/go-pg/migrations"
)

func init() {
	m.MustRegisterTx(func(db m.DB) error {
		fmt.Println("Creating resources table")
		_, err := db.Exec(
			`CREATE TABLE resources (
				"id" BIGSERIAL PRIMARY KEY,
				"location" TEXT NOT NULL,
				"created_at" DATETIME WITH TIMEZONE NOT NULL,
				UNIQUE ("location")
			)`,
		)
		return err
	}, func(db m.DB) error {
		fmt.Println("Dropping resources table")
		_, err := db.Exec("DROP TABLE resources")
		return err
	})
}
