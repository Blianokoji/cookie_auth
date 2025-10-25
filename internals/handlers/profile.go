package handlers

import (
	"fmt"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("auth")
	if err != nil {
		http.Error(w, "Unauthorized — please log in first", http.StatusUnauthorized)
		return
	}

	username := cookie.Value
	fmt.Fprintf(w, "Welcome back, %s!", username)
}
