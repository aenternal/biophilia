package main

import (
	_ "biophilia/docs"
	"biophilia/internal/data/clients"
	"biophilia/internal/data/repositories/database"
	"biophilia/internal/data/repositories/storage"
	"biophilia/internal/domain/services"
	"biophilia/internal/infrastructure/config"
	"biophilia/internal/infrastructure/logging"
	"biophilia/internal/presentation/http/rest"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           Biophilia API
// @version         0.0.1
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Получение конфига
	cfg := config.MustLoad()

	// Инициализация slog
	logger := logging.SetupLogger(cfg.Env)

	// Инициализация MinIO
	minioRepo, err := storage.NewMinioRepository(
		fmt.Sprintf("%s:%s", cfg.StorageHost, cfg.StoragePort),
		cfg.StorageUser,
		cfg.StoragePassword,
		cfg.StorageBucket,
		cfg.StorageUseSSL,
	)
	if err != nil {
		panic("Failed to connect to MinIO: " + err.Error())
	}

	logger.Info("Successfully connected to MinIO")

	// Инициализация БД
	db, err := sqlx.Open("postgres", cfg.DataBaseDSN())
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic("Failed to ping the database: " + err.Error())
	}

	// Инициализация БД-репозитория
	dbRepo := database.NewBiomoleculeRepository(db)

	// Инициализация инфраструктурного сервиса для работы с изображениями
	imageService := services.NewPlotImageService(logger)

	// Инициализация клиента Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	// Инициализация blast-репозитория
	blastRepository := clients.NewEBIBlastClient()

	// Инициализация сервиса с бизнес-логикой
	biomoleculeService := services.NewBiomoleculeService(logger, dbRepo, minioRepo, imageService, blastRepository, redisClient)

	// Инициализация Echo
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// Инициализация маршрутов
	rest.InitRoutes(e, biomoleculeService)

	// Запуск сервера
	logger.Info("Starting server on port :8080")
	if err := e.Start(":8080"); err != nil {
		logger.Error("Failed to start server", "error", err)
	}
}
