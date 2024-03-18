package main

import "fmt"

func main() {
	fmt.Println("Demo: Array range operator")

	array := [...]int{1: 100, 50: 10000, 3: 40, 10: 12345}

	for idx, val := range array {
		fmt.Println(idx, val)
	}
}
