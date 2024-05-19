package main

import "fmt"

func main() {
	foo()
	bar("Ronald")
	greet := aloha("Ronald")
	fmt.Println(greet)

	s1, age := dogYears("Todd", 5)
	fmt.Println(s1, age)
}

func foo() {
	fmt.Println("I am from foo")
}

func bar(name string) {
	fmt.Println("My name is:", name)
}

func aloha(name string) string {
	return fmt.Sprint("Aloha ", name)
}

func dogYears(name string, age int) (string, int) {
	age *= age
	return fmt.Sprint(name, " is this old in dog years "), age
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { <your code here> }

/*

func syntax

no params, no returns
1 param, no returns
1 param, 1 return
2 params, 2 returns
*/
