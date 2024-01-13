package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	// Accessing Array Elements
	fmt.Println("Fruit list is: ", fruitList)
	// The below program will ouput 4 even if it was only three elements we added to it
	// This is because the len() function outputs the max number of elements an array can take
	// And above, we specified that the friutList array should take a max of 4 elements [var fruitList [4]string]
	fmt.Println("The length of the the fruit list is: ", len(fruitList))


	vegList := [3]string{"potato", "beans", "mushroom"}

	fmt.Println("The vegy list is: ", vegList)

}
