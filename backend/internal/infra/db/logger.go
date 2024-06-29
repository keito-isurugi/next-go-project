package db

import (
	"time"
	"runtime"
	"strings"
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

const (
	SlowThresholdDuration = 100 * time.Millisecond
)

type GormLogger struct {
	*zap.Logger
	LogLevel logger.LogLevel
	SlowThreshold time.Duration
	SkipCallerLookup bool
}

func initGormLogger(zapLogger *zap.Logger) *GormLogger {
	return &GormLogger{
		Logger:           zapLogger,
		LogLevel:         logger.Info,
		SlowThreshold:    SlowThresholdDuration,
		SkipCallerLookup: false,
	}
}


func (l *GormLogger) logger() *zap.Logger {
	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		default:
			return l.Logger.WithOptions(zap.AddCallerSkip(i))
		}
	}
	return l.Logger
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	traceID := getTraceID(ctx)
	amznTraceID := getAmznTraceID(ctx)
	l.Logger.Info(
		"sql",
		zap.String("msg", msg),
		zap.Any("data", data),
		zap.String("trace_id", traceID),
		zap.String("amzn_trace_id", amznTraceID),
	)
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	traceID := getTraceID(ctx)
	amznTraceID := getAmznTraceID(ctx)
	l.Logger.Warn(
		"sql",
		zap.String("msg", msg),
		zap.Any("data", data),
		zap.String("trace_id", traceID),
		zap.String("amzn_trace_id", amznTraceID),
	)
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	traceID := getTraceID(ctx)
	amznTraceID := getAmznTraceID(ctx)
	l.Logger.Error(
		"sql",
		zap.String("msg", msg),
		zap.Any("data", data),
		zap.String("trace_id", traceID),
		zap.String("amzn_trace_id", amznTraceID),
	)
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	traceID := getTraceID(ctx)
	amznTraceID := getAmznTraceID(ctx)
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		l.Logger.Error(
			"sql",
			zap.String("msg", "trace"),
			zap.Error(err),
			zap.Duration("elapsed", elapsed),
			zap.String("sql", sql),
			zap.Int64("rows", rows),
			zap.String("trace_id", traceID),
			zap.String("amzn_trace_id", amznTraceID),
		)
	} else {
		l.Logger.Debug(
			"sql",
			zap.String("msg", "trace"),
			zap.Duration("elapsed", elapsed),
			zap.String("sql", sql),
			zap.Int64("rows", rows),
			zap.String("trace_id", traceID),
			zap.String("amzn_trace_id", amznTraceID),
		)
	}
}

func getTraceID(ctx context.Context) string {
	traceID, ok := ctx.Value("trace_id").(string)
	if !ok {
		return ""
	}
	return traceID
}

func getAmznTraceID(ctx context.Context) string {
	amznTraceID, ok := ctx.Value("amzn_trace_id").(string)
	if !ok {
		return ""
	}
	return amznTraceID
}

