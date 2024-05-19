package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	x := doMath(42, 16, add)
	fmt.Println(x)

	y := doMath(42, 16, subtract)
	fmt.Println(y)

	z := concatenateString("Abimbola", "Ronald", concat)
	fmt.Println(z)
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func concatenateString(first string, last string, f func(string, string) string) string {
	return f(first, last)
}

func concat(first string, last string) string {
	return fmt.Sprint(first, " ", last)
}
