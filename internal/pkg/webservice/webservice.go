package webservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/tarmalonchik/entain/internal/pkg/logger"
)

type WebService struct {
	conf   Config
	router *mux.Router
	logger *logger.Client
}

func NewWebService(conf Config, router *mux.Router, logger *logger.Client) *WebService {
	return &WebService{
		conf:   conf,
		router: router,
		logger: logger,
	}
}

func (s *WebService) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", s.conf.AppHost, s.conf.AppPort),
		Handler:      s.router,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	errC := make(chan error, 1)
	go func() {
		s.logger.Infof("http servers listening %s", fmt.Sprintf("%s:%s", s.conf.AppHost, s.conf.AppPort))
		errC <- server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		logrus.Info("stop http server")

		timeout, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		err := server.Shutdown(timeout)
		if err != nil {
			s.logger.Errorf(err, "failed to shutdown http server")

		}
		return nil
	case err := <-errC:
		return err
	}
}
