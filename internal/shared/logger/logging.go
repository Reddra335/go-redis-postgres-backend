package logger

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	Log *logrus.Logger
}

func New() *Logger {

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	log.SetLevel(logrus.InfoLevel)

	return &Logger{
		Log: log,
	}
}

func (l *Logger) Info(msg string) {

	l.Log.Info(msg)
}
func (l *Logger) Error(msg string) {

	l.Log.Error(msg)
}

func (l *Logger) WithError(err error) *logrus.Entry {

	return l.Log.WithError(err)

}
func (l *Logger) WithFields(fields logrus.Fields) *logrus.Entry {

	return l.Log.WithFields(fields)
}
