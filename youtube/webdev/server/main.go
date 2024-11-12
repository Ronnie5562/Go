package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func url(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Query().Get("name")))
}

func body(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(body)
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/url":
		url(w, r)
		// http://localhost:8080/url?name=Ronald
	case "/body":
		body(w, r) // returns the request body to the client
		// http://localhost/body
	default:
		w.Write([]byte("Hello World"))
	}
}

func main() {
	http.HandleFunc("/", handleFunc)

	log.Fatal(http.ListenAndServe(":80", nil))
}
