package main

import (
	"go-rest-practice/cmd/api"
	"go-rest-practice/config"
	"go-rest-practice/db"
	"log"
)

func main() {
	connStr := config.InitConfig()
	db, err := db.NewPostgresStorage(connStr)

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)
	err2 := server.Run()
	if err2 != nil {
		log.Fatal(err2)
	}
}
