package main

import "fmt"

func main() {
	fmt.Println("Array with SDO")

	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(array)

	fmt.Println(len(array))
}
