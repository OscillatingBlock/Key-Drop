package models

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/uptrace/bun"
)

type Paste struct {
	ID         string    `bun:"id,pk"`
	Ciphertext string    `bun:"ciphertext,notnull"`
	Signature  string    `bun:"signature,notnull"`
	PublicKey  string    `bun:"public_key,notnull"`
	ExpiresAt  time.Time `bun:"expires_at,notnull"`
}

func CreatePasteTables(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().Model((*Paste)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		slog.Error("Error while creating Paste table", "table", "paste", "error", err)
	}
	return nil
}

func CreatePaste(ctx context.Context, db *bun.DB, paste *Paste) error {
	if paste.ID == "" || paste.Ciphertext == "" || paste.Signature == "" || paste.PublicKey == "" {
		return fmt.Errorf("missing required fields")
	}
	if paste.ExpiresAt.IsZero() {
		paste.ExpiresAt = time.Now().Add(time.Hour)
	}
	_, err := db.NewInsert().Model(paste).Exec(ctx)
	if err != nil {
		slog.Error("Error while creating paste", "id", paste.ID, "error", err)
	}
	return nil
}

func GetPasteByID(ctx context.Context, db *bun.DB, id string) (*Paste, error) {
	var paste Paste
	err := db.NewSelect().Model(&paste).Where("id = ?", id).Scan(ctx)
	if err != nil {
		slog.Error("Error while getting paste", "pasteid", id, "error", err)
	}
	return &paste, nil
}

// TODO: func DeleteExpiredPastes
