package main

import (
	"log"
	"net/http"
)

type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello World</h1>"
	_, err := w.Write([]byte(msg))

	if err != nil {
		log.Printf("An error occurred: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/chapter1", func(w http.ResponseWriter, r *http.Request) {
		msg := "<h1>Chapter 1</h1>"
		_, err := w.Write([]byte(msg))

		if err != nil {
			log.Printf("An error occurred: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	http.Handle("/", hello{})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
