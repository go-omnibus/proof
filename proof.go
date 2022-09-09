package proof

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *Proof

	encoderNameToConstructor = map[string]func(zapcore.EncoderConfig) zapcore.Encoder{
		ConsoleEncoder: func(encoderConfig zapcore.EncoderConfig) zapcore.Encoder {
			return zapcore.NewConsoleEncoder(encoderConfig)
		},
		JSONEncoder: func(encoderConfig zapcore.EncoderConfig) zapcore.Encoder {
			return zapcore.NewJSONEncoder(encoderConfig)
		},
	}
)

type Proof struct {
	Z *zap.Logger
}

func infoLevel() zap.LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		return level < zapcore.WarnLevel
	}
}

func warnLevel() zap.LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		return level >= zapcore.WarnLevel
	}
}

func Sync() error {
	return Logger.Z.Sync()
}

func Info(msg string, args ...zap.Field) {
	Logger.Z.Info(msg, args...)
}

func Error(msg string, args ...zap.Field) {
	Logger.Z.Error(msg, args...)
}

func Warn(msg string, args ...zap.Field) {
	Logger.Z.Warn(msg, args...)
}

func Debug(msg string, args ...zap.Field) {
	Logger.Z.Debug(msg, args...)
}

func Fatal(msg string, args ...zap.Field) {
	Logger.Z.Fatal(msg, args...)
}

func Infof(format string, args ...interface{}) {
	logMsg := fmt.Sprintf(format, args...)
	Logger.Z.Info(logMsg)
}

func Errorf(format string, args ...interface{}) {
	logMsg := fmt.Sprintf(format, args...)
	Logger.Z.Error(logMsg)
}

func Warnf(format string, args ...interface{}) {
	logMsg := fmt.Sprintf(format, args...)
	Logger.Z.Warn(logMsg)
}

func Debugf(format string, args ...interface{}) {
	logMsg := fmt.Sprintf(format, args...)
	Logger.Z.Debug(logMsg)
}

func Fatalf(format string, args ...interface{}) {
	logMsg := fmt.Sprintf(format, args...)
	Logger.Z.Fatal(logMsg)
}
