package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghp234CFEA"

func main() {
	fmt.Println("Wrlcome to handling URLs in golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)


	qparams := result.Query()
	fmt.Printf("Type of qparams is %T\n", qparams)
	fmt.Println(qparams)
	fmt.Println(qparams["coursename"])

	// Constructing a Url

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawQuery: "user=ronald",
	}

	mainUrl := partsOfUrl.String()
	fmt.Println(mainUrl)

}
