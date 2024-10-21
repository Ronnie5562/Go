package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println("Chapter 3: Core Types")

	if passwordChecker("hello") {
		fmt.Println("Password is valid!")
	} else {
		fmt.Println("Password is invalid!")
	}
}


func passwordChecker(pw string) bool {
	pwR := []rune(pw)

	if len(pwR) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}
