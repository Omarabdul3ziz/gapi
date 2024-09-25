package standardhttp

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/tfgrid-sdk-go/gapi/internal/api/rest"
)

func (c *Server) Ping(w http.ResponseWriter, r *http.Request) {
	err := c.dbClient.Ping(r.Context())
	if err != nil {
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Error().Err(err).Msg("failed to encode return object")
		}
		return
	}

	if err := json.NewEncoder(w).Encode("PONG"); err != nil {
		log.Error().Err(err).Msg("failed to encode return object")
	}
}

func (c *Server) Nodes(w http.ResponseWriter, r *http.Request) {
	nodes, err := c.dbClient.GetNodes(r.Context())
	if err != nil {
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Error().Err(err).Msg("failed to encode return object")
		}
		return
	}

	res := rest.ConvertNodes(nodes)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Error().Err(err).Msg("failed to encode return object")
	}
}
