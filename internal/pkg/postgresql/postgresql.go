package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

const (
	errCodeViolatesUniqueConstraint = "23505"
)

type Postgres struct {
	db *sql.DB
}

func New(cfg Config) (*Postgres, error) {
	if cfg.GetPGDSN() != "" {
		sql.Register("instrumented-postgres", pq.Driver{})
		dbSQL, err := sql.Open("instrumented-postgres", cfg.GetPGDSN())
		if err != nil {
			return nil, fmt.Errorf("failed to connect database: %w", err)
		}

		dbSQL.SetMaxIdleConns(cfg.GetPGMaxIdleConns())
		dbSQL.SetMaxOpenConns(cfg.GetPGMaxOpenConns())
		dbSQL.SetConnMaxLifetime(cfg.GetPGTimeout())

		return &Postgres{db: dbSQL}, nil
	}

	return nil, nil
}

func (p *Postgres) GetDB() *sql.DB {
	return p.db
}

func (p *Postgres) Alive(ctx context.Context) error {
	return p.db.PingContext(ctx)
}

func IsNotFound(err error) bool {
	return strings.Contains(err.Error(), sql.ErrNoRows.Error())
}

type sqlState interface {
	SQLState() string
}

func ViolatesUniqueConstraint(err error) bool {
	pe, ok := err.(sqlState)
	if !ok {
		return false
	}
	return pe.SQLState() == errCodeViolatesUniqueConstraint
}
