package main

import "fmt"

func main() {
	person1 := Person{
		first: "Ronald",
		age: 18,
	}

	person1.speak()
}


type Person struct {
	first string
	age int
}


func (p Person) speak() {
	fmt.Printf("My name is %v and I am %v years old",  p.first, p.age)
}



/*
Create a user defined struct with the identifier "person"
- the fields:
	- first
	- age
- Attach a method to type person with the identifier "speak"
- The method should have the person say their name and age
- Create a value of type person
- Call the method from the value of type person
*/
