package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study in golang")

	presentTime := time.Now()
	fmt.Println("Present time is: ", presentTime.Format("01-02-2006 15:04:05 Monday"))


	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println("Created date is: ", createdDate.Format("01-02-2006 15:04:05 Monday"))
}
