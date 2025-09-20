package main

import (
	"encoding/json"
	"net/http"
)

func handleGetNodes(db Database) http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {

		nodes, err := db.GetNodes()
		if err != nil {
			http.Error(w, "failed to get nodes from db", http.StatusInternalServerError)
			return
		}

		resp, err := json.Marshal(nodes)
		if err != nil {
			http.Error(w, "failed to encode nodes response", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(200)
		w.Write(resp)
	}

	return http.HandlerFunc(h)
}
