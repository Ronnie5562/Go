package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}

	fmt.Println("The age of Henry was", am["Henry"])

	fmt.Println(am)
	fmt.Printf("%#v\n", am)

	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37
	an["George"] = 78
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println(len(an))

	delete(an, "George")

	fmt.Println("--- accessing keys that don't exist")
	delete(an, "Georage") // won't panic
	fmt.Println(an["Georgey"]) // won't panic
	fmt.Printf("James Age: %#v\n", an["James"]) // James does not exist and the compiler won't panic [instead it prints 0 - which is the zero-value for type int]
	fmt.Println("------------------------")

	// for range over a MAP
	for k, v := range an {
		fmt.Println(k, v)
	}

	for _, v := range an {
		fmt.Println(v)
	}

	for k := range an {
		fmt.Println(k)
	}

	// for range with a SLICE
	xi := []int{42,43,44}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	for _, v := range xi {
		fmt.Println(v)
	}

	for i := range xi {
		fmt.Println(i)
	}
}
