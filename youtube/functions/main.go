package main

import "fmt"

func main() {
	fmt.Println("Functions in golang")

	result := adder(2, 3)
	fmt.Println("Result is", result)

	result2, message := proAdder(2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("The result is:", result2)
	fmt.Println(message)
}

func adder(value1 int, value2 int) int {
	return value1 + value2
}

func proAdder(values ...int) (int, string) {
	result := 0
	for _, value := range values {
		result += value
	}

	return result, "Thank you for using our calculator function"
}