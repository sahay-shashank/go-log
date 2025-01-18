package middleware_test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	middleware "github.com/mainframematrix/go-log/Middleware"
)

func TestBaseLogger(t *testing.T) {

	var logBuffer bytes.Buffer
	t.Log("Buffer Created")
	log.SetOutput(&logBuffer)
	t.Log("Logs assigned to Buffer")
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Serving '/'")
	})
	t.Log("Handler Created")
	middlewareLogger := middleware.PathLogger(handler)
	t.Log("Middleware Created")
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Couldn't create request! Reason: %v", err)
	}
	t.Log("Request Created")
	responseRecorder := httptest.NewRecorder()
	t.Log("Response Recorder Created")
	middlewareLogger.ServeHTTP(responseRecorder, req)
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("Expected 200 OK, but recieved: %v", status)
	}
	t.Log("HTTP request served successfully")

	expectedTimestamp := time.Now().Format(time.RFC3339)
	expectedMessage := fmt.Sprintf("[%s] GET /\n", expectedTimestamp)
	if !bytes.Contains(logBuffer.Bytes(), []byte(expectedMessage)) {
		t.Errorf("expected log output to contain %v but got %v", expectedMessage, logBuffer.String())
	}
	t.Log("Log messages match")

}
