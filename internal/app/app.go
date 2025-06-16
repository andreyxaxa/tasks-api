package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/andreyxaxa/tasks-api/config"
	"github.com/andreyxaxa/tasks-api/internal/controller/http"
	"github.com/andreyxaxa/tasks-api/internal/repo/inmemory"
	"github.com/andreyxaxa/tasks-api/internal/usecase/tasks"
	"github.com/andreyxaxa/tasks-api/pkg/httpserver"
)

func Run(cfg *config.Config) {
	// Repository
	repo := inmemory.New()

	// Use-Case
	tasksUseCase := tasks.New(repo)

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port))
	http.NewRouter(httpServer.Router, tasksUseCase)

	// Start server
	httpServer.Start()
	log.Printf("http server started at :%s", cfg.HTTP.Port)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("app - Run - signal: %s", s.String())
	case err := <-httpServer.Notify():
		log.Fatal(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := httpServer.Shutdown(ctx)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	log.Println("graceful shutdown complete")
}
