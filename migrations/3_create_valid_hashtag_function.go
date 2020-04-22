package main

import (
	"fmt"

	m "github.com/go-pg/migrations"
)

// Regex from https://stackoverflow.com/a/42065927/5225575
func init() {
	m.MustRegisterTx(func(db m.DB) error {
		fmt.Println("Creating is_valid_hashtag function")
		_, err := db.Exec(
			`CREATE OR REPLACE FUNCTION is_valid_hashtag(candidate text) RETURNS bool AS $$
				BEGIN
					RETURN candidate ~ '^#\w+$';
				END;
			$$ LANGUAGE plpgsql`,
		)
		return err
	}, func(db m.DB) error {
		fmt.Println("Dropping is_valid_hashtag function")
		_, err := db.Exec("DROP FUNCTION IF EXISTS is_valid_hashtag(text)")
		return err
	})
}
