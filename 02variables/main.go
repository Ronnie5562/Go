package main

import "fmt"

const LoginToken = "mumbojombogibberish"

func main() {
	fmt.Println(LoginToken)


	var emptyString string
	fmt.Println(emptyString)

	var username string = "Ronnie5562"
	fmt.Println(emptyString)
	fmt.Printf("The Varible is of type: %T \n", username)

	// Implicit variable declarion (witout specifying types)
	var newVar = 10
	fmt.Println(newVar)

	//Without using the var keyword
	lastName := "Abimbola"
	fmt.Println(lastName)

}
