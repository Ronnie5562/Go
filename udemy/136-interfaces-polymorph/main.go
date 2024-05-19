package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

type SecretAgent struct {
	profile Person
	licensetokill bool
}


func (p Person) speak() {
	fmt.Println("I am", p.lastName, p.firstName)
}


func (sa SecretAgent) speak() {
	fmt.Println("I am secret Agent", sa.profile.firstName)
}


type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}


func main() {
	secret_agent_1 := SecretAgent{
		profile: Person{
			firstName: "Tamar",
			lastName: "Rabinya",
		},
		licensetokill: true,
	}

	person_1 := Person{
		firstName: "Doe",
		lastName: "John",
	}

	// secret_agent_1.speak()
	// person_1.speak()

	saySomething(person_1)
	saySomething(secret_agent_1)
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }
