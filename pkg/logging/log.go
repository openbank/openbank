package logging

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	outputIsTerminal = isTerminal(os.Stdout)
)

// NewLogger creates a new logrus logger with a fluentd formatter that will
// work natively with stackdriver logging
func NewLogger() *logrus.Logger {
	// Prepare a new logger
	logger := logrus.New()

	logger.Out = os.Stdout
	if !outputIsTerminal {
		logger.Formatter = &FluentdFormatter{
			TimestampFormat: time.RFC3339,
		}
	} else {
		logger.Formatter = &logrus.TextFormatter{
			TimestampFormat: time.RFC3339,
			FullTimestamp:   true,
		}
	}

	logger.SetReportCaller(true)

	return logger
}

// WithError takes an error and logger and returns a standardised error logger
func WithError(err error, logger logrus.FieldLogger) *logrus.Entry {
	return logger.WithError(err).WithField("stacktrace", fmt.Sprintf("%+v", err))
}

// isTerminal returns true if the connected os.Writer is a terminal
//
// Public method removed from logrus - copied old code
// https://github.com/sirupsen/logrus/commit/e66f22976feaed4e0a6ab43bc2b9226be8b0b979
func isTerminal(f io.Writer) bool {
	switch v := f.(type) {
	case *os.File:
		return terminal.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}
