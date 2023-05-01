package main

import (
	"flag"
	"log"
	api "spotless/API"
	"spotless/database"
)

var (
	addr = flag.String("addr", ":8080", "Local address for the HTTP API")
)

func main(){
	err := database.ConnectDatabase()
	if err != nil{
		log.Fatal(err)
	}

	api := api.NewRestApi()
	if err = api.Server(*addr); err != nil{
		log.Fatal(err)
	}
}