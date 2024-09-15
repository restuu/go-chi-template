package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type app struct {
	srv *http.Server
}

func (a *app) start() {
	go func() {
		slog.Info("starting server", slog.String("addr", a.srv.Addr))
		if err := a.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-sigint
	a.shutdown()
}

func (a *app) shutdown() {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dead := make(chan struct{})

	go func() {
		innerCtx := context.Background()
		if err := a.srv.Shutdown(innerCtx); err != nil {
			log.Fatal("failed to shutdown", err)
		}

		close(dead)
	}()

	select {
	case <-timeoutCtx.Done():
		log.Print("timed out shutting down")
	case <-dead:
		log.Print("server shutdown")
	}
}
