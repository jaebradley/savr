package main

import (
	"fmt"

	m "github.com/go-pg/migrations"
)

func init() {
	m.MustRegisterTx(func(db m.DB) error {
		fmt.Println("Creating users table")
		_, err := db.Exec(
			`CREATE TABLE users (
				"id" BIGSERIAL PRIMARY KEY,
				"email_address" TEXT UNIQUE NOT NULL
			)`,
		)
		return err
	}, func(db m.DB) error {
		fmt.Println("Dropping users table")
		_, err := db.Exec("DROP TABLE users")
		return err
	})
}
