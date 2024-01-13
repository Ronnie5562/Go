package main

import "fmt"

func main() {
	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("The value staored by the pointer is ", *ptr)

	*ptr = *ptr * 2

	fmt.Println("The value of myNumber is ", myNumber)
}
