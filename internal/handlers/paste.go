package paste

import (
	"fmt"
	"log/slog"
	"net/http"

	"pasteBin-backend/internal/db/models"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type PasteRequest struct {
	CipherText string `json:"ciphertext" validate:"required"`
}

func ResponseWithError(c echo.Context, ErrorResponseCode int, message string, err error) error {
	slog.Error(message, "err", err)
	return c.JSON(ErrorResponseCode, map[string]string{
		"error": message,
	})
	var bun bun.DB
}
