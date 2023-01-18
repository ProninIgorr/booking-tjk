package driver

import (
	"database/sql"

	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxLifetimeDbConn = 5 * time.Minute

// ConnectSQL create database pool for postgresql
func ConnectSQl(dsn string) (*DB, error) {
	d, err := NewDataBase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxLifetimeDbConn)

	dbConn.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

// testDB try to ping database
func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

// NewDataBase create new database for the application
func NewDataBase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
