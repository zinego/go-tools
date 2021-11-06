package log

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
