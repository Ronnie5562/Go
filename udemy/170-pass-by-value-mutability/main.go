package main

import "fmt"

func intDelta(n *int) {
	*n = 45
}

func sliceDelta(li []int){
	li[0] = 99
}

func mapDelta(md map[string]int, name string) {
	md[name] = 99
}

func main() {
	a := 40
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	xi := []int{1,2,3,4,5}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)

	m := map[string] int {
		"James": 32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	mapDelta(m, "James")
	fmt.Println(m)
}
