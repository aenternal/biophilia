package main

import (
	"biophilia/internal/config"
	"biophilia/internal/domain/services"
	"biophilia/internal/repositories/storage"
	"biophilia/internal/transport/http/rest"
	"fmt"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Получение конфига
	cfg := config.MustLoad()

	// Инициализация slog
	logger := setupLogger(cfg.Env)

	// Инициализация MinIO
	minioRepo, err := storage.NewMinioRepository(
		fmt.Sprintf("%s:%s", cfg.StorageHost, cfg.StoragePort),
		cfg.StorageSecretKey,
		cfg.StorageAccessKey,
		cfg.StorageBucket,
		false,
	)
	if err != nil {
		logger.Error("failed to connect to MinIO", "error", err)
		return
	}

	logger.Info("Successfully connected to MinIO")

	// Инициализация сервиса
	biomoleculeService := services.NewBiomoleculeService(minioRepo)

	// Инициализация Echo
	e := echo.New()

	// Заменим Echo логгер на slog
	e.Logger.SetOutput(logger.Writer()) // Используем writer slog для вывода логов Echo

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logger.Writer(), // Логи запросов через slog
	}))
	e.Use(middleware.Recover())

	// Инициализация маршрутов
	rest.InitRoutes(e, biomoleculeService)

	// Запуск сервера
	logger.InfoContext(ctx, "Starting server on port :8080")
	if err := e.Start(":8080"); err != nil {
		logger.ErrorContext(ctx, "Failed to start server", "error", err)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
