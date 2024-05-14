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

	// for range over a MAP
	for key, value := range an {
		fmt.Println(key, value)
	}

	for _, value := range an {
		fmt.Println(value)
	}

	for key := range an {
		fmt.Println(key)
	}

	for key, value := range an {
		fmt.Printf("Key: %#v\n", key)
		fmt.Printf("Value: %#v\n", value)
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