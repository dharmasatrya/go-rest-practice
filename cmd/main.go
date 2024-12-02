package main

import (
	"go-rest-practice/cmd/api"
	"log"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
