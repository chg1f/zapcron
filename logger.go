package zapcron

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type CronLogger struct {
	L *zap.Logger
}

func NewLogger(l *zap.Logger) cron.Logger {
	return &CronLogger{L: l.WithOptions(zap.AddCallerSkip(2))} // FIXME: caller skip depth
}

// Info logs routine messages about cron's operation.
func (cl *CronLogger) Info(msg string, keysAndValues ...interface{}) {
	if keysAndValues == nil || len(keysAndValues) == 0 {
		cl.L.Sugar().Info(msg)
		return
	} else if len(keysAndValues)%2 != 0 {
		cl.L.Sugar().Info(msg, keysAndValues)
		return
	}
	fields := make([]zap.Field, len(keysAndValues))
	for ix := 0; ix < len(keysAndValues); ix += 2 {
		key, ok := keysAndValues[ix].(string)
		if !ok {
			key = fmt.Sprintf("%v", keysAndValues[ix])
		}
		fields = append(fields, zap.Any(key, keysAndValues[ix+1]))
	}
	cl.L.Info(msg, fields...)
}

// Error logs an error condition.
func (cl *CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	if keysAndValues == nil || len(keysAndValues) == 0 {
		cl.L.Sugar().Error(msg)
		return
	} else if len(keysAndValues)%2 != 0 {
		cl.L.Sugar().Error(msg, keysAndValues)
		return
	}
	fields := make([]zap.Field, len(keysAndValues))
	for ix := 0; ix < len(keysAndValues); ix += 2 {
		key, ok := keysAndValues[ix].(string)
		if !ok {
			key = fmt.Sprintf("%v", keysAndValues[ix])
		}
		fields = append(fields, zap.Any(key, keysAndValues[ix+1]))
	}
	cl.L.Error(msg, fields...)
}
