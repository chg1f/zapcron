package zapcron

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type CronLogger struct {
	L *zap.Logger
}

func NewLogger(l *zap.Logger) cron.Logger {
	return &CronLogger{L: l.WithOptions(zap.AddCallerSkip(2))} // FIXME: caller skip depth
}

func (cl *CronLogger) log(lvl zapcore.Level, msg string, keysAndValues ...interface{}) {
	if keysAndValues == nil || len(keysAndValues) == 0 {
		cl.L.Log(lvl, msg)
	} else if len(keysAndValues)%2 != 0 {
		cl.L.Log(lvl, fmt.Sprintf("%s; %v", msg, keysAndValues))
	} else {
		fields := make([]zap.Field, 0, len(keysAndValues)/2)
		for ix := 0; ix < len(keysAndValues); ix += 2 {
			key, ok := keysAndValues[ix].(string)
			if !ok {
				key = fmt.Sprintf("%v", keysAndValues[ix])
			}
			fields = append(fields, zap.Any(key, keysAndValues[ix+1]))
		}
		cl.L.Log(lvl, msg, fields...)
	}
}

// Info logs routine messages about cron's operation.
func (cl *CronLogger) Info(msg string, keysAndValues ...interface{}) {
	cl.log(zapcore.InfoLevel, msg, keysAndValues...)
}

// Error logs an error condition.
func (cl *CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	cl.log(zapcore.ErrorLevel, msg, keysAndValues...)
}
