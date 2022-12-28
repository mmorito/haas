package logger

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/blendle/zapdriver"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/mnes/haas/src/services/imaging-study/config"
)

var TraceCtxKey = &contextKey{"trace"}

type contextKey struct {
	name string
}

type Trace struct {
	TraceID string
	SpanID  string
	Sampled bool
}

func ZapLogger() *zap.Logger {
	// create our uber zap configuration
	config := zapdriver.NewProductionConfig()
	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.EncodeTime = jstTimeEncoder
	config.EncoderConfig.LevelKey = "severity"
	config.EncoderConfig.StacktraceKey = "stack_trace"
	config.EncoderConfig.NameKey = "logger"

	// creates our logger instance
	zapLogger, err := config.Build()
	if err != nil {
		log.Fatalf("zap.config.Build(): %v", err)
	}

	zap.ReplaceGlobals(zapLogger)

	return zapLogger
}

func ExtractRequestField(zapLogger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			fields := []zapcore.Field{
				zap.String("remote_ip", c.RealIP()),
				zap.String("latency", time.Since(start).String()),
				zap.String("host", req.Host),
				zap.String("request", fmt.Sprintf("%s %s", req.Method, req.RequestURI)),
				zap.Int("status", res.Status),
				zap.Int64("size", res.Size),
				zap.String("user_agent", req.UserAgent()),
			}

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}
			fields = append(fields, zap.String("request_id", id))

			traceID := req.Header.Get("X-Cloud-Trace-Context")
			if traceID != "" {
				env := config.Env()
				traceID, spanID, sampled := deconstructXCloudTraceContext(traceID)
				fields = append(fields, zapdriver.TraceContext(traceID, spanID, sampled, env.ProjectID)...)
				fields = append(fields, zap.String("logName", fmt.Sprintf("projects/%s/logs/haas", env.ProjectID)))
			}

			log := zapLogger.With(fields...)

			n := res.Status
			switch {
			case n >= 500:
				log.With(zap.Error(err)).Error("Server error", fields...)
			case n >= 400:
				log.With(zap.Error(err)).Warn("Client error", fields...)
			case n >= 300:
				log.Info("Redirection", fields...)
			default:
				log.Info("Success", fields...)
			}

			return nil
		}
	}
}

var reCloudTraceContext = regexp.MustCompile(
	// Matches on "TRACE_ID"
	`([a-f\d]+)?` +
		// Matches on "/SPAN_ID"
		`(?:/([a-f\d]+))?` +
		// Matches on ";0=TRACE_TRUE"
		`(?:;o=(\d))?`)

func deconstructXCloudTraceContext(s string) (traceID, spanID string, traceSampled bool) {
	matches := reCloudTraceContext.FindStringSubmatch(s)

	traceID, spanID, traceSampled = matches[1], matches[2], matches[3] == "1"

	if spanID == "0" {
		spanID = ""
	}

	return
}

func jstTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	const layout = "2006-01-02T15:04:05+09:00"
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	enc.AppendString(t.In(jst).Format(layout))
}
