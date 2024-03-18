package main

import "fmt"

func main() {
	fmt.Println("Demo for loop")

	var i = -1

	for {
		i++
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
		i++
		if i > 100 {
			break
		}
	}
}
