package main

import (
	"fmt"
	"github.com/Blianokoji/cookie_auth.git/internals/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Root)
	http.HandleFunc("/ping", handlers.Ping)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/profile", handlers.Profile)
	http.HandleFunc("/logout", handlers.Logout)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}
