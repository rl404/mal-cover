package main

import (
	"os"
	"time"

	"github.com/allegro/bigcache"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Cfg is config for the project.
var Cfg Config

// Cache is cache instance.
var Cache *bigcache.BigCache

// Config is configuration model from `.env`.
type Config struct {
	// HTTP port.
	Port string `envconfig:"PORT"`
	// Caching time (in seconds).
	Cache int `envconfig:"CACHE"`
}

const (
	// envPrefix is env prefix name for this project.
	envPrefix = "MC"
	// defaultPort is default HTTP port.
	defaultPort = "34001"
	// defaultCache is default caching time (1 day).
	defaultCache = 86400
)

// GetConfig to read and parse config from `.env`.
func GetConfig() error {
	// Get default config.
	Cfg = defaultConfig()

	// Load .env.
	godotenv.Load()

	// Parse .env.
	err := envconfig.Process(envPrefix, &Cfg)
	if err != nil {
		return err
	}

	// Override PORT env.
	port := os.Getenv("PORT")
	if port != "" {
		Cfg.Port = port
	}

	// Prepare the port.
	Cfg.Port = ":" + Cfg.Port

	// Set cache.
	Cache, err = bigcache.NewBigCache(bigcache.DefaultConfig(time.Duration(Cfg.Cache) * time.Second))
	if err != nil {
		return err
	}

	return nil
}

// defaultConfig to  get default config.
func defaultConfig() (cfg Config) {
	return Config{
		Port:  defaultPort,
		Cache: defaultCache,
	}
}
