package main

import "fmt"


func main() {
	xi := []int{0,1,2,3,4,5,6,7,8,9,10}

	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))
}

func foo(variadic ...int) int {
	sum := 0

	for _, value := range variadic {
		sum += value
	}

	return sum
}


func bar(slice_v []int) int {
	sum := 0

	for _, value := range slice_v {
		sum += value
	}

	return sum
}


/*
create a func with the identifier foo that
takes in a variadic parameter of type int
pass in a value of type []int into your func (unfurl the []int)
returns the sum of all values of type int passed in

create a func with the identifier bar that
takes in a parameter of type []int
returns the sum of all values of type int passed in

*/
