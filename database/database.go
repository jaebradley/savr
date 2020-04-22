package database

import (
	"github.com/jackc/pgx"
)

var (
	// ConnectionConfiguration is the connection configuration for connecting to database
	ConnectionConfiguration pgx.ConnConfig
	// ConnectionPoolConfiguration is the connection configuration for connecting to database
	ConnectionPoolConfiguration pgx.ConnPoolConfig
	// ConnectionPool is connection pool for database
	ConnectionPool *pgx.ConnPool
)
