package pencil

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/chien-dd/library/clock"
	"github.com/sirupsen/logrus"
)

const (
	// Formatter
	DefaultPattern = "[{time}][{level}]{message}"
	ProductPattern = "[{time}][{level}]{message}"
	DevelopPattern = "[{time}][{level}][{file}-{line}]{message}"
	// Pattern
	PatternTime     = "{time}"
	PatternLevel    = "{level}"
	PatternMessage  = "{message}"
	PatternPackage  = "{package}"
	PatternFile     = "{file}"
	PatternFunction = "{function}"
	PatternLine     = "{line}"
	// Environment
	EnvironmentDevelop = "develop"
	EnvironmentProduct = "product"
)

// Formatter decorates log entries with function name and package name (optional) and line number (optional)
type Formatter struct {
	// Pattern
	Pattern string
	Time    string
}

var (
	DefaultFormatter = NewFormatter(DefaultPattern, clock.FormatRFC3339)
	DevelopFormatter = NewFormatter(DevelopPattern, clock.FormatRFC3339)
	ProductFormatter = NewFormatter(ProductPattern, clock.FormatRFC3339)

	Environment = map[string]logrus.Formatter{
		"develop": DevelopFormatter,
		"product": ProductFormatter,
	}
)

func NewFormatter(pattern, time string) logrus.Formatter {
	if pattern == "" {
		pattern = DefaultPattern
	}
	if time == "" {
		time = clock.FormatRFC3339Z
	}
	// Success
	return &Formatter{
		Pattern: pattern,
		Time:    time,
	}
}

// Format the current log entry by adding the function name and line number of the caller.
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	pattern := f.Pattern
	/*
		Start
	*/
	// Time
	pattern = strings.Replace(pattern, PatternTime, clock.Format(entry.Time, f.Time), 1)
	// Level
	pattern = strings.Replace(pattern, PatternLevel, strings.ToUpper(entry.Level.String()), 1)
	// Message
	pattern = strings.Replace(pattern, PatternMessage, entry.Message, 1)
	// Check
	okPackage := strings.Contains(pattern, PatternPackage)
	okFile := strings.Contains(pattern, PatternFile)
	okFunction := strings.Contains(pattern, PatternFunction)
	okLine := strings.Contains(pattern, PatternLine)
	if okPackage || okFile || okFunction || okLine {
		// Get position
		function, file, line := f.getPosition(entry)
		packageEnd := strings.LastIndex(function, ".")
		// Package
		if okPackage {
			pattern = strings.Replace(pattern, PatternPackage, function[:packageEnd], 1)
		}
		// File
		if okFile {
			pattern = strings.Replace(pattern, PatternFile, filepath.Base(file), 1)
		}
		// Function
		if okFunction {
			pattern = strings.Replace(pattern, PatternFunction, function[packageEnd+1:], 1)
		}
		// Line
		if okLine {
			pattern = strings.Replace(pattern, PatternLine, line, 1)
		}
	}
	for k, val := range entry.Data {
		switch v := val.(type) {
		case string:
			pattern = strings.Replace(pattern, "{"+k+"}", v, 1)
		case int:
			s := strconv.Itoa(v)
			pattern = strings.Replace(pattern, "{"+k+"}", s, 1)
		case bool:
			s := strconv.FormatBool(v)
			pattern = strings.Replace(pattern, "{"+k+"}", s, 1)
		}
	}
	// Success
	return []byte(pattern), nil
}

func (f *Formatter) getPosition(entry *logrus.Entry) (string, string, string) {
	skip := 4
	if len(entry.Data) == 0 {
		skip = 6
	}
start:
	pc, file, line, _ := runtime.Caller(skip)
	lineNumber := fmt.Sprintf("%d", line)
	function := runtime.FuncForPC(pc).Name()
	if strings.LastIndex(function, "sirupsen/logrus.") != -1 {
		skip++
		goto start
	}
	return function, file, lineNumber
}
