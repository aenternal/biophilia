package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Env                string `env:"ENV" envDefault:"local"`
	PostgresHost       string `env:"POSTGRES_HOST" envDefault:"localhost"`
	PostgresPort       string `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresUser       string `env:"POSTGRES_USER" envDefault:"postgres"`
	PostgresPassword   string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
	PostgresDBName     string `env:"POSTGRES_DB" envDefault:"biophilia"`
	PostgresTestDBName string `env:"POSTGRES_TEST_DB" envDefault:"biophilia-tests"`
	StorageHost        string `env:"STORAGE_HOST" envDefault:"localhost"`
	StoragePort        string `env:"STORAGE_PORT" envDefault:"9000"`
	StorageUser        string `env:"STORAGE_USER" envDefault:"minioadmin"`
	StorageAccessKey   string `env:"STORAGE_ACCESS_KEY"`
	StorageSecretKey   string `env:"STORAGE_SECRET_KEY"`
	StoragePassword    string `env:"STORAGE_PASSWORD" envDefault:"miniopassword"`
	StorageBucket      string `env:"STORAGE_BUCKET" envDefault:"biophilia"`
	StorageUseSSL      bool   `env:"STORAGE_USE_SSL" envDefault:"false"`
	StorageUrl         string `env:"STORAGE_URL" envDefault:"http://localhost:9000"`
	MigrationsPath     string `env:"MIGRATIONS_PATH"`
	RedisHost          string `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort          string `env:"REDIS_PORT" envDefault:"6379"`
	RedisPassword      string `env:"REDIS_PASSWORD"`
	RedisDB            int    `env:"REDIS_DB" envDefault:"0"`
}

func (cfg *Config) DataBaseDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDBName,
	)
}

// LoadConfig - загрузка конфигурации с поддержкой .env
func LoadConfig() (*Config, error) {
	// Загружаем .env файл по умолчанию, если он существует
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found, using environment variables directly")
	}

	configPath := fetchConfigPath()
	var cfg Config

	// Если путь не указан, читаем только env
	if configPath == "" {
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			return nil, fmt.Errorf("cannot read environment variables: %w", err)
		}
	} else {
		// Проверяем файл конфигурации
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			return nil, fmt.Errorf("config file does not exist: %s", configPath)
		}

		if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
			return nil, fmt.Errorf("cannot read config file: %w", err)
		}
	}

	return &cfg, nil
}

// MustLoad - загрузка конфигурации с паникой при ошибке
func MustLoad() *Config {
	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}

// fetchConfigPath - получение пути до конфигурационного файла
// Приоритет: флаг > переменная окружения CONFIG_PATH > пустое значение.
func fetchConfigPath() string {
	var path string
	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
