package main

import (
	"context"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/vkidmode/server-core/pkg/core"

	"github.com/tarmalonchik/entain/internal/core/bootstrap"
	"github.com/tarmalonchik/entain/internal/core/config"
	"github.com/tarmalonchik/entain/internal/pkg/version"
	"github.com/tarmalonchik/entain/internal/pkg/webservice"
)

func init() {
	if version.Service == "" {
		version.Service = "core"
	}
}

func main() {
	ctx := context.Background()

	runtime.GOMAXPROCS(runtime.NumCPU())

	conf, err := config.GetConfig()
	if err != nil {
		logrus.Errorf("failed to load environment: %v", err)
		return
	}

	services, err := bootstrap.GetServices(ctx, conf)
	if err != nil {
		logrus.Errorf("failed to initiate services: %v", err)
		return
	}

	handlers := bootstrap.GetHandlers(services)
	router, err := bootstrap.GetRouter(conf, services, handlers)
	if err != nil {
		logrus.Errorf("failed to initiate router")
		return
	}

	ws := webservice.NewWebService(conf.WebService, router, services.GetCustomLogger())

	app := core.NewCore(services.GetCustomLogger(), conf.GracefulTimeout, 10)

	app.AddRunner(ws.Run, true)

	err = app.Launch(ctx)
	if err != nil {
		logrus.Errorf("application error: %v", err)
	}
}
