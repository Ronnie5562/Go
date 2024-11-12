package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to the class on requests in golang")
	// performGetRequest()
	// performPostJsonRequest()
	performPostFormRequest()
}

func performGetRequest() {
	const myurl = "http://localhost:3000/get"

	response, err := http.Get(myurl)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	// fmt.Println(response)

	fmt.Println("The status Code is: ", response.StatusCode)
	fmt.Println("The Content Length is: ", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	var responseString strings.Builder
	content2, err := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content2)
	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())

}

func performPostJsonRequest() {
	const myurl = "http://localhost:3000/post"

	requestBody := strings.NewReader(`
		{
			"name": "John Doe",
			"age": 22,
			"school": "African Leadership University",
			"church": "Deeper Life Bible Church",
			"Hobby": "Programming"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func performPostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	data := url.Values{}

	data.Add("firstname", "Ronald")
	data.Add("lastname", "Abimbola")
	data.Add("email", "starboy@google.dev")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
