package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	for i := 0; i < len(days); i++ {

		if i == 2 {
			goto ronnie
		}

		fmt.Println(days[i])
	}

	// This is a label
	ronnie: {
		fmt.Println("Ronnie")
	}

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Println(index, day)
	// }
}
