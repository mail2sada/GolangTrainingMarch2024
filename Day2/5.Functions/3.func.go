package main

import "fmt"

func main() {
	fmt.Println("Multiple arguments")
	add(10, 20)
}

func add(a int, b int) {
	fmt.Println(a + b)
}
