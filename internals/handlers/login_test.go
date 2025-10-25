package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/login?user=testuser", nil)
	w := httptest.NewRecorder()

	Login(w, req)

	resp := w.Result()
	cookies := resp.Cookies()
	if len(cookies) == 0 {
		t.Fatal("Expected cookie to be set")
	}
	if cookies[0].Name != "auth" || cookies[0].Value != "testuser" {
		t.Fatalf("Expected auth cookie with value testuser, got %v", cookies[0])
	}
}
