package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Declaring variables
var (
    Debug bool = false
    LogLevel string = "info"
    startUpTime time.Time = time.Now()
)


func main() {
	rand.Seed(time.Now().UnixNano())

    r := rand.Intn(5) + 1
    stars := strings.Repeat("*", r)
    fmt.Println(stars)

    Debug := true

    fmt.Printf("Log details: %v %s %#v \n", Debug, LogLevel, startUpTime)
    fmt.Println(Debug, LogLevel, startUpTime)

    fmt.Println()
    PrintUserData()


    fmt.Println()
    fmt.Println("Operators")
    Operators()

    fmt.Println()
    fmt.Println("Test time")
    testTime()


    fmt.Println("Pointer Illustration")
    var count int
    add5Value(count)
    fmt.Println("add5Value: ", count)

    add5Pointer(&count)
    fmt.Println("add5Pointer: ", count)

    a, b := 5, 10
    swap(&a, &b)
    fmt.Println(a==10, b==5)

}


func PrintUserData() {
    first_name := "Ronald"
    last_name := "Abimbola"
    age := 18
    peanut_allergy := false


    fmt.Println("First Name: ", first_name)
    fmt.Println("Last Name: ", last_name)
    fmt.Println("Age: ", age)
    fmt.Println("Peanut Allergy: ", peanut_allergy)


    fmt.Println(Debug, LogLevel, startUpTime)
}


func Operators() {
    fmt.Println("Operators")
    fmt.Println("Arithmetic Operators")
    fmt.Println("Addition: ", 1 + 1)
    fmt.Println("Subtraction: ", 1 - 1)
    fmt.Println("Multiplication: ", 2 * 2)
    fmt.Println("Division: ", 4 / 2)
    fmt.Println("Modulus: ", 5 % 2)

    fmt.Println("Comparison Operators")
    fmt.Println("Equal: ", 1 == 1)
    fmt.Println("Not Equal: ", 1 != 1)
    fmt.Println("Greater Than: ", 2 > 1)
    fmt.Println("Less Than: ", 1 < 2)
    fmt.Println("Greater Than or Equal: ", 2 >= 1)
    fmt.Println("Less Than or Equal: ", 1 <= 2)

    fmt.Println("Logical Operators")
    fmt.Println("And: ", true && true)
    fmt.Println("Or: ", true || false)
    fmt.Println("Not: ", !true)
}


func testTime() {
    t := time.Time{}
    fmt.Println(t)
}


func add5Value(count int) {
    count += 5
    fmt.Println("add5Value: ", count)
}

func add5Pointer(count *int) {
    *count += 5
    fmt.Println("add5Pointer: ", *count)
}

func swap(a *int, b *int) {
    *b, *a = *a, *b
}
