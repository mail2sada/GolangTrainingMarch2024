package main

import "fmt"

func main() {
	fmt.Println("Demo for loop")

	var i = 0

	for {
		fmt.Println(i)
		i++
		if i > 100 {
			break
		}
	}
}
