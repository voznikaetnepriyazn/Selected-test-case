package tests

import (
	"log"
	"log/slog"

	"go.uber.org/zap"
)

func testLowercase() {
	slog.Info("Starting server")
	log.Print("Error occurred")
	zap.L().Warn("Connection lost")

	slog.Info("starting server")
	log.Print("error occurred")
	zap.L().Warn("connection lost")

	slog.Info("127.0.0.1 connected")
	log.Print("/api/v1/users called")

	slog.Info("  Capital after spaces")
}

func testEnglish() {
	slog.Info("запуск сервера")
	log.Print("ошибка подключения")

	slog.Info("started 🚀")

	slog.Info("server started")
	log.Print("connection error")
}

func testSymbols() {
	slog.Info("started!!!")
	slog.Info("waiting...")

	slog.Info("launched 🚀")

	slog.Info("started!")
	slog.Info("waiting..")
	slog.Info("version 1.2.3")
}

func testSensitive() {
	password := "secret123"
	apiKey := "key-abc"

	log.Info("password: " + password)
	slog.Debug("api_key=" + apiKey)

	slog.Info("user authenticated")
	log.Print("api request completed")
}
