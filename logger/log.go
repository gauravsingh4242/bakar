package logger

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func InitLogger(logLevel string) {
	if len(logLevel) == 0 {
		logLevel = "info"
	}
	//logger, _ := zap.Config{
	//	Encoding: "json",
	//	// Level:            zap.NewAtomicLevelAt(getLogLevel(logLevel)),
	//	OutputPaths:      []string{"stdout"},
	//	ErrorOutputPaths: []string{"stdout"},
	//	EncoderConfig: zapcore.EncoderConfig{TimeKey: "timestamp", EncodeTime: zapcore.ISO8601TimeEncoder,
	//		MessageKey: "message", LevelKey: "level", EncodeLevel: zapcore.LowercaseLevelEncoder},
	//}.Build()

	logger, _ := zap.NewProduction()
	Log = logger.Sugar()
}

func Sync() {
	Log.Sync()
}
