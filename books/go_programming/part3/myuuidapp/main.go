package main

import (
	"fmt"
	"github.com/google/uuid"
	"rsc.io/quote"
)


func main() {
	uuid := uuid.New()
	fmt.Println(uuid)

	quote_text := quote.Go()
	fmt.Println(quote_text)
}
