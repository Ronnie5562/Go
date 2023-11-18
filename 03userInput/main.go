package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rate our Pizza out of a 5")

	// comma (ok || err)
	input, err := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)
	fmt.Println(err)
}
