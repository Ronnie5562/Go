package main

import "fmt"

func main() {
	loopToTen()
	loopOverSlice()
	loopOverSliceWithRange()
}

func loopToTen() {
	fmt.Println("Looping to 10 ...")
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func loopOverSlice() {
	fmt.Println("Looping over slice ...")
	names := []string{"Alice", "Bob", "Charlie", "John", "Doe"}
	fmt.Println("Slice Capacity:", cap(names))
	fmt.Println("Slice Length:", len(names))

	for i := 0; i < len(names); i++ {
		fmt.Println("Index:", i, "Name:", names[i])
	}
}

func loopOverSliceWithRange() {
	fmt.Println("Looping over slice with range ...")
	names := []string{"Alice", "Bob", "Charlie", "John", "Doe"}

	for i, name := range names {
		fmt.Println("Index:", i, "Name:", name)
	}
}
