package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Only log the warning severity or above.
	// Default log level
	logger.SetLevel(logrus.InfoLevel)

	EnvLogLevel := os.Getenv("LOG_LEVEL")
	if EnvLogLevel == "debug" {
		logger.SetLevel(logrus.DebugLevel)
	} else if EnvLogLevel == "info" {
		logger.SetLevel(logrus.InfoLevel)
	} else if EnvLogLevel == "error" {
		logger.SetLevel(logrus.ErrorLevel)
	} else if EnvLogLevel == "fatal" {
		logger.SetLevel(logrus.FatalLevel)
	}
}

func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}

func Info(msg string) {
	logger.Info(msg)
}

func Infof(msg string) {
	logger.Infof(msg)
}

func Debug(msg string) {
	logger.Debug(msg)
}

func Error(trace string, err error) {
	logger.WithFields(logrus.Fields{
		"line": trace,
	}).Error(err)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Println(args ...interface{}) {
	logger.Println(args...)
}

// Printf ...
func Printf(msg string, args ...interface{}) {
	logger.Printf(msg, args...)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(fields)
}

func WithError(err error) *logrus.Entry {
	return logger.WithField("error", err)
}
