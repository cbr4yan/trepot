package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Environment string
	Database    Database
	HTTPServer  HTTPServer
}

type Database struct {
	Dsn             string
	MaxPoolSize     int
	MaxIdlePoolSize int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

type HTTPServer struct {
	Addr string
}

func Load(prefix string) (*Config, error) {
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	cfg := &Config{}
	if err := envconfig.Process(prefix, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c Config) IsDevelopment() bool {
	return c.Environment == "development"
}

func (c Config) IsProduction() bool {
	return c.Environment == "production"
}
