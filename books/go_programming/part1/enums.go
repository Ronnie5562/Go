// Enums are a way of defining a set of named constants.
// In Go, enums are typically defined using the iota keyword, which generates a sequence of numbers starting from 0.
// Enums are often used to represent a set of related values, such as days of the week, colors, or states in a state machine.

package main

import "fmt"

const (
	sunday = iota
	_ // Skip 1 which is monday
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)


func main() {
	fmt.Println(sunday, monday, tuesday, wednesday, thursday, friday, saturday)
}
