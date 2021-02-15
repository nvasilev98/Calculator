package connection

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func ConnectDB(cfg ConfigDatabase) (*sql.DB, error) {

	postgres, err := sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	))
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database")
	}
	if errPing := postgres.Ping(); errPing != nil {
		return nil, errors.Wrap(errPing, "failed to ping database")
	}
	return postgres, nil
}
