package glog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

// writerHook is a structure that defines a hook for logrus.
// It contains a slice of io.Writer where logs will be written,
// and LogLevels, which define the log levels that the hook will process.
type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

// Fire is called by logrus when a log entry is fired.
// It writes the log entry to all the writers defined in the hook.
// Returns an error if writing to any writer fails.
func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}

	for _, w := range hook.Writer {
		if _, err := w.Write([]byte(line)); err != nil {
			return fmt.Errorf("failed to write log entry to writer: %v", err)
		}
	}
	return err
}

// Levels returns the log levels that the hook will process.
// This method is required by logrus' Hook interface.
func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

// Logger is a wrapper around logrus.Entry to provide additional methods.
type Logger struct {
	*logrus.Entry
}

// GetLogger returns the global logger instance.
func GetLogger() *Logger {
	return &Logger{e}
}

// GetLoggerWithField returns a new Logger with a specific field added to the log entries.
func (l *Logger) GetLoggerWithField(k string, v interface{}) *Logger {
	return &Logger{l.WithField(k, v)}
}

// init initializes the logger, sets formatting, output writers, and log levels.
func init() {
	l := logrus.New()

	l.SetReportCaller(true)

	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0755)
	if err != nil {
		panic(err)
	}

	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("all log file: %s\n", allFile.Name())

	l.SetOutput(io.Discard)

	l.AddHook(&writerHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
