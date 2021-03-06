package xlog

import (
	"io"

	"github.com/AleksandrGurkin/xerror"
)

var ErrorInitLogger = xerror.CommonError{
	Code:    10101,
	Message: "can't initialize logger",
	Err:     nil,
}

// Logger interface for integration_common
//
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Tracef(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Trace(msg string)
	Error(msg string)
	Fatal(msg string)

	WithXFields(fields Fields) Logger
	WithXField(key string, value interface{}) Logger
}

type Fields map[string]interface{}

type LoggerCfg struct {
	Level      string
	TimeFormat string
	Out        io.Writer
}
