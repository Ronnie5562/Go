package main

import "fmt"

type Book struct {
	title string
	author string
}

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%v\t%T\n", &x, &x)

	s := "James"
	fmt.Printf("%v\t%T\n", &s, &s)

	Book1 := Book {
		title: "Harry Potter",
		author: "J.K. Rowling",
	}
	fmt.Printf("%v\t%T\n", &Book1, &Book1)
	fmt.Printf("%v\t%T\n", &Book1.title, &Book1.title)
	fmt.Printf("%v\t%T\n", &Book1.author, &Book1.author)
}
