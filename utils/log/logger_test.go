package log

import "testing"

func TestDebug(t *testing.T) {
	Init()
	Debug(1, 2, "3")
	Info(1, 2, "3")
	Warn(1, 2, "3")
	Error(1, 2, "3")
	// DPanic(1, 2, "3")
	// Panic(1, 2, "3")
	// Fatal(1, 2, "3")
	Debugf("template string %d-%d-%s", 1, 2, "3")
	Infof("template string %d-%d-%s", 1, 2, "3")
	Warnf("template string %d-%d-%s", 1, 2, "3")
	Errorf("template string %d-%d-%s", 1, 2, "3")
	// DPanicf("template string %d-%d-%s", 1, 2, "3")
	// Panicf("template string %d-%d-%s", 1, 2, "3")
	// Fatalf("template string %d-%d-%s", 1, 2, "3")
	Debugw("tag", "key", 1)
	Infow("tag", "key", 1)
	Warnw("tag", "key", 1)
	Errorw("tag", "key", 1)
	// DPanicw("tag", "key", 1)
	// Panicw("tag", "key", 1)
	// Fatalw("tag", "key", 1)

}
