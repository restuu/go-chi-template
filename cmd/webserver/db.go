package main

import (
	"os"

	"go-chi-template/internal/pkg/db"
	"go-chi-template/internal/pkg/db/sql"
)

func connectDB() (db.SQL, error) {
	return sql.Connect(os.Getenv("DATABASE_DSN"))
}
