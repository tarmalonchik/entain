package config

import (
	"time"

	"github.com/tarmalonchik/entain/internal/pkg/config"
	"github.com/tarmalonchik/entain/internal/pkg/postgresql"
	"github.com/tarmalonchik/entain/internal/pkg/routes"
	"github.com/tarmalonchik/entain/internal/pkg/webservice"
)

// Config contains all environment variables
type Config struct {
	Postgres   postgresql.Config
	CORS       routes.CORSConfig
	WebService webservice.Config

	GracefulTimeout time.Duration `envconfig:"CORE_GRACEFUL_TIMEOUT" default:"10s"`
}

func GetConfig() (conf *Config, err error) {
	conf = &Config{}
	err = config.Load(conf)
	return conf, err
}
