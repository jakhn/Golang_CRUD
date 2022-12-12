package db

import (
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/asadbekGo/golang_crud/config"
)

func ConnectionDB(cfg *config.Config) (*sql.DB, error) {

	connect := fmt.Sprintf(
		"host=%s user=%s database=%s password=%s port=%s",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	)

	db, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
