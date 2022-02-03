package empty

import (
	"github.com/AleksandrGurkin/xlog"
)

type Empty struct{}

func (Empty) Debugf(format string, args ...interface{}) {
}

func (Empty) Infof(format string, args ...interface{}) {
}

func (Empty) Warnf(format string, args ...interface{}) {
}

func (Empty) Tracef(format string, args ...interface{}) {
}

func (Empty) Errorf(format string, args ...interface{}) {
}

func (Empty) Fatalf(format string, args ...interface{}) {
}

func (Empty) Debug(msg string) {
}

func (Empty) Info(msg string) {
}

func (Empty) Warn(msg string) {
}

func (Empty) Trace(msg string) {
}

func (Empty) Error(msg string) {
}

func (Empty) Fatal(msg string) {
}

func (Empty) WithXFields(fields xlog.Fields) xlog.Logger {
	return Empty{}
}

func (Empty) WithXField(key string, value interface{}) xlog.Logger {
	return Empty{}
}
