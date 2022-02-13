package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		_, err = w.Write([]byte(line))
		if err != nil {
			return err
		}
	}
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func (l *Logger) GetLoggerWithField(k string, v interface{}) *Logger {
	return &Logger{l.WithField(k, v)}
}

func init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		FullTimestamp: true,
		ForceColors:   true,
	}
	err := os.MkdirAll("logs", 0777)
	if err != nil {
		panic(err)
	}
	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	l.SetOutput(io.Discard)
	l.AddHook(&writerHook{
		Writer: []io.Writer{allFile, os.Stdout},

		// PanicLevel level, highest level of severity. Logs and then calls panic with the message passed to Debug, Info, ...

		// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the logging level is set to Panic.

		// ErrorLevel level. Logs. Used for myErrors that should definitely be noted.
		// Commonly used for hooks to send myErrors to an error tracking service.

		// WarnLevel level. Non-critical entries that deserve eyes.

		// InfoLevel level. General operational entries about what's going on inside the application.

		// DebugLevel level. Usually only enabled when debugging. Very verbose logging.

		// TraceLevel level. Designates finer-grained informational events than the Debug.
		LogLevels: logrus.AllLevels,
	})
	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)

}
