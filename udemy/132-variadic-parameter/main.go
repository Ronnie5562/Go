package main

import "fmt"

func main() {
	x := sum(2, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The sum is", x)
}

// func (r receiver) identifier(p parameters) (return(s)) {code}

func sum(first int, variadic ...int) int {
	fmt.Println(variadic)
	fmt.Printf("%T\n", variadic)

	n := 0
	for _, value := range variadic {
		n += value
	}

	return first * n
}
