package main

import "fmt"

func main() {
	fmt.Println("Demo: Array initializaion at specific index")

	array := [10]int{0: 10, 4: 100, 9: 1000}

	fmt.Println(array)
}
