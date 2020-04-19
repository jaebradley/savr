package main

import (
	"fmt"

	m "github.com/go-pg/migrations"
)

func init() {
	m.MustRegisterTx(func(db m.DB) error {
		fmt.Println("Creating provider_type enum")
		_, err := db.Exec(
			`CREATE TYPE provider_type AS ENUM (
				'GOOGLE', 
				'FACEBOOK', 
				'TWITTER'
			);`,
		)
		return err
	}, func(db m.DB) error {
		fmt.Println("Dropping provider_type enum")
		_, err := db.Exec("DROP TYPE provider_type;")
		return err
	})
}
