package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file_data, err := readFile("poem1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file_data)
	fmt.Println(string(file_data))
}

func readFile(filename string) ([]byte, error) {
	file_data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("there was an issue in reading the file")
	}

	return file_data, nil
}
