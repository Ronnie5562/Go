package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	person_1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       100,
	}

	person_2 := Person{
		firstName: "Jane",
		lastName:  "Henry",
		age:       200,
	}

	fmt.Printf("%T, %#v, %v\n", person_1, person_1, person_1)
	fmt.Printf("%T, %#v, %v\n", person_2, person_2, person_2)

	fmt.Println(person_1.firstName, person_1.age)
	fmt.Println(person_2.firstName, person_2.age)
}
