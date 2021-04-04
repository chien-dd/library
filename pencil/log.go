package pencil

import (
	"io"

	"github.com/sirupsen/logrus"
)

type (
	Level = uint
)

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func NewLogger(out io.Writer, level Level, formatter logrus.Formatter) *logrus.Logger {
	if formatter == nil {
		formatter = DefaultFormatter
	}
	logger := &logrus.Logger{
		Out:       out,
		Level:     logrus.Level(level),
		Formatter: formatter,
	}
	// Success
	return logger
}
