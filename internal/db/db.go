package db

import (
	"context"
	"database/sql"
	"log/slog"

	// ... other imports
	"pasteBin-backend/internal/db/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func initDB(DSN string) *bun.DB {
	var db *bun.DB
	if DSN == "" {
		slog.Error("Failed to get DSN from .env file.")
	}
	sqldb, err := sql.Open("mysql", DSN)
	if err != nil {
		slog.Error("Failed to open MySQL connection: ", "error", err)
	}
	db = bun.NewDB(sqldb, mysqldialect.New())
	return db
}

func createDB(db *bun.DB, ctx context.Context) error {
	// TODO: call createPasteTable and createUserTable from models/paste.go and models/users.go
	err := models.CreateUserTables(ctx, db)
	if err != nil {
		slog.Error("Error while making users table", "error", err)
	}
	err = models.CreatePasteTables(ctx, db)
	if err != nil {
		slog.Error("Error while creating paste tables", "error", err)
	}
	return nil
}
