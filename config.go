package main

import (
	"os"
	"time"

	"github.com/allegro/bigcache"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var cfg config
var cache *bigcache.BigCache

type config struct {
	// HTTP port.
	Port string `envconfig:"PORT" default:"34001"`
	// Caching time (in seconds).
	Cache int `envconfig:"CACHE" default:"86400"`
}

const envPrefix = "MC"

func setConfig() (err error) {
	// Load .env.
	_ = godotenv.Load()

	// Parse .env.
	if err := envconfig.Process(envPrefix, &cfg); err != nil {
		return err
	}

	// Override PORT env.
	if port := os.Getenv("PORT"); port != "" {
		cfg.Port = port
	}

	// Prepare the port.
	cfg.Port = ":" + cfg.Port

	// Set cache.
	if cache, err = bigcache.NewBigCache(bigcache.DefaultConfig(time.Duration(cfg.Cache) * time.Second)); err != nil {
		return err
	}

	return nil
}
