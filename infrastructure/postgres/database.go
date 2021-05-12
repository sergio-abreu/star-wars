package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func NewDatabase(config Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.Database))
	if err != nil {
		return nil, errors.Wrap(err, "failed to open postgres database connection")
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "failed to ping postgres database")
	}
	return db, nil
}
