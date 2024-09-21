//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	userprovider "go-chi-template/internal/app/user/provider"

	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

func initApp() (*app, error) {
	wire.Build(
		connectDB,

		userprovider.UserUsecaseProvider,

		wire.Struct(new(usecases), "*"),

		initRouter,
		wire.Bind(new(http.Handler), new(chi.Router)),
		initServer,
		wire.Struct(new(app), "*"),
	)

	return &app{}, nil
}
