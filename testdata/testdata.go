package testdata

import (
	"log"
	"log/slog"

	"go.uber.org/zap"
)

func testLogs() {

	slog.Info("Starting server on port 8080") // want "lowercase"

	log.Info("запуск сервера") // want "English"

	slog.Info("server started!🚀") // want "invalid symbols"

	password := "secret123"
	log.Info("user password: " + password) // want "avoid sensitive"

	// correct examples
	slog.Info("starting server on port 8080")
	log.Info("server started")
	zap.L().Info("connection established")
}
