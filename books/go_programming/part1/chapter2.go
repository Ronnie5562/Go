package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Chapter 2: Command and Control")
	fmt.Println()

	fmt.Println("if-else statement")
	isEven(14)
	fmt.Println()

	fmt.Println("switch-case statement")
	switchCase()
	fmt.Println()

	fmt.Println("Loops")
	loops()
	fmt.Println()

	fmt.Println("FizzBuzz implementation")
	fizzBuzz(15)
	fmt.Println()
}

func valiateInput(input int) error {
	if input < 0 {
		return errors.New("Input cannot be a negative number")
	} else if input > 100 {
		return errors.New("Input cannot be greater than 100")
	} else if input%7 == 0 {
		return errors.New("Input should not be divisible by 7")
	} else {
		return nil
	}
}

func isEven(input int) {
	if err := valiateInput(input); err != nil {
		fmt.Println(err)
	} else if input%2 == 0 {
		fmt.Println("The input value is even")
	} else {
		fmt.Println("The input value is odd")
	}
}

func switchCase() {
	dayBorn := time.Monday

	switch dayBorn {
	case time.Monday:
		fmt.Println("Born on Monday")
	case time.Tuesday:
		fmt.Println("Born on Tuesday")
	case time.Wednesday:
		fmt.Println("Born on Wednesday")
	case time.Thursday:
		fmt.Println("Born on Thursday")
	case time.Friday:
		fmt.Println("Born on Friday")
	case time.Saturday, time.Sunday:
		fmt.Println("Born during the weekend. Must be a cool kid, haha!")
	default:
		fmt.Println("Error, day born is not valid")
	}
}

func loops() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	names := []string{"Richard", "Sharon", "Benita", "Ronald"}
	for index, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", index, name)
	}

	config := map[string]string{
		"debug":    "1",
		"logLevel": "warn",
		"version":  "1.2.1",
	}

	for key, value := range config {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}

func fizzBuzz(input int) {
	for i := 1; i <= input; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
