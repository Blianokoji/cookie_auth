package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProfileHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/profile", nil)
	req.AddCookie(&http.Cookie{Name: "auth", Value: "tester"})
	w := httptest.NewRecorder()

	Profile(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	got := w.Body.String()
	want := "Welcome back, tester!"
	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}
