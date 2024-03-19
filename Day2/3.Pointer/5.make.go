package main

import "fmt"

func main() {
	fmt.Println("exploring make")

	intPtr := make([]int, 10)

	fmt.Println(intPtr)
	intPtr[0] = 100

	fmt.Println(intPtr)

}
