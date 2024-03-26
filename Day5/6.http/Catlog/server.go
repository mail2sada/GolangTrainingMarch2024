package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Product catlog server...")

	// Initialize and get handler..

	err := http.ListenAndServe(":8080", CreateCatlogMux())

	if err != nil {
		log.Fatalln("Error while create server", err)
	}
}
