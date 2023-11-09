package log

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type LogHelper struct {
	*log.Logger
}

func New(level log.Level, jsonFormat bool) *LogHelper {
	var baseLogger = log.New()
	baseLogger.SetLevel(level)

	if jsonFormat {
		baseLogger.SetFormatter(&log.JSONFormatter{})
	} else {
		baseLogger.SetFormatter(&log.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
			DisableQuote:    true,
		})
	}

	return &LogHelper{baseLogger}
}

func (l *LogHelper) Info(message string, fields ...log.Fields) {
	if len(fields) > 0 {
		l.WithFields(fields[0]).Info(message)
	} else {
		l.Info(message)
	}
}

func (l *LogHelper) Error(message string, fields ...log.Fields) {
	if len(fields) > 0 {
		l.WithFields(fields[0]).Error(message)
	} else {
		l.Error(message)
	}
}
