package log

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	projectRoot := getProjectRoot()
	Init(
		WithMaxAge(7),
		WithMaxByteSize(1024*1024),
		WithMaxBackups(16),
		WithFileName(fmt.Sprintf("%s/output/logs/%s.log", projectRoot, filepath.Base(projectRoot))),
	)
}

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
		pwd = filepath.Dir(pwd)
	}
}

func Init(opts ...Option) {
	hook := lumberjack.Logger{
		Compress:  true,
		LocalTime: true,
	}
	for _, opt := range opts {
		opt.f(&hook)
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

type Option struct {
	f func(l *lumberjack.Logger)
}

func WithMaxByteSize(maxSize int) Option {
	return Option{func(l *lumberjack.Logger) {
		l.MaxSize = maxSize
	}}
}

func WithMaxBackups(maxBackups int) Option {
	return Option{func(l *lumberjack.Logger) {
		l.MaxBackups = maxBackups
	}}
}

func WithMaxAge(maxAge int) Option {
	return Option{func(l *lumberjack.Logger) {
		l.MaxAge = maxAge
	}}
}

func WithFileName(fileName string) Option {
	return Option{func(l *lumberjack.Logger) {
		l.Filename = fileName
	}}
}
