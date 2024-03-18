package main

import "fmt"

func main() {
	fmt.Println("Demo: Array filtering")

	var ar = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(ar[0:3])
	fmt.Println(ar[3:7])
	fmt.Println(ar[7:])
	fmt.Println(ar[:5])

}
