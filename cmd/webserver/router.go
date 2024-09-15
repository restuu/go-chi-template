package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/hello", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("world"))
	})

	return r
}
