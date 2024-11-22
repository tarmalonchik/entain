package routes

import (
	"net/http"
	"net/url"
	"strings"

	alexHealth "github.com/alexliesenfeld/health"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/tarmalonchik/entain/internal/pkg/health"
)

func InitDefaultRoutes(r *mux.Router, healthCheckers []alexHealth.CheckerOption) {
	health.InitHealthRoute(r, healthCheckers...)
}

const (
	originRequestKey = "Origin"
	originKey        = "Access-Control-Allow-Origin"
	methodsKey       = "Access-Control-Allow-Methods"
	headersKey       = "Access-Control-Allow-Headers"
)

type CORSConfig struct {
	CORSOrigin          string   `envconfig:"CORS_DEFAULT_ORIGIN" default:"*"`
	CORSAllowedReferrer []string `envconfig:"CORS_ALLOWED_REFERER" default:""`
	CORSMethods         string   `envconfig:"CORS_DEFAULT_METHODS" default:"POST, GET, OPTIONS, PUT, DELETE"`
	CORSHeaders         string   `envconfig:"CORS_DEFAULT_HEADERS" default:"Accept, Content-Type, Content-Length, Accept-Encoding"`
}

func (c *CORSConfig) GetCORSOrigin() string            { return c.CORSOrigin }
func (c *CORSConfig) GetCORSAllowedReferrer() []string { return c.CORSAllowedReferrer }
func (c *CORSConfig) GetCORSMethods() string           { return c.CORSMethods }
func (c *CORSConfig) GetCORSHeaders() string           { return c.CORSHeaders }

type configImpl interface {
	GetCORSOrigin() string
	GetCORSAllowedReferrer() []string
	GetCORSMethods() string
	GetCORSHeaders() string
}

func GetDefaultCors(conf configImpl) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var origin string

			if requestOrigin, err := url.Parse(r.Header.Get(originRequestKey)); err == nil {
				for _, domain := range conf.GetCORSAllowedReferrer() {
					if strings.HasSuffix(requestOrigin.Host, domain) {
						origin = requestOrigin.String()
						break
					}
				}
			} else {
				logrus.WithError(err).Warn("can't parse request origin")
			}

			if origin == "" {
				origin = conf.GetCORSOrigin()
			}

			w.Header().Set(originKey, origin)
			w.Header().Set(methodsKey, conf.GetCORSMethods())
			w.Header().Set(headersKey, conf.GetCORSHeaders())
			if r.Method == http.MethodOptions {
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
