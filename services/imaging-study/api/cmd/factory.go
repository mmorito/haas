package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/brpaz/echozap"
	openapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/mnes/haas/services/imaging-study/openapi"
	"github.com/mnes/haas/services/imaging-study/src/config"
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

	swagger, err := openapi.GetSwagger()
	if err != nil {
		log.Fatalf("GetSwagger failed: %v", err)
	}
	swagger.Servers = nil

	s.echo.Use(echozap.ZapLogger(zapLogger()))
	s.echo.Use(middleware.Recover())
	s.echo.Use(openapimiddleware.OapiRequestValidator(swagger))
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

func zapLogger() *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "time"
	loggerConfig.EncoderConfig.EncodeTime = jstTimeEncoder
	loggerConfig.EncoderConfig.LevelKey = "severity"
	loggerConfig.EncoderConfig.StacktraceKey = "stack_trace"
	loggerConfig.EncoderConfig.NameKey = "logger"

	zapLogger, err := loggerConfig.Build()
	if err != nil {
		log.Fatalf("logger failed: %v", err)
	}

	zap.ReplaceGlobals(zapLogger)

	return zapLogger
}

func jstTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	const layout = "2006-01-02T15:04:05+09:00"
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	enc.AppendString(t.In(jst).Format(layout))
}
