package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	a := rune('a')

	fmt.Printf("%T\n", a)

	num := 10_000_000
	fmt.Println(num)
}
