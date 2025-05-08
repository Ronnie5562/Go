package main

import "fmt"

// removeElement removes the element at index i from slice s.
func removeElement[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s // Index out of range, return original slice
	}
	return append(s[:i], s[i+1:]...)
}

func main() {
	mySlice := []string{"apple", "banana", "cherry", "date"}

	// Remove "banana" at index 1
	mySlice = removeElement(mySlice, 1)
	fmt.Println(mySlice) // Output: [apple cherry date]

	// Remove "apple" at index 0
	mySlice = removeElement(mySlice, 0)
	fmt.Println(mySlice) // Output: [cherry date]

	// Remove element from empty slice
	emptySlice := []int{}
	emptySlice = removeElement(emptySlice, 0)
	fmt.Println(emptySlice) // Output: []

	//Remove element from out-of-bound index
	mySlice = removeElement(mySlice, 5)
	fmt.Println(mySlice) // Output: [cherry date]
}
