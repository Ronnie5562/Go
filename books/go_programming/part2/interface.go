package main

import "fmt"

type Speaker interface {
	Speak() string
}

type person struct {
	name      string
	age       int
	isMarried bool
}

func (p person) Speak() string {
	return "Hi my name is: " + p.name + " and I am " + fmt.Sprint(p.age) + " years old."
}

func (p person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Married: %t", p.name, p.age, p.isMarried)
}

func main() {
	p := person{
		name:      "John Doe",
		age:       30,
		isMarried: true,
	}
	fmt.Println(p.Speak())
	fmt.Println(p)
}
