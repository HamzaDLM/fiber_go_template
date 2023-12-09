package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct{ *zap.Logger }

var logger *zap.Logger
var once sync.Once

// env string, appConfigFile embed.FS
func Get() *zap.Logger {
	once.Do(func() {
		stdout := zapcore.AddSync(os.Stdout)

		file := zapcore.AddSync(&lumberjack.Logger{
			Filename:   "app.log",
			MaxSize:    10, // megabytes
			MaxBackups: 3,
			MaxAge:     20, // days
		})

		// Here we define log level (debug - info - ....))
		level := zap.NewAtomicLevelAt(zap.DebugLevel)

		productionCfg := zap.NewProductionEncoderConfig()
		productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

		developmentCfg := zap.NewDevelopmentEncoderConfig()
		developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

		consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
		fileEncoder := zapcore.NewJSONEncoder(productionCfg)

		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, stdout, level),
			zapcore.NewCore(fileEncoder, file, level),
		)

		logger = zap.New(core)
		defer logger.Sync()
	})
	return logger
}
