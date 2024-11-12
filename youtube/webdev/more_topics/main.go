package main

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	// ...
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	// ...
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		login(w, r)
	case "/login-submit":
		loginSubmit(w, r)
	default:
		fmt.Fprintf(w, "Sup Ninja")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
