package logger

import (
	"errors"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(debug bool) (*zap.Logger, error) {
	logLevel := zap.InfoLevel
	if debug {
		logLevel = zap.DebugLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		logLevel,
	)

	options := []zap.Option{
		zap.AddCaller(),                       // ここで呼び出し元の情報を追加
		zap.AddStacktrace(zapcore.ErrorLevel), // ここでエラーレベル以上のスタックトレースを追加
	}

	logger := zap.New(core, options...)
	if logger == nil {
		return nil, errors.New("failed to initialize logger")
	}

	return logger, nil
}
