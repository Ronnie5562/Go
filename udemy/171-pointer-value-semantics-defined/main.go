package main

import "fmt"

// Value semantic
func addOne(v int) int {
	return v + 1
}

// Pointer semantic
func addOneP(v *int) {
	*v += 1
}


func main() {
	// Value semantic
	value := 10
	fmt.Println(value) // 10
	fmt.Println(addOne(value)) // 11
	fmt.Println(value) // 10

	// Pointer semantic
	value2 := 20
	fmt.Println(value2) // 20
	addOneP(&value2)
	fmt.Println(value2) // 21

}