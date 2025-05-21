package main

import (
	"github/ronnie5562/fingreat_backend/api"
)

func main() {
	server := api.NewServer(".")

	server.Start(3000)
}
