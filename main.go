package main

import (
	"github.com/golang-clean-architecture/app/adapters"
	"github.com/golang-clean-architecture/app/routes/transaction"
	"github.com/golang-clean-architecture/core/depedencies"
	"github.com/golang-clean-architecture/pkg/config"
	"github.com/golang-clean-architecture/pkg/health"
	"github.com/golang-clean-architecture/pkg/http"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
)

func checkFatal(logger adapters.LoggerAdapter, err error) {
	if err != nil {
		logger.Fatal("Fatal error", err)
	}
}

func main() {
	envConfig := &config.Config{}
	logger := adapters.NewLoggerAdapter()

	godotenv.Load()
	err := envconfig.Process("", envConfig)
	checkFatal(logger, err)

	server := http.NewServer(logger, envConfig.HttpAddress)
	server.Use(health.GinHandler("/health"))

	logger.Info("Loading modules")
	_, err = server.Load(
		"/v1",
		envConfig,
		transaction_route.New(),
	)
	checkFatal(logger, err)

	rootCmd := &cobra.Command{
		Use:                   "golang-clean-architecture [-h]",
		Short:                 "golang-clean-architecture",
		Version:               "0.0.1",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("Starting HTTP server", depedencies.Event{"address": envConfig.HttpAddress})
			err := server.Start()
			checkFatal(logger, err)
		},
	}

	checkFatal(logger, rootCmd.Execute())
}
