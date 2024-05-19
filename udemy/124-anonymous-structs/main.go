package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	person_1 := struct {
		firstName string
		lastName string
		age int
	}{
		firstName: "Ronnie",
		lastName: "Golang",
		age: 100,
	}

	person_2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Printf("%T \n", person_1)
	fmt.Printf("%T \n", person_2)
}
