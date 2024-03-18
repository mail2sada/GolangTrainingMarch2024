package main

import "fmt"

func main() {
	fmt.Println("Demo: Array initializaion at specific index")

	array := [...]int{0: 10, 4: 100, 9: 1000, 10000: 10}

	fmt.Println(array)
	fmt.Println(len(array))
}
