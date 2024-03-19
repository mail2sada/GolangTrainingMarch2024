package main

import "fmt"

func main() {
	fmt.Println("Demo: return value from function")

	fmt.Println("Sum 10, 20", add(10, 20))

	res := add(add(10, 20), add(30, 40))
	fmt.Println(res)
}

func add(a, b int) int {

	c := a + b
	return c
}
