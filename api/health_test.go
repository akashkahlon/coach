package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/health", nil)
    if err != nil {
        t.Fatalf("Could not create request: %v", err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(HealthCheckHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned notOK status code: got %v want %v", status, http.StatusOK)
    }

    expected := map[string]string{"status": "healthy"}
    var response map[string]string
    err = json.NewDecoder(rr.Body).Decode(&response)
    if err != nil {
        t.Fatalf("Could not decode response: %v", err)
    }

    if response["status"] != expected["status"] {
        t.Errorf("Handler returned unexpected body: got %v want %v", response["status"], expected["status"])
    }

    if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
        t.Errorf("Handler returned wrong content type: got %v want %v", contentType, "application/json")
    }
}