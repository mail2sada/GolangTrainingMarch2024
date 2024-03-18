package main

import "fmt"

func main() {
	fmt.Println("Demo: Arrays")

	var array [7]string = [7]string{"Sunday", "Monday", "Tuesday", "Wednsday", "Thursday", "Friday", "Saturday"}

	for i, v := range array {
		fmt.Println(i, v)
	}
}
