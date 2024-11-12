package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	res, err := http.Get("http://localhost:80/url?name=Ronald")

	if err != nil {
		log.Fatal(err)
	} else {
		all, _ := ioutil.ReadAll(res.Body)
		log.Println(string(all))
	}

	res, err = http.PostForm(
		"http://localhost:80/body",
		url.Values{"name": {"Ronald"}},
	)

	if err != nil {
		log.Fatal(err)
	} else {
		all, _ := ioutil.ReadAll(res.Body)
		log.Println(string(all))
	}

	type Ninja struct {
		Name string
	}

	ronald := Ninja{Name: "Ronald"}
	ronaldJson, _ := json.Marshal(ronald)

	res, err = http.Post(
		"http://localhost:80/body",
		"application/json",
		bytes.NewBuffer(ronaldJson),
	)

	if err != nil {
		log.Fatal(err)
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
	}
}
