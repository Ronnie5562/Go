package main

import "fmt"

func main() {
	fmt.Println("Interface")

	c := cat{}
	fmt.Println(c.Speak())
	c.Greeting()
}

type Speaker interface {
	Speak() string
}

type cat struct {
}

func (c cat) Greeting() {
	fmt.Println("Meow, Meow !!!!!!!!!!!! mmeeeeoowwww")
}

func (c cat) Speak() string {
	return "Purr Meow"
}
