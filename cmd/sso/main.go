package main

import (
	"biophilia/internal/config"
	"biophilia/internal/domain/services"
	"biophilia/internal/repositories/database"
	"biophilia/internal/repositories/storage"
	"biophilia/internal/transport/http/rest"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
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

	// Инициализация БД
	db, err := sqlx.Connect("postgres", cfg.DbDSN())
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic("Failed to ping the database: " + err.Error())
	}

	// Инициализация БД-репозитория
	dbRepo := database.NewBiomoleculeRepository(db)

	// Инициализация сервиса
	biomoleculeService := services.NewBiomoleculeService(logger, dbRepo, minioRepo)

	// Инициализация Echo
	e := echo.New()

	// Инициализация маршрутов
	rest.InitRoutes(e, biomoleculeService)

	// Запуск сервера
	logger.Info("Starting server on port :8080")
	if err := e.Start(":8080"); err != nil {
		logger.Error("Failed to start server", "error", err)
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
