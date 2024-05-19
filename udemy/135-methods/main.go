package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}


func (p Person) speak() {
	fmt.Println("I am", p.lastName, p.firstName)
}


func main() {
	person_1 := Person{
		firstName: "James",
		lastName: "Mary",
	}

	person_2 := Person{
		firstName: "Doe",
		lastName: "John",
	}

	person_1.speak()
	person_2.speak()
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }
