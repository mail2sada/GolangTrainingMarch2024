package main

import "fmt"

func Test() {
	fmt.Println("Inside Test")
	defer fmt.Println("Calling defer from Test")
	fmt.Println("Exiting test")
}

func main() {
	fmt.Println("Exploring defer")
	defer Test()
	fmt.Println("Exiting main")
}
