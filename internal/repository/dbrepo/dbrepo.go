package dbrepo

import (
	"database/sql"

	"github.com/ProninIgorr/booking-tjk/internal/config"
	"github.com/ProninIgorr/booking-tjk/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// NewPostgresRepo create new database repository

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DataBaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

// NewTestingRepo create new test database repository
func NewTestingRepo(a *config.AppConfig) repository.DataBaseRepo {
	return &testDBRepo{
		App: a,
	}
}
