package main

import (
	"log"
	"net/http"
)

type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		log.Printf("An error occurred: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", MyHandler{}))
}
