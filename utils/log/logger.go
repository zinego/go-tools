package log

import (
	"fmt"
	"os"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.SugaredLogger

func getProjectRoot() string {
	pwd, _ := os.Getwd()
	for {
		if pwd == "/" {
			return ""
		}
		entryList, err := os.ReadDir(pwd)
		if err != nil {
			return ""
		}
		for _, v := range entryList {
			if v.Name() == "go.mod" {
				return pwd
			}
		}
		pwd = path.Dir(pwd)
	}
}

func Init() {
	projectRoot := getProjectRoot()
	hook := lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/output/logs/%s.log", projectRoot, path.Base(projectRoot)),
		MaxSize:    1024,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   true,
		LocalTime:  true,
	}
	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), zapcore.AddSync(os.Stdout), zap.DebugLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(&hook), zap.DebugLevel),
	)
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
	defer logger.Sync()
}

func Debug(args ...interface{})                        { logger.Debug(args...) }
func Info(args ...interface{})                         { logger.Info(args...) }
func Warn(args ...interface{})                         { logger.Warn(args...) }
func Error(args ...interface{})                        { logger.Error(args...) }
func DPanic(args ...interface{})                       { logger.DPanic(args...) }
func Panic(args ...interface{})                        { logger.Panic(args...) }
func Fatal(args ...interface{})                        { logger.Fatal(args...) }
func Debugf(template string, args ...interface{})      { logger.Debugf(template, args...) }
func Infof(template string, args ...interface{})       { logger.Infof(template, args...) }
func Warnf(template string, args ...interface{})       { logger.Warnf(template, args...) }
func Errorf(template string, args ...interface{})      { logger.Errorf(template, args...) }
func DPanicf(template string, args ...interface{})     { logger.DPanicf(template, args...) }
func Panicf(template string, args ...interface{})      { logger.Panicf(template, args...) }
func Fatalf(template string, args ...interface{})      { logger.Fatalf(template, args...) }
func Debugw(msg string, keysAndValues ...interface{})  { logger.Debugw(msg, keysAndValues...) }
func Infow(msg string, keysAndValues ...interface{})   { logger.Infow(msg, keysAndValues...) }
func Warnw(msg string, keysAndValues ...interface{})   { logger.Warnw(msg, keysAndValues...) }
func Errorw(msg string, keysAndValues ...interface{})  { logger.Errorw(msg, keysAndValues...) }
func DPanicw(msg string, keysAndValues ...interface{}) { logger.DPanicw(msg, keysAndValues...) }
func Panicw(msg string, keysAndValues ...interface{})  { logger.Panicw(msg, keysAndValues...) }
func Fatalw(msg string, keysAndValues ...interface{})  { logger.Fatalw(msg, keysAndValues...) }
