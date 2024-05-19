package main

import "fmt"

func main() {

	x := foo()
	fmt.Println(x)

	y := bar()
	fmt.Println(y())
	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", bar)
	fmt.Printf("%T\n", y)

	z := runFoo()
	fmt.Println(z)
}


func foo() int {
	return 42
}


func runFoo() int {
	return foo()
}


func bar() func() int {
	return func() int {
		return 43
	}
}
