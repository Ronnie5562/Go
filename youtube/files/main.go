package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to a class on files in golang")

	content := "This needs to be in a file - abimbolaronald.tech"

	file, err := os.Create("./ronniegopher.txt")

	// if err != nil {
	// 	panic(err)
	// }

	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./ronniegopher.txt")
}

func readFile(filename string) {
	datebyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text file content is: ", string(datebyte))
}


func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}