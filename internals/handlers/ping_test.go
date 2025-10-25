package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	Ping(w, req)
	if w.Body.String() != "Pong\n" {
		t.Fatalf("Expected Pong,got %q", w.Body.String())
	}
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}
}
