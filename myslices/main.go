package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices class in golang")

	var fruitList = []string{"Apple", "Pineapple", "Watermelon"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Cashew", "Plantain", "carrot")
	fmt.Println("Fruit list is", fruitList)


	// This works like slicing an array in python
	someFruits := append(fruitList[2:5])
	fmt.Println("This is a list of fruits sliced from the fruitList slices(Array): ", someFruits)


	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 945
	highscores[2] = 465
	highscores[3] = 867

	// The below will throw an index out of range exception
	// highscores[4] = 867

	// But the below will still work (even though we specified that the size should be 4 with the make() function)
	highscores = append(highscores, 555, 666, 777)

	fmt.Println(highscores)


	// How to remove a value from slices based on index

	var courses = []string{"Python", "Django", "Ruby", "Golang", "Swift"}
	fmt.Println(courses)
	var index int = 1
	courses = append(courses[:index], courses[index+2:]...)
	fmt.Println(courses)
}
