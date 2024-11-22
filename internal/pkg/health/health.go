package health

import (
	"net/http"

	"github.com/alexliesenfeld/health"
	"github.com/alexliesenfeld/health/middleware"
	"github.com/gorilla/mux"
)

func InitHealthRoute(r *mux.Router, healthCheckers ...health.CheckerOption) {
	handler := health.NewHandler(
		NewChecker(healthCheckers...),
		health.WithResultWriter(health.NewJSONResultWriter()),
		health.WithMiddleware(middleware.FullDetailsOnQueryParam("full")),
		health.WithStatusCodeUp(http.StatusOK),
		health.WithStatusCodeDown(http.StatusServiceUnavailable),
	)

	r.HandleFunc("/_healthz", handler)
}
