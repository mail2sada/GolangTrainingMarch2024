package main

import "fmt"

func main() {
	fmt.Println("Exploring slice capacity")
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(len(slice), cap(slice))

	slice = append(slice, 6)
	fmt.Println(len(slice), cap(slice))

	sl := make([]int, 10, 1000)
	fmt.Println(sl)
	fmt.Println(len(sl), cap(sl))

	sl = append(sl, 1)
	sl[0] = 1000

	fmt.Println(sl)

	var sl1 []int
	fmt.Println(sl1)
	sl1 = append(sl1, 1)
	fmt.Println(sl1)
}
