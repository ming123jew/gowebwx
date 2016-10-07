package webwx

import (
	"fmt"
	"testing"
)

const ENABLED = false

func _doLog() {
	msg := "this is a test msg %s %v"
	s := "OK"
	i := 123.456
	// no format
	log.Trace(msg, s, i)
	log.Debug(msg, s, i)
	log.Info(msg, s, i)
	log.Warn(msg, s, i)
	log.Error(msg, s, i)
	log.Critical(msg, s, i)
	// with format
	log.Tracef(msg, s, i)
	log.Debugf(msg, s, i)
	log.Infof(msg, s, i)
	log.Warnf(msg, s, i)
	log.Errorf(msg, s, i)
	log.Criticalf(msg, s, i)
}

func TestLogger(t *testing.T) {
	if !ENABLED {
		t.Skip("not enabled!")
	}
	for _, level := range []string{"trace", "debug", "info", "warn", "error", "critical", "off"} {
		fmt.Println("---- " + level + " ----")
		SetLogLevel(level)
		_doLog()
		log.Flush()
	}
}
