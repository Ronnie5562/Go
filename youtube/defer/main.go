package main

import "fmt"

func main() {
	// In golang, when the defer keyword is place in front of a statement, it specifies that that statement should run after every other statement has completed execution
	// When you have multiple defer statement, they use the LIFO rule to execute i.e the last defer statement will be the first the be executed

	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")

	fmt.Println("Welcome to a class on the defer keyword in golang")

	moreDefer()
}



func moreDefer() {
	for value := 0; value < 5; value++ {
		defer fmt.Println(value)
	}
}

// The Output of the whole program should look like this:

// Welcome to a class on the defer keyword in golang
// 4
// 3
// 2
// 1
// 0
// Third defer
// Second defer
// First defer