package main

import "net/http"

func registerHandlers(
	mux *http.ServeMux,
	db Database,
) {
	mux.Handle("GET /api/v1/nodes", handleGetNodes(db))
}
