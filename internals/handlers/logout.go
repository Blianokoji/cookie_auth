package handlers

import (
	"fmt"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:   "auth",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "Logged out successfully.")
}
