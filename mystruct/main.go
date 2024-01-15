package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// Disclaimer: Golang does not have inheritance; No super or parent.
	// It uses a concept called composition instead

	Ronald := User{"Ronald", "abimbola@google.dev", true, 17}
	fmt.Println(Ronald)
	fmt.Printf("%+v", Ronald)
	fmt.Printf("My name is %v, I am %v years old. You can reach out to me on %v", Ronald.Name, Ronald.Age, Ronald.Email)
}

type User struct {
	Name string
	Email string
	verified bool
	Age int
}