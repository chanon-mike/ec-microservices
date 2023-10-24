package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chanon-mike/ec-microservices/config"
	"github.com/chanon-mike/ec-microservices/modules/middleware/middlewareHandler"
	"github.com/chanon-mike/ec-microservices/modules/middleware/middlewareRepository"
	"github.com/chanon-mike/ec-microservices/modules/middleware/middlewareUsecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	server struct {
		app        *echo.Echo
		db         *mongo.Client
		cfg        *config.Config
		middleware middlewareHandler.MiddlewareHandlerService
	}
)

func newMiddleware(cfg *config.Config) middlewareHandler.MiddlewareHandlerService {
	repo := middlewareRepository.NewMiddlewareRepository()
	usecase := middlewareUsecase.NewMiddlewareUsecase(repo)
	return middlewareHandler.NewMiddlewareHandler(cfg, usecase)
}

func (s *server) gracefulShutdown(pctx context.Context, quit <-chan os.Signal) {
	log.Printf("Start service: %s\n", s.cfg.App.Name)
	<-quit // Sendout a channel signal
	log.Printf("Shutting down service: %s\n", s.cfg.App.Name)

	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	if err := s.app.Shutdown(ctx); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func (s *server) httpListening() {
	if err := s.app.Start(s.cfg.App.Url); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error: %v", err)
	}
}

func Start(pctx context.Context, cfg *config.Config, db *mongo.Client) {
	s := &server{
		app:        echo.New(),
		db:         db,
		cfg:        cfg,
		middleware: newMiddleware(cfg),
	}

	// Basic middleware
	// Request Timeout
	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		Timeout:      30 * time.Second,
		ErrorMessage: "Error: Request Timeout",
	}))

	// CORS
	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	// Body Limit: maximum allowed size for request body
	s.app.Use(middleware.BodyLimit("10M"))

	switch s.cfg.App.Name {
	case "auth":
	case "player":
	case "item":
	case "inventory":
	case "payment":
	}

	// Graceful Shutdown
	// Wait for interrupt signal and use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1) // Open 1 channel for retrive 1 interrupt signal
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Logging
	s.app.Use(middleware.Logger())

	go s.gracefulShutdown(pctx, quit)

	// Listening
	s.httpListening()
}
