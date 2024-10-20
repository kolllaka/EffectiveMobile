package logging

import "github.com/sirupsen/logrus"

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func (l *Logger) GetLoggerWithField(k string, v interface{}) Logger {
	return Logger{l.WithField(k, v)}
}

func (l *Logger) SetLevel(lvl string) {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		level = logrus.DebugLevel
	}

	logrus.SetLevel(level)
}
