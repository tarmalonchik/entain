package bootstrap

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/tarmalonchik/entain/internal/core/config"
	"github.com/tarmalonchik/entain/internal/pkg/logger"
	"github.com/tarmalonchik/entain/internal/pkg/postgresql"
	"github.com/tarmalonchik/entain/internal/pkg/storage"
)

type ClientsContainer struct {
	db      *postgresql.Postgres
	sqlx    *postgresql.SQLXClient
	storage *storage.Model
	logger  *logger.Client
}

func getClients(_ context.Context, conf *config.Config) (*ClientsContainer, error) {
	clients := &ClientsContainer{}

	pgClient, err := postgresql.New(conf.Postgres)
	if err != nil {
		return nil, fmt.Errorf("init pg client: %w", err)
	}
	clients.db = pgClient
	clients.sqlx = postgresql.NewSQLXClient(clients.db.GetDB(), conf.Postgres)

	purl, err := url.Parse(conf.Postgres.GetPGDSN())
	if err != nil {
		return nil, fmt.Errorf("dns url parse: %w", err)
	}

	err = postgresql.RunMigrations(
		clients.db.GetDB(),
		conf.Postgres.GetPGMigrationsPath(),
		strings.Trim(purl.Path, "/"))
	if err != nil {
		return nil, fmt.Errorf("running migrations: %v", err)
	}

	clients.storage = storage.NewModel(clients.sqlx)
	clients.logger = logger.NewLogger()

	return clients, nil
}
