package main

import "fmt"

func main() {
	scores := []int{67, 90, 42, 56, 85, 77, 89, 94, 39, 86}
	fmt.Println(scores)

	// Variadic paramrters
	scores = append(scores, 100, 50, 99, 10)
	fmt.Println(scores)
}
