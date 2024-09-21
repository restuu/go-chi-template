package main

import (
	"net/http"
	"time"

	userhttp "go-chi-template/internal/app/user/delivery/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog/v2"
)

func initRouter(uc usecases) chi.Router {
	r := chi.NewRouter()

	setupMiddlewares(r)

	r.Get("/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	})

	setupAppRoutes(r, uc)

	return r
}

func setupMiddlewares(r chi.Router) {
	logger := httplog.NewLogger("app", httplog.Options{
		JSON:           true,
		RequestHeaders: true,
		QuietDownRoutes: []string{
			"/ping",
			"/healthcheck",
		},
		QuietDownPeriod: 10 * time.Second,
	})

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(httplog.RequestLogger(logger))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	}))
}

func setupAppRoutes(r chi.Router, uc usecases) {
	userhttp.SetUserHTTPHandler(r, uc.UserUsecase)
}
