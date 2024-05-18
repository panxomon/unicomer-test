package main

import (
	router "unicomer-test/cmd/api"
	"unicomer-test/cmd/bootstrap"
	"unicomer-test/config"

	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.LoadEnvVars()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load env vars")
	}

	components := bootstrap.LoadComponents(cfg.UrlHolidays)

	r := router.SetupRouter(components)

	log.Info().Msg("Starting server on :8080")

	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatal().Err(err).Msg("could not start server")
	}
}
