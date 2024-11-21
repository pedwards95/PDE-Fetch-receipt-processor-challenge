package logger

import (
	"context"
	"fmt"
)

// simple logger implementation

// Logger ...
type Logger struct{}

// Log ...
type Log struct {
	level     string
	message   string
	requestID string
	tags      map[string]string
}

// New constructor for logger
func New() *Logger {
	return &Logger{}
}

// printLog very simple fill-in for a logging structure
func (lg *Logger) printLog(log *Log) {
	msg := fmt.Sprintf("%s: %s; REQUEST_ID:%s; TAGS: %+v", log.level, log.message, log.requestID, log.tags)
	fmt.Println(msg)
}

// Infof info level log, format string
func (lg *Logger) Infof(ctx context.Context, fStr string, params ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			lg.Errorf(ctx, "Panic in Infof log: %+v", err)
		}
	}()
	msg := fmt.Sprintf(fStr, params...)
	requestID := "-1"
	if rid := ctx.Value("request_id"); rid != nil {
		requestID = fmt.Sprintf("%+v", rid)
	}
	log := &Log{
		level:     "info",
		message:   msg,
		requestID: requestID,
		tags:      make(map[string]string),
	}
	lg.printLog(log)
}

// Errorf error level log, format string
func (lg *Logger) Errorf(ctx context.Context, fStr string, params ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			lg.Errorf(ctx, "Panic in Errorf log: %+v", err)
		}
	}()
	msg := fmt.Sprintf(fStr, params...)
	requestID := "-1"
	if rid := ctx.Value("request_id"); rid != nil {
		requestID = fmt.Sprintf("%+v", rid)
	}
	log := &Log{
		level:     "error",
		message:   msg,
		requestID: requestID,
		tags:      make(map[string]string),
	}
	lg.printLog(log)
}
