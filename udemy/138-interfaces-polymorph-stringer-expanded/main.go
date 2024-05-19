// package main

// import (
// 	"fmt"
// 	"log"
// 	"strconv"
// )

// type book struct {
// 	title string
// }

// func (b book) String() string {
// 	return fmt.Sprint("The title of the book is ", b.title)
// }

// type count int

// func (c count) String() string {
// 	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
// }

// func logInfo(s fmt.Stringer) {
// 	log.Println("LOG FROM 138", s.String())
// }

// func main() {
// 	b := book{
// 		title: "West With The Night",
// 	}

// 	var n count = 42

// 	logInfo(b)
// 	logInfo(n)
// }

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }

package main

import (
	"fmt"
	"log"
	"strconv"
)

type FootBaller struct {
	name string
	age int
	club string
	shirtNumber int
}


func (f FootBaller) String() string {
	return fmt.Sprint(f.name, " is ", f.age, ", plays for ", f.club, ", and wears T-shirt Number ", f.shirtNumber)
}


type count int

func (c count) String() string {
	return fmt.Sprint("The value of this number is ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("SERVER LOG: ", s.String())
}


func main() {
	player_1 := FootBaller{
		name: "Lionel Messi",
		age: 33,
		club: "Paris Saint-Germain",
		shirtNumber: 30,
	}

	var c count = 10

	logInfo(player_1)
	logInfo(c)

}

