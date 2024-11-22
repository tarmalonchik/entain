package webservice

type Config struct {
	AppPort string `envconfig:"APP_PORT" default:"8080"`
	AppHost string `envconfig:"APP_HOST" default:"localhost"`
}
