package logger

import (
	"fmt"
	"os"
	"time"
)

const (
	INFO = iota
	DEBUG
	ERROR
)

type Logger struct {
	level      int
	output     *os.File
	timeEnable bool
}

func CreateLogger(level int, output string, timeEnable bool) (*Logger, error) {
	var outputFile *os.File
	if output == "stdout" {
		outputFile = os.Stdout
	} else if output == "stderr" {
		outputFile = os.Stderr
	} else {
		var err error
		outputFile, err = os.OpenFile(output, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return nil, fmt.Errorf("error opening output file %s: %v", output, err)
		}
	}
	return &Logger{
		level:      level,
		output:     outputFile,
		timeEnable: timeEnable,
	}, nil
}

func (l *Logger) log(level int, msg string) {
	if level >= l.level {
		levelStr := ""
		switch level {
		case INFO:
			levelStr = "INFO"
		case DEBUG:
			levelStr = "DEBUG"
		case ERROR:
			levelStr = "ERROR"
		}
		timeStr := ""
		if l.timeEnable {
			timeStr = time.Now().Format(time.RFC3339)
		}
		fmt.Printf("%s [%s] %s", timeStr, levelStr, msg)
	}
}

func (l *Logger) Info(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.log(INFO, msg)
}

func (l *Logger) Debug(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.log(DEBUG, msg)
}
func (l *Logger) Error(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.log(ERROR, msg)
}

func (l *Logger) Close() {
	if l.output != os.Stdout && l.output != os.Stderr {
		_ = l.output.Close()
	}
}
