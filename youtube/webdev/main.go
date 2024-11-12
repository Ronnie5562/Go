package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleFunction(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "Hello World!")
	case "/ninja":
		fmt.Fprintf(w, "Ronald!")
	case "/gopher":
		fmt.Fprintf(w, "Star Boy!")
	default:
		http.Error(w, "404 Not Found", 404)
	}

	log.Printf("Request Method: %s\n", r.Method)
}

func htmlVsPlain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("htmlVsPlain")
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}

func main() {
	http.HandleFunc("/", htmlVsPlain)

	log.Fatal(http.ListenAndServe(":80", nil))
}
