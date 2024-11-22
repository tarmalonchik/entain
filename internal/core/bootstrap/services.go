package bootstrap

import (
	"context"
	"fmt"

	"github.com/alexliesenfeld/health"

	"github.com/tarmalonchik/entain/internal/core/config"
	"github.com/tarmalonchik/entain/internal/core/svc/wallet"
	"github.com/tarmalonchik/entain/internal/pkg/logger"
)

type ServiceContainer struct {
	conf           *config.Config
	clients        *ClientsContainer
	transactionSvc *wallet.Service
	healthCheckers []health.CheckerOption
}

func GetServices(ctx context.Context, conf *config.Config) (*ServiceContainer, error) {
	var (
		err error
		sv  = &ServiceContainer{conf: conf}
	)

	if sv.clients, err = getClients(ctx, conf); err != nil {
		return nil, fmt.Errorf("getting clients: %w", err)
	}

	sv.healthCheckers = sv.getCheckers()

	sv.transactionSvc = wallet.NewService(sv.clients.storage)
	return sv, nil
}

func (sc *ServiceContainer) GetCustomLogger() *logger.Client {
	return sc.clients.logger
}

func (sc *ServiceContainer) getCheckers() []health.CheckerOption {
	return []health.CheckerOption{
		health.WithCheck(health.Check{
			Name:    "database",
			Timeout: sc.conf.Postgres.GetPGTimeout(),
			Check:   sc.clients.db.Alive,
		}),
	}
}
