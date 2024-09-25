package standardhttp

import (
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/tfgrid-sdk-go/gapi/internal/api/rest"
	"github.com/threefoldtech/tfgrid-sdk-go/gapi/internal/db"
	gormpgquery "github.com/threefoldtech/tfgrid-sdk-go/gapi/internal/db/gorm-pg-query"
)

var _ rest.HttpServer = (*Server)(nil)

type Server struct {
	dbClient db.QueryClient
}

func NewServer(dns string) (*Server, error) {
	dbClient, err := gormpgquery.NewDBQueryClient(dns)
	if err != nil {
		return nil, err
	}

	return &Server{
		dbClient: dbClient,
	}, nil
}

func (s *Server) Start(port int) error {
	http.HandleFunc("/ping", s.Ping)
	http.HandleFunc("/nodes", s.Nodes)

	address := ":" + strconv.Itoa(port)
	log.Info().Str("address", address).Msg("standard http server started")
	return http.ListenAndServe(address, nil)
}
