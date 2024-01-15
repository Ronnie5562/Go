package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class on maps in golang")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["GO"] = "Golang"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["RT"] = "Rust"

	fmt.Println("These are the languages: ", languages)

	// Accessing values using keys
	fmt.Println("GO is:", languages["GO"])

	// To Delete
	delete(languages, "RT")
	delete(languages, "RB")

	fmt.Println("These are the remaining languages: ", languages)


	// Looping through a map in golang

	
	fmt.Println()
	fmt.Println("<== Looping through a map in golang ==>")
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}