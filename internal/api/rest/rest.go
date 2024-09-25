package rest

import (
	"net/http"
)

type HttpServer interface {
	Start(int) error

	// handlers
	Ping(http.ResponseWriter, *http.Request)
	Nodes(http.ResponseWriter, *http.Request)
}
