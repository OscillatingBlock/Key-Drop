package models

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/uptrace/bun"
)

type User struct {
	ID        string `bun:"id,pk"`
	PublicKey string `bun:"public_key,notnull,unique"`
}

func CreateUserTables(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().Model((*User)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		slog.Error("Error while creating User table", "table", "user", "error", err)
		return err
	}
	return nil
}

func CreateUser(ctx context.Context, db *bun.DB, user *User) error {
	if user.ID == "" || user.PublicKey == "" {
		return fmt.Errorf("Missing required fields for user")
	}

	// TODO: add validation for public key using utils/crypto 's validatePublicKeyFunc
	_, err := db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		slog.Error("Error while creating user", "id", user.ID, "error", err)
		return err
	}
	return nil
}

func GetUserByID(ctx context.Context, db *bun.DB, id string) (*User, error) {
	if id == "" {
		slog.Error("Cannot get user by ID, invalid user ID", "id", id)
		return nil, fmt.Errorf("invalid user ID")
	}
	var user User
	err := db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		slog.Error("Error while getting user", "id", id, "error", err)
		return nil, err
	}
	return &user, nil
}

func GetUserByPublicKey(ctx context.Context, db *bun.DB, publicKey string) (*User, error) {
	if publicKey == "" {
		slog.Error("Cannot get user by Public Key, Invalid public key", "public key", publicKey)
		return nil, fmt.Errorf("invalid public key")
	}
	// TODO: add validation for public key using utils/crypto 's validatePublicKeyFunc
	var user User
	err := db.NewSelect().Model(&user).Where("public_key = ?", publicKey).Scan(ctx)
	if err != nil {
		slog.Error("Error while getting user by public key", "public_key", publicKey[:8], "error", err)
		return nil, err
	}
	return &user, nil
}

func UserExists(ctx context.Context, db *bun.DB, publicKey string) (bool, error) {
	exists, err := db.NewSelect().Model((*User)(nil)).Where("public_key = ?", publicKey).Exists(ctx)
	if err != nil {
		slog.Error("Error checking user existence", "public_key", publicKey[:8], "error", err)
		return false, err
	}
	return exists, nil
}
