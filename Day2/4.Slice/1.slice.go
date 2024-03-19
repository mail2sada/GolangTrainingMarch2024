package main

import "fmt"

func main() {
	fmt.Println("Exploring Slice")
	var slice []int = []int{1, 2, 3, 4, 5, 6}

	fmt.Println(slice)

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for idx, v := range slice {
		fmt.Println(idx, v)
	}

	fmt.Println("Length of slice is ", len(slice))
}
