package main

import (
	"graphql/internal"
	"log"
	"net/http"
)

var portNum = ":8080"

func main() {
	app := internal.GetApp()
	log.Println("Started on port", portNum)
	log.Fatal(http.ListenAndServe(portNum, app))
}
