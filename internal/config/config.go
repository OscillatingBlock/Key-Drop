package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
	DNS  string
}

func GetConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Warning: .env file not found or could not be loaded.", "error", err)
		return nil, err
	} else {
		slog.Info(".env file loaded successfully.")
	}
	port_str := os.Getenv("PORT")
	if port_str == "" {
		slog.Warn("PORT not found in .env, defaulting to 8080")
		port_str = "8080"
	}

	dns_str := os.Getenv("DNS")
	if dns_str == "" {
		slog.Error("DNS not found in .env")
		return nil, fmt.Errorf("Error while getting DNS from .env")
	}

	var conf Config
	conf = Config{PORT: port_str, DNS: dns_str}
	return &conf, nil
}
