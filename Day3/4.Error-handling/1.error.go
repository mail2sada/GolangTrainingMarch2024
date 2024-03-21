package main

import "fmt"

func Divide(a, b int) int {
	defer func() {
		a := recover()

		if a != nil {
			fmt.Println("We have received a FATAL error...")
		}
	}()
	return a / b
}

func main() {

	fmt.Println("Demo: Error handling")

	var slice []int = []int{1, 2, 3, 4, 5}

	fmt.Println(len(slice), cap(slice))

	fmt.Println(Divide(1, 0))

	fmt.Println("Reaching end of main..")
}
