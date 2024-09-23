package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
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
	StorageAccessKey   string `env:"STORAGE_ACCESS_KEY" envDefault:""`
	StorageSecretKey   string `env:"STORAGE_SECRET_KEY" envDefault:""`
	StoragePassword    string `env:"STORAGE_PASSWORD" envDefault:"miniopassword"`
	StorageBucket      string `env:"STORAGE_BUCKET" envDefault:"biophilia"`
	StorageUrl         string `env:"STORAGE_URL" envDefault:"http://localhost:9000"`
	MigrationsPath     string `env:"MIGRATIONS_PATH" envDefault:""`
}

func (cfg *Config) DbDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDBName,
	)
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
