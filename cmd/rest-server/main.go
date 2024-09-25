package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	standardhttp "github.com/threefoldtech/tfgrid-sdk-go/gapi/internal/api/rest/standard-http"
	"github.com/threefoldtech/tfgrid-sdk-go/gapi/pkg/logging"
)

const (
	dns  = "postgresql://postgres:postgres@172.17.0.2:5432/tfgrid-graphql?sslmode=disable"
	port = 8080
)

func main() {
	logging.Init(zerolog.LevelInfoValue)

	s, err := standardhttp.NewServer(dns)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	if err := s.Start(port); err != nil {
		log.Fatal().Err(err).Send()
	}
}
