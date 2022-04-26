package main

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rl404/fairy/cache"
	"github.com/rl404/fairy/log"
	"github.com/rl404/mal-cover/internal/utils"
)

type config struct {
	App   appConfig   `envconfig:"APP"`
	Cache cacheConfig `envconfig:"CACHE"`
	Log   logConfig   `envconfig:"LOG"`
}

type appConfig struct {
	Port            string        `envconfig:"PORT" default:"34001" validate:"required" mod:"no_space"`
	ReadTimeout     time.Duration `envconfig:"READ_TIMEOUT" default:"1m" validate:"required,gt=0"`
	WriteTimeout    time.Duration `envconfig:"WRITE_TIMEOUT" default:"1m" validate:"required,gt=0"`
	GracefulTimeout time.Duration `envconfig:"GRACEFUL_TIMEOUT" default:"10s" validate:"required,gt=0"`
}

type cacheConfig struct {
	Dialect  string        `envconfig:"DIALECT" default:"inmemory" validate:"required,oneof=nocache redis inmemory memcache" mod:"no_space,lcase"`
	Address  string        `envconfig:"ADDRESS"`
	Password string        `envconfig:"PASSWORD"`
	Time     time.Duration `envconfig:"TIME" default:"24h" validate:"required,gt=0"`
}

type logConfig struct {
	Type  log.LogType  `envconfig:"TYPE" default:"2"`
	Level log.LogLevel `envconfig:"LEVEL" default:"-1"`
	JSON  bool         `envconfig:"JSON" default:"false"`
	Color bool         `envconfig:"COLOR" default:"true"`
}

const envPath = "../../.env"
const envPrefix = "MC"

var cacheType = map[string]cache.CacheType{
	"nocache":  cache.NoCache,
	"redis":    cache.Redis,
	"inmemory": cache.InMemory,
	"memcache": cache.Memcache,
}

func getConfig() (*config, error) {
	var cfg config

	// Load .env file.
	_ = godotenv.Load(envPath)

	// Convert env to struct.
	if err := envconfig.Process(envPrefix, &cfg); err != nil {
		return nil, err
	}

	// Override PORT env.
	if port := os.Getenv("PORT"); port != "" {
		cfg.App.Port = port
	}

	// Validate.
	if err := utils.Validate(&cfg); err != nil {
		return nil, err
	}

	// Init global log.
	if err := utils.InitLog(cfg.Log.Type, cfg.Log.Level, cfg.Log.JSON, cfg.Log.Color); err != nil {
		return nil, err
	}

	return &cfg, nil
}
