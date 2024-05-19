package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.first))
}

func main() {
	person_1 := person{
		first: "Dotun",
	}

	f, err := os.Create("Dotun.txt")
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	defer f.Close()

	var b bytes.Buffer

	person_1.writeOut(f)
	person_1.writeOut(&b)

	fmt.Println(b.String())
}
