package main

import (
	_http "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/rl404/fairy/cache"
	"github.com/rl404/mal-cover/internal/delivery/rest/api"
	"github.com/rl404/mal-cover/internal/delivery/rest/ping"
	malRepository "github.com/rl404/mal-cover/internal/domain/mal/repository"
	malCache "github.com/rl404/mal-cover/internal/domain/mal/repository/cache"
	malHTTP "github.com/rl404/mal-cover/internal/domain/mal/repository/http"
	"github.com/rl404/mal-cover/internal/pkg/http"
	"github.com/rl404/mal-cover/internal/service"
	"github.com/rl404/mal-cover/internal/utils"
)

func server() error {
	// Get config.
	cfg, err := getConfig()
	if err != nil {
		return err
	}
	utils.Info("config initialized")

	// Init cache.
	c, err := cache.New(cacheType[cfg.Cache.Dialect], cfg.Cache.Address, cfg.Cache.Password, cfg.Cache.Time)
	if err != nil {
		return err
	}
	utils.Info("cache initialized")
	defer c.Close()

	// Init in-memory.
	im, err := cache.New(cache.InMemory, "", "", time.Minute)
	if err != nil {
		return err
	}
	utils.Info("in-memory initialized")
	defer im.Close()

	// Init mal.
	var mal malRepository.Repository
	mal = malHTTP.New(_http.Client{Timeout: 10 * time.Second})
	mal = malCache.New(c, mal)

	// Init service.
	service := service.New(mal)
	utils.Info("service initialized")

	server := http.New(http.Config{
		Port:            cfg.App.Port,
		ReadTimeout:     cfg.App.ReadTimeout,
		WriteTimeout:    cfg.App.WriteTimeout,
		GracefulTimeout: cfg.App.GracefulTimeout,
	})
	utils.Info("server initialized")

	r := server.Router()
	r.Use(middleware.RealIP)
	r.Use(utils.Recoverer)
	utils.Info("server middleware initialized")

	// Register ping route.
	ping.New().Register(r)
	utils.Info("route ping initialized")

	// Register api route.
	api.New(service).Register(r)
	utils.Info("route api initialized")

	// Run web server.
	serverChan := server.Run()
	utils.Info("server listening at :%s", cfg.App.Port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case err := <-serverChan:
		if err != nil {
			return err
		}
	case <-sigChan:
	}

	return nil
}
