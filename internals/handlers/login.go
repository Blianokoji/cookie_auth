package handlers

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	username := r.URL.Query().Get("user")
	if username == "" {
		http.Error(w, "missing ?user=username", http.StatusBadRequest)
		return
	}
	cookie := &http.Cookie{
		Name:     "auth",
		Value:    username,
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
	}
	http.SetCookie(w, cookie)
	fmt.Fprintf(w, "User %s logged in successfully\n", username)

}
