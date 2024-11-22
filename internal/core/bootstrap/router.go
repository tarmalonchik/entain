package bootstrap

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tarmalonchik/entain/internal/core/config"
	"github.com/tarmalonchik/entain/internal/pkg/routes"
)

func GetRouter(conf *config.Config, sc *ServiceContainer, handlers *HandlerContainer) (*mux.Router, error) {
	r := mux.NewRouter()

	r.Use(routes.GetDefaultCors(&conf.CORS))
	routes.InitDefaultRoutes(r, sc.healthCheckers)

	r.HandleFunc("/user/{userId}/balance", handlers.walletHandler.GetUserBalance).Methods(http.MethodGet)
	r.HandleFunc("/user/{userId}/transaction", handlers.walletHandler.UpdateBalance).Methods(http.MethodPost)
	return r, nil
}
