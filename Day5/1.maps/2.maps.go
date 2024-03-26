package main

import "fmt"

func main() {
	fmt.Println("Maps...")

	url_store := map[int]string{0: "http://mavenir.com", 1: "http://media.mavenir.com"}

	fmt.Println("Value for key 0 is", url_store[0])

	str, ok := url_store[10]
	if !ok {
		fmt.Println("Invalid key")
	}
	fmt.Println(str)
}
