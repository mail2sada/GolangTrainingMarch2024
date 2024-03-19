package main

import "fmt"

func main() {
	fmt.Println("Arguments to function")
	Greet("Hello")
	Greet("Hi")
}

func Greet(msg string) {
	fmt.Println(msg)
}
