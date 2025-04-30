package logger_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	logger "github.com/sahay-shashank/go-log"
)

func TestLogger(t *testing.T) {
	t.Run("Standard Output with INFO level", func(t *testing.T) {
		var buffer bytes.Buffer
		l, err := logger.CreateLogger(logger.INFO, "stdout", false)
		if err != nil {
			t.Fatalf("Error creating logger: %v", err)
		}

		l.SetOutput(&buffer)
		l.Log(logger.INFO, "Info message")
		l.Log(logger.DEBUG, "Debug message")
		l.Log(logger.ERROR, "Error message")

		out := buffer.String()

		if !contains(string(out), "Info message") {
			t.Fatalf("Expected 'Info message' but got %v", out)
		}
		if !contains(string(out), "Error message") {
			t.Fatalf("Expected 'Error message' but got %v", out)
		}

		// fail if debug is displayed
		if contains(string(out), "Debug message") {
			t.Fatalf("Didn't expected 'Debug message' but found it")
		}
	})
	t.Run("Standard Error with ERROR level", func(t *testing.T) {
		var buffer bytes.Buffer
		l, err := logger.CreateLogger(logger.ERROR, "stderr", true)
		if err != nil {
			t.Fatalf("Error creating logger: %v", err)
		}

		l.SetOutput(&buffer)
		l.Log(logger.INFO, "Info message")
		l.Log(logger.DEBUG, "Debug message")
		l.Log(logger.ERROR, "Error message")

		out := buffer.String()

		if !contains(string(out), "Error message") {
			t.Fatalf("Expected 'Error message' but got %v", out)
		}

		// fail if info or debug are displayed
		if contains(string(out), "Info message") {
			t.Fatalf("Didn't expected 'Info message' but found it")
		}
		if contains(string(out), "Debug message") {
			t.Fatalf("Didn't expected 'Debug message' but found it")
		}
	})

	t.Run("File with DEBUG level", func(t *testing.T) {
		tmp := "temp.log"
		defer os.Remove(tmp)

		l, err := logger.CreateLogger(logger.DEBUG, tmp, true)
		if err != nil {
			t.Fatalf("Error creating logger: %v", err)
		}

		l.Log(logger.INFO, "Info message")
		l.Log(logger.DEBUG, "Debug message")
		l.Log(logger.ERROR, "Error message")

		out, err := os.ReadFile(tmp)
		if err != nil {
			t.Fatalf("Error reading file :%v", err)
		}

		if !contains(string(out), "Info message") {
			t.Fatalf("Expected 'Info message' but got %v", out)
		}
		if !contains(string(out), "Debug message") {
			t.Fatalf("Expected 'Debug message' but got %v", out)
		}
		if !contains(string(out), "Error message") {
			t.Fatalf("Expected 'Error message' but got %v", out)
		}
	})
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
