package sql

import (
	"context"
	"time"

	"go-chi-template/internal/pkg/db"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

func Connect(dsn string) (db.SQL, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	// TODO: make configurable
	cfg.MaxConns = 20
	cfg.MinConns = 10
	cfg.MaxConnIdleTime = 5 * time.Minute
	cfg.MaxConnLifetime = 5 * time.Minute
	cfg.HealthCheckPeriod = time.Minute

	ctx := context.Background()
	conn, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(ctx); err != nil {
		return nil, err
	}

	return stdlib.OpenDBFromPool(conn), nil
}
