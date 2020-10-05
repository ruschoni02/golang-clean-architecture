package config

type Config struct {
	AppName     string `envconfig:"APP_NAME" default:"golang_clean_architecture"`
	AppEnv      string `envconfig:"APP_ENV" default:"dev"`
	HttpAddress string `envconfig:"HTTP_ADDRESS" default:"0.0.0.0:8000"`
	Debug       bool   `envconfig:"DEBUG" default:"false"`
}
