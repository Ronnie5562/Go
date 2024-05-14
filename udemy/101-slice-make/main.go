package main

import "fmt"

func main() {
	slice_one := []string{"a", "b", "c"}
	fmt.Println(slice_one)

	// Make a slice using the make() function
	slice_two := make([]int, 0, 10)
	fmt.Println(slice_two)
	fmt.Println(len(slice_two))
	fmt.Println(cap(slice_two))

	fmt.Println("`````````````````")

	slice_two = append(slice_two, 0,1,2,3,4,5,6,7,8,9)
	fmt.Println(slice_two)
	fmt.Println(len(slice_two))
	fmt.Println(cap(slice_two))

	fmt.Println("`````````````````")

	slice_two = append(slice_two, 10,11,12,13)
	fmt.Println(slice_two)
	fmt.Println(len(slice_two))
	fmt.Println(cap(slice_two))
}

/*
	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-------------")
*/
