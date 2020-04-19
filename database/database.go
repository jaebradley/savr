package database

import (
	"github.com/jackc/pgx"
)

var (
	// DatabaseConnectionConfiguration is the connection configuration for connecting to database
	DatabaseConnectionConfiguration pgx.ConnConfig
)
