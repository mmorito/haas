package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"

	"github.com/mnes/haas/src/services/imaging-study/config"
	"github.com/mnes/haas/src/services/imaging-study/logger"
)

type server struct {
	echo *echo.Echo
}

func newServer() *server {
	return &server{echo: echo.New()}
}

func (s *server) setConfig() {
	s.echo.HideBanner = true
	s.echo.HidePort = true

	// swagger, err := openapi.GetSwagger()
	// if err != nil {
	// 	log.Fatalf("GetSwagger failed: %v", err)
	// }
	// swagger.Servers = nil
	// s.echo.Use(openapimiddleware.OapiRequestValidator(swagger))

	s.echo.Use(logger.ExtractRequestField(logger.ZapLogger()))
	s.echo.Use(middleware.Recover())
}

func (s *server) start() {
	zap.L().Info("Start Server")
	if err := s.echo.Start(":" + config.Env().Port); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}

func (s *server) notifySignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	zap.L().Info("Shutdown Server")
}

func (s *server) gracefulStop() {
	if err := s.echo.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
	zap.L().Info("Server exiting")
}
