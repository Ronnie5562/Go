package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Guiness", "Wolverine Tracks"}
	fmt.Println(jb)
	fmt.Println(jm)

	two_dimensional_slice := [][]string{jb, jm}
	fmt.Println(two_dimensional_slice)
}
