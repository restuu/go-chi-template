//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

func initApp() (*app, error) {
	wire.Build(
		initRouter,
		wire.Bind(new(http.Handler), new(chi.Router)),
		initServer,
		wire.Struct(new(app), "*"),
	)

	return &app{}, nil
}
