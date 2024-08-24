package logger

import (
	"context"
	"tt-go-sample-api/util"

	"github.com/sirupsen/logrus"
)

// APILogger is a wrapper for a log engine (e.g. logrus).
type APILogger struct {
	engine *logrus.Entry
}

// NewWithLogrus returns a pointer to APILogger,
// using logrus as its inner log engine.
//
// It takes an `appName` that will be ecoed
// throughout the application calls, and may be
// useful for Modern Application Performance
// Monitoring (APM).
func NewWithLogrus(appName string) *APILogger {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		DisableHTMLEscape: true,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
	})

	contextualizedLog := log.WithFields(logrus.Fields{
		"apiVersion": util.GetAPIVersion(),
		"env":        util.GetEnv(),
	})

	return &APILogger{
		engine: contextualizedLog,
	}
}

// Info logs an information at info level.
func (l *APILogger) Info(ctx context.Context, input LogInput) {
	l.engine.WithContext(ctx).WithFields(input.GetFields()).Infof(input.GetMessage())
}

// Warn logs an information at warn level.
func (l *APILogger) Warn(ctx context.Context, input LogInput) {
	l.engine.WithContext(ctx).WithFields(input.GetFields()).Warnf(input.GetMessage())
}

// Error logs an information at error level.
func (l *APILogger) Error(ctx context.Context, input LogInput) {
	l.engine.WithContext(ctx).WithFields(input.GetFields()).Errorf(input.GetMessage())
}

// Fatal logs an information at fatal level,
// thus shutting the application down.
func (l *APILogger) Fatal(ctx context.Context, input LogInput) {
	l.engine.WithContext(ctx).WithFields(input.GetFields()).Fatalf(input.GetMessage())
}
