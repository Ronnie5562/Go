package main

import "fmt"

func main() {
	defer fmt.Println("I am done") // This will be the last thing to run because of the defer statement.
	fmt.Println("Functions")

	fmt.Println()
	fmt.Println("Variadic functions")
	variadic(10, 20, 10)
	fmt.Println()


	fmt.Println()
	fmt.Println("Anonymous functions")
	message := "Hey Rony!"
	func(str string) {
		fmt.Println("This is how to declare an anon function.")
		fmt.Println(str)
	}(message)

	// Create an anon
	number := 9
	square := func(num int) int {
		return num * num
	}

	fmt.Printf("The Square of %v is %v", number, square(number))
	fmt.Println()
}


func variadic(nums ...int) {
	fmt.Println(nums)
	fmt.Printf("%T\n", nums)
	fmt.Printf("%#v", nums)

}

// Closure
// Regular functions cannot reference variables outside of themselves; however,
// an anonymous function can reference variables external to their definition

