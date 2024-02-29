package logger

import (
	"github.com/atpuxiner/grapi/app/initializer/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var Logger *zap.Logger

func InitLogger() {
	// Encoder
	encoder := getEncoder()
	// WriteSyncer
	infoWriteSyncer := getWriterSyncer(
		"log/info.log",
		conf.Conf.Log.MaxSize,
		conf.Conf.Log.MaxAge,
		conf.Conf.Log.MaxBackup,
		conf.Conf.Log.LocalTime,
		conf.Conf.Log.Compress,
	)
	errorWriteSyncer := getWriterSyncer(
		"log/error.log",
		conf.Conf.Log.MaxSize,
		conf.Conf.Log.MaxAge,
		conf.Conf.Log.MaxBackup,
		conf.Conf.Log.LocalTime,
		conf.Conf.Log.Compress,
	)
	// LevelEnabler
	lowLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { // info|debug
		lowLevel := zap.InfoLevel
		if conf.Conf.Log.Debug {
			lowLevel = zap.DebugLevel
		}
		return lvl < zap.ErrorLevel && lvl >= lowLevel
	})
	highLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { // error级别
		return lvl >= zap.ErrorLevel
	})
	// Logger
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoWriteSyncer, zapcore.AddSync(os.Stdout)), lowLevel)
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorWriteSyncer, zapcore.AddSync(os.Stdout)), highLevel)
	coreArr := []zapcore.Core{
		infoFileCore,
		errorFileCore,
	}
	Logger = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller())
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getWriterSyncer(filename string, maxSize, maxAge, maxBackup int, localTime, compress bool) zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackup,
		LocalTime:  localTime,
		Compress:   compress,
	}
	return zapcore.AddSync(lumberjackLogger)
}

// ==========format string, args ...interface{}==========

func Debugf(format string, args ...interface{}) {
	Logger.Sugar().Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	Logger.Sugar().Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	Logger.Sugar().Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	Logger.Sugar().Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	Logger.Sugar().Panicf(format, args...)
}

// ==========msg string, fields ...Field==========

func Debug(msg string, fields ...zapcore.Field) {
	Logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	Logger.Error(msg, fields...)
}

func Panic(msg string, fields ...zapcore.Field) {
	Logger.Panic(msg, fields...)
}
