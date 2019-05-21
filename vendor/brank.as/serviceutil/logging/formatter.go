package logging

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// FluentdFormatter is similar to logrus.JSONFormatter but with log level that are recongnized
// by kubernetes fluentd.
type FluentdFormatter struct {
	TimestampFormat string
}

// Format the log entry. Implements logrus.Formatter.
func (f *FluentdFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		if err, ok := v.(error); ok {
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			v = err.Error()
		}
		data[k] = v
	}
	prefixFieldClashes(data)

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.RFC3339
	}

	data["time"] = entry.Time.Format(timestampFormat)
	data["message"] = entry.Message
	data["severity"] = entry.Level.String()

	if entry.HasCaller() {
		source := fmt.Sprintf("%s:%d %s()", entry.Caller.File, entry.Caller.Line, entry.Caller.Function)
		data["source"] = source
	}

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON: %v", err)
	}
	return append(serialized, '\n'), nil
}

// prefixFieldClashes will copy fields that are reserved for stackdriver to
// a new field to retain all information
func prefixFieldClashes(data logrus.Fields) {
	if t, ok := data["time"]; ok {
		data["fields.time"] = t
	}
	if m, ok := data["msg"]; ok {
		data["fields.msg"] = m
	}
	if l, ok := data["level"]; ok {
		data["fields.level"] = l
	}
}
