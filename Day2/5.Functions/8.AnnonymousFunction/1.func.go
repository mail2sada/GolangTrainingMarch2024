package main

import "fmt"

func main() {
	fmt.Println("Annonymous functions")
	func() {
		fmt.Println("Inside Annonynous function")
	}()
	fmt.Println("Exiting main")

}
