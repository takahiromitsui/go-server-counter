package main

import (
	"net/http"

	"github.com/takahiromitsui/go-server-counter/internal/handlers"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(("GET /counter"), handlers.Repo.Counter)
	return mux
}