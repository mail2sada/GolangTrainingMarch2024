package main

import "fmt"

func main() {
	fmt.Println("inserting in slice...")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	idx := 5
	num := 100

	slice = append(slice[:idx+1], slice[idx:]...)

	slice[idx] = num

	fmt.Println(slice)

	slice = append(slice[:idx], slice[idx+1:]...)

	fmt.Println("Just slice", slice)

	fmt.Println(slice...)

}
